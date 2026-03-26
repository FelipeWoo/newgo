package tests

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/FelipeWoo/newgo/internal/api/routes"
	"github.com/FelipeWoo/newgo/internal/config"
)

func TestSystemRoutes(t *testing.T) {
	mux := http.NewServeMux()
	routes.RegisterSystemRoutes(mux, config.AppConfig{
		Name: "newgo",
		Env:  "test",
		Port: "8000",
	})

	req := httptest.NewRequest(http.MethodGet, "/health", nil)
	rec := httptest.NewRecorder()
	mux.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("expected status %d, got %d", http.StatusOK, rec.Code)
	}
}

func TestSumOnePlusOne(t *testing.T) {
	if 1+1 != 2 {
		t.Fatal("expected 1 + 1 to equal 2")
	}
}
