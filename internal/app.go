package internal

import (
	"awesomeProject/internal/io/http"
	"context"
	"log"
	"os"
	"strconv"
)

const defaultServerPort = 8081

type App struct {
	ctx context.Context
}

func NewApp(ctx context.Context) *App {
	return &App{
		ctx: ctx,
	}
}

func (a *App) Run() error {
	serverPort := defaultServerPort
	if serverPortStr := os.Getenv("APP_PORT"); serverPortStr != "" {
		var envParseErr error
		serverPort, envParseErr = strconv.Atoi(serverPortStr)
		if envParseErr != nil {
			log.Printf("error on parse env APP_PORT. Error: %s. Server run on default port: %d", envParseErr, serverPort)
		}
	}
	log.Println("server init")
	server := http.New()
	server.Run(serverPort)
	log.Println("server run and ready to accept connection")
	<-a.ctx.Done()
	log.Println("get shutdown signal")
	return server.Shutdown()
}
