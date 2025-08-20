package server

import (
	"context"
	"net/http"
	"testing"
	"time"

	"wb_l0/internal/cache"
	"wb_l0/internal/db"

	"github.com/gorilla/mux"
	"go.uber.org/zap"
)

func TestServer(t *testing.T) {
	log, err := zap.NewDevelopment()
	if err != nil {
		t.Fatal(err)
	}

	if err := db.InitDB(log); err != nil {
		t.Fatal("Failed to initialize database:", err)
	}

	if err := cache.InitCache(log); err != nil {
		t.Fatal("Failed to initialize cache:", err)
	}

	logger = log

	router := mux.NewRouter()
	router.HandleFunc("/order/{id}", GetOrderHandler).Methods("GET")

	srv := &http.Server{
		Handler: router,
		Addr:    ":8081",
	}

	errChan := make(chan error, 1)
	go func() {
		errChan <- srv.ListenAndServe()
	}()

	time.Sleep(time.Second)

	resp, err := http.Get("http://localhost:8081/order/nonexistent")
	if err != nil {
		t.Fatal("Failed to make GET request:", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusNotFound {
		t.Errorf("Expected status 404, got %d", resp.StatusCode)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		t.Fatal("Failed to shutdown server:", err)
	}

	if err := <-errChan; err != http.ErrServerClosed {
		t.Errorf("Unexpected error from ListenAndServe: %v", err)
	}
}
