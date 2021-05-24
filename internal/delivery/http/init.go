package http

import (
	"context"
	"database/sql"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	config2 "kz.nitec.digidocs.pcr/internal/config"
	models2 "kz.nitec.digidocs.pcr/internal/models"
	"log"
	"net/http"
	"time"
)

type App struct {
	Config *config2.MainConfig
	DB     *sql.DB
	Router *gin.Engine

	Pcr models2.PcrRepository
}

func (a *App) Run(ctx context.Context) {

	srv := http.Server{
		Addr:    a.Config.App.Port,
		Handler: a.Router,
	}

	go func() {
		log.Println("Starting server on port", a.Config.App.Port)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal(err)
		}
	}()

	<-ctx.Done()

	log.Println("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown:", err)
	}

	log.Println("Server exiting")
}
