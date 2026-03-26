package api

import (
	"fmt"
	"net/http"

	"github.com/FelipeWoo/newgo/internal/api/routes"
	"github.com/FelipeWoo/newgo/internal/config"
)

func NewServer(cfg config.AppConfig) *http.Server {
	mux := http.NewServeMux()
	routes.RegisterSystemRoutes(mux, cfg)

	return &http.Server{
		Addr:    fmt.Sprintf(":%s", cfg.Port),
		Handler: mux,
	}
}
