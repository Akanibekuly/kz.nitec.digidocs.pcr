package handler

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	config2 "kz.nitec.digidocs.pcr/internal/config"
	models2 "kz.nitec.digidocs.pcr/internal/models"
	repository2 "kz.nitec.digidocs.pcr/internal/repository"
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

func (a *App) Initialize(conf *config2.MainConfig) {
	a.Config = conf

	gin.SetMode(conf.App.Mode)
	a.Router = gin.New()
	a.Router.Use(gin.Logger())
	a.Router.Use(gin.Recovery())
	a.setRoutes()
	db, err := a.getConnection()
	if err != nil {
		log.Printf("could not connect to db %s", err.Error())
		return
	}
	a.Pcr = repository2.PcrRepositoryInit(db)
}

func (a *App) getConnection() (*sql.DB, error) {
	var db *sql.DB
	var err error

	switch a.Config.DB.Dialect {
	case "postgres":
		dbURI := fmt.Sprintf("port=%d host=%s user=%s password=%s dbname=%s sslmode=disable",
			a.Config.DB.Port,
			a.Config.DB.Host,
			a.Config.DB.Username,
			a.Config.DB.Password,
			a.Config.DB.DBName,
		)
		db, err = sql.Open(a.Config.DB.Dialect, dbURI)

	case "mysql":
		// TODO:
	case "cassandra":
		// TODO:
	default:
		return nil, fmt.Errorf("could not connect to %s", a.Config.DB.Dialect)
	}

	return db, err
}

func (a *App) setRoutes() {
	pcr := a.Router.Group("/digilocker/pcr-cert/api")
	{
		pcr.POST("/pcr-result", a.Process)
	}
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
