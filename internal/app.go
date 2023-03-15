package internal

import (
	"awesomeProject/internal/io/http"
	"context"
	"fmt"
	"log"
	"os"
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
	host := fmt.Sprintf(":%d", defaultServerPort)
	if s := os.Getenv("APP_PORT"); s != "" {
		host = s
	}

	log.Println("server init")
	server := http.New()

	errCh := make(chan error)
	go func() {
		if err := server.Run(host); err != nil {
			errCh <- err
		}
	}()

	log.Println("server run and ready to accept connection")

	select {
	case <-a.ctx.Done():
		log.Println("get shutdown signal")
	case err := <-errCh:
		return err
	}

	return server.Shutdown()
}
