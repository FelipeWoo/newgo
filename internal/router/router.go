package router

import (
	"fmt"
	"net/http"

	"newgo/internal/config"
	"newgo/internal/router/routes"
)

func NewServer(cfg config.AppConfig) *http.Server {
	mux := http.NewServeMux()
	routes.RegisterSystemRoutes(mux)

	return &http.Server{
		Addr:    fmt.Sprintf(":%d", cfg.Port),
		Handler: mux,
	}
}
