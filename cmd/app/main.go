package app

import (
	"awesomeProject/internal"
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	ctx, stop := context.WithCancel(context.Background())
	go onSignal(stop)
	defer graceful(stop)
	app := internal.NewApp(ctx)
	if err := app.Run(); err != nil {
		log.Fatalln(err)
	}
}

func onSignal(stop context.CancelFunc) {
	ch := make(chan os.Signal)
	signal.Notify(ch, syscall.SIGHUP, syscall.SIGINT, syscall.SIGQUIT, syscall.SIGILL)
	defer signal.Stop(ch)
	<-ch
	stop()
}

func graceful(stop context.CancelFunc) {
	if err := recover(); err != nil {
		log.Println(fmt.Sprintf("Application panic: %v\n", err))
		stop()
	}
}
