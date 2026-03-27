package main

import (
	"context"
	"os"
	"os/signal"

	"newgo/cmd/cli"
	"newgo/internal/bootstrap"
	"newgo/internal/config"
	"newgo/internal/logger"
)

func main() {
	bootstrap.InitAll()

	logger.SetModule("main")
	logger.Info("Starting %s app", config.Config.AppName)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Canal de señales
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt)

	logger.Info("Ready to process requests ...")

	// 1) Listen Ctrl+C
	go func() {
		<-sigChan
		logger.Warn("Canceled by Ctrl+C")
		cancel()
	}()

	// 2) Run CLI
	go func() {
		err := cli.Menu(ctx)
		if err != nil {
			logger.Fail("app error: %v", err)
			cancel()
		}
		cancel() // for now
	}()

	// 3) Main
	<-ctx.Done()

	logger.Info("Shutting down...")
}
