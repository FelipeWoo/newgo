package tests

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"newgo/internal/router/routes"
)

func init() {
	InitTestEnv("sample_test")
}

func TestSystemRoutes(t *testing.T) {
	mux := http.NewServeMux()
	routes.RegisterSystemRoutes(mux)

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
