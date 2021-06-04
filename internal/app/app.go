package app

import (
	"context"
	"fmt"
	"github.com/joho/godotenv"
	"kz.nitec.digidocs.pcr/internal/config"
	delivery "kz.nitec.digidocs.pcr/internal/delivery/http/v1"
	"kz.nitec.digidocs.pcr/internal/repository"
	"kz.nitec.digidocs.pcr/internal/service"
	"kz.nitec.digidocs.pcr/pkg/logger"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func Run() {
	if err := godotenv.Load("build/.local_env"); err != nil {
		logger.PrintLog("ERROR", "PCR-TM", "", logger.CreateMessageLog(err))
		return
	}

	configs, err := config.GetConfig()
	if err != nil {
		logger.PrintLog("ERROR", "PCR-TM", "", logger.CreateMessageLog(err))
		return
	}

	db, err := repository.NewPostgresDB(configs.DB)
	if err != nil {
		logger.PrintLog("ERROR", "PCR-TM", "", err)
		return
	}
	defer db.Close()

	// initialize pcr repository
	repos := repository.NewRepositories(db)
	services := service.NewServices(service.Deps{
		repos,
		configs.Shep,
		configs.Services,
	})

	handlers := delivery.NewHandler(services)

	srv := http.Server{
		Addr:    configs.App.Port,
		Handler: handlers.InitRoutes(),
	}

	go func() {
		msg:=fmt.Sprintf("Starting server on port: %s" , configs.App.Port)
		logger.PrintLog("INFO","PCR_TM", "",msg)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.PrintLog("ERROR", "PCR-TM", "", logger.CreateMessageLog(err))
			return
		}
	}()

	// Graceful Shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)

	<-quit

	msg:=fmt.Sprintf("Shutting down server...")
	logger.PrintLog("INFO", "PCR_TM", "", msg)

	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		msg:=fmt.Sprintf("Server forced to shutdown: %s", err)
		logger.PrintLog("ERROR", "PCR_TM", "",msg)
	}
}
