package main

import (
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"os"
	"github.com/enjoys-in/BhaiyaChalo/services/realtime-core/connection-manager-service/internal/config"
	"github.com/enjoys-in/BhaiyaChalo/services/realtime-core/connection-manager-service/internal/handler"
)

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	cfg := config.Load()
	_ = cfg

	// TODO: init db, repo, event publisher, service, handler
	_ = handler.NewHandler

	mux := http.NewServeMux()
	mux.HandleFunc("GET /healthz", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("ok"))
	})

	addr := fmt.Sprintf(":%d", cfg.Port)
	logger.Info("starting server", "addr", addr, "service", "connection-manager-service")
	if err := http.ListenAndServe(addr, mux); err != nil {
		log.Fatal(err)
	}
}
