package App

import (
	"context"
	"github.com/joho/godotenv"
	"kz.nitec.digidocs.pcr/internal/Service"
	"kz.nitec.digidocs.pcr/internal/config"
	delivery "kz.nitec.digidocs.pcr/internal/delivery/http"
	"kz.nitec.digidocs.pcr/internal/repository"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func Run() {
	if err := godotenv.Load("build/.env", "build/.local_env"); err != nil {
		log.Println(err)
		return
	}

	configs := config.GetConfig()

	db, err := repository.NewPostgresDB(configs.DB)
	if err != nil {
		log.Println(err)
	}

	// initialize pcr repository
	repos := repository.NewRepositories(db)
	services := Service.NewServices(Service.Deps{
		repos,
		configs.Shep,
	})

	handlers := delivery.NewHandler(services)

	srv := http.Server{
		Addr:    configs.App.Port,
		Handler: handlers.InitRoutes(),
	}

	go func() {
		log.Println("Starting server on port", configs.App.Port)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal(err)
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
		log.Fatal("Server forced to shutdown:", err)
	}
}
