package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"

	"newgo/cmd/api/router"
	"newgo/internal/config"
	//"newgo/cmd/router"
)

func main() {

	cfg := config.Init()
	defer config.Shutdown()

	config.SetModule("main")
	config.Info("Starting %s app", cfg.AppName)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Signal handling
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt)

	config.SetModule("main")
	config.Info("Ready to process requests ...")
	server := router.NewServer(cfg)

	go func() {
		<-sigChan
		config.Warn("Canceled by Ctrl+C")
		cancel()
	}()

	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			config.Fail("server error: %v", err)
		}
	}()

	// err := cli.Menu(ctx)
	// if err != nil {
	// 	logger.Fail("app error: %v", err)
	// }

	// logger.Info("Shutting down...")

	// Wait for shutdown signal (Ctrl+C, kill, etc.)
	<-ctx.Done()
	config.Warn("Shutdown signal received")

	// Close the server gracefully
	if err := server.Shutdown(context.Background()); err != nil {
		config.Fail("server shutdown error: %v", err)
	} else {
		config.Info("API server shut down gracefully")
	}

}
