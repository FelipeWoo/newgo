package main

import (
	"fmt"
	"log"

	"github.com/FelipeWoo/newgo/internal/api"
	"github.com/FelipeWoo/newgo/internal/bootstrap"
)

func main() {
	app, err := bootstrap.Boot("main")
	if err != nil {
		log.Fatalf("bootstrap failed: %v", err)
	}

	server := api.NewServer(app.Config)

	app.Logger.Info("application ready", "name", app.Config.Name, "port", app.Config.Port)
	fmt.Printf("Application: %s\n", app.Config.Name)

	if err := server.ListenAndServe(); err != nil {
		app.Logger.Error("server stopped", "error", err)
		log.Fatal(err)
	}
}
