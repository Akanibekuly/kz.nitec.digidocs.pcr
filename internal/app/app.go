package app

import (
	"context"
	"github.com/joho/godotenv"
	"kz.nitec.digidocs.pcr/internal/config"
	delivery "kz.nitec.digidocs.pcr/internal/delivery/http/v1"
	"kz.nitec.digidocs.pcr/internal/repository"
	"kz.nitec.digidocs.pcr/internal/service"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func Run() {
	if err := godotenv.Load("build/.local_env"); err != nil {
		log.Println(err)
		return
	}

	configs, err := config.GetConfig()
	if err != nil {
		log.Println(err)
		return
	}

	db, err := repository.NewPostgresDB(configs.DB)
	if err != nil {
		log.Println(err)
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
		log.Println("Starting server on port", configs.App.Port)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Println(err)
			return
		}
	}()

	// Graceful Shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)

	<-quit

	log.Println("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Println("Server forced to shutdown:", err)
	}
}
