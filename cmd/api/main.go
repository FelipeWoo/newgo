package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"newgo/internal/bootstrap"
	"newgo/internal/config"
	"newgo/internal/logger"
	"newgo/internal/router"
)

func main() {

	bootstrap.InitAll()
	defer bootstrap.Shutdown()

	logger.SetModule("main")
	logger.Info("Starting %s API", config.Config.AppName)

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	logger.SetModule("main")
	logger.Info("Ready to process requests ...")
	server := router.NewServer(config.Config)

	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Fail("server error: %v", err)
		}
	}()

	// Esperar señal de cierre (Ctrl+C, kill, etc.)
	<-ctx.Done()
	logger.Warn("Shutdown signal received")

	// Cerrar servidor web limpio
	if err := server.Shutdown(context.Background()); err != nil {
		logger.Fail("server shutdown error: %v", err)
	} else {
		logger.Info("API server shut down gracefully")
	}

}
