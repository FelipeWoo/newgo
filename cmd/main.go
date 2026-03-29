package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"

	"newgo/internal/bootstrap"
	"newgo/internal/config"
	"newgo/internal/logger"
	"newgo/internal/router"
	//"newgo/cmd/router"
)

func main() {

	bootstrap.InitAll()
	defer bootstrap.Shutdown()

	logger.SetModule("main")
	logger.Info("Starting %s app", config.Config.AppName)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Signal handling
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt)

	logger.SetModule("main")
	logger.Info("Ready to process requests ...")
	server := router.NewServer(config.Config)

	go func() {
		<-sigChan
		logger.Warn("Canceled by Ctrl+C")
		cancel()
	}()

	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Fail("server error: %v", err)
		}
	}()

	// err := cli.Menu(ctx)
	// if err != nil {
	// 	logger.Fail("app error: %v", err)
	// }

	// logger.Info("Shutting down...")

	// Wait for shutdown signal (Ctrl+C, kill, etc.)
	<-ctx.Done()
	logger.Warn("Shutdown signal received")

	// Close the server gracefully
	if err := server.Shutdown(context.Background()); err != nil {
		logger.Fail("server shutdown error: %v", err)
	} else {
		logger.Info("API server shut down gracefully")
	}

}
