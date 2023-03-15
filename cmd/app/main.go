package main

import (
	"context"
	"fmt"
	"log"
	"os/signal"
	"syscall"

	"awesomeProject/internal"
)

func main() {
	ctx, stop := signal.NotifyContext(
		context.Background(),
		syscall.SIGHUP, syscall.SIGINT, syscall.SIGQUIT, syscall.SIGILL,
	)

	go func() {
		if err := recover(); err != nil {
			log.Println(fmt.Sprintf("Application panic: %v\n", err))
			stop()
		}
	}()

	app := internal.NewApp(ctx)

	if err := app.Run(); err != nil {
		log.Fatalln(err)
	}

	log.Println("application has been stopped")
}
