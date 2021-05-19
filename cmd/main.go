package main

import (
	"context"
	config2 "kz.nitec.digidocs.pcr/internal/config"
	handler2 "kz.nitec.digidocs.pcr/internal/handler"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load("./../build/.env", "./../build/.local_env"); err != nil {
		log.Print("No .env file found")
	}
}

func main() {
	conf := config2.GetConfig()

	app := &handler2.App{}
	app.Initialize(conf)

	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		quit := make(chan os.Signal)
		signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
		<-quit
		cancel()
	}()

	app.Run(ctx)
}
