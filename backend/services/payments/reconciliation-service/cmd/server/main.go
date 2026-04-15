package main

import (
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"os"

	"github.com/enjoys-in/BhaiyaChalo/services/payments/reconciliation-service/internal/config"
	"github.com/enjoys-in/BhaiyaChalo/services/payments/reconciliation-service/internal/event"
	"github.com/enjoys-in/BhaiyaChalo/services/payments/reconciliation-service/internal/handler"
	"github.com/enjoys-in/BhaiyaChalo/services/payments/reconciliation-service/internal/repository"
	"github.com/enjoys-in/BhaiyaChalo/services/payments/reconciliation-service/internal/service"
)

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	cfg := config.Load()

	// TODO: init db connection
	repo := repository.NewRepository(nil)
	pub := event.NewKafkaPublisher(logger)
	svc := service.NewReconciliationService(repo, pub)
	h := handler.NewReconciliationHandler(svc)

	mux := http.NewServeMux()
	h.RegisterRoutes(mux)

	mux.HandleFunc("GET /healthz", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("ok"))
	})

	addr := fmt.Sprintf(":%d", cfg.Port)
	logger.Info("starting server", "addr", addr, "service", "reconciliation-service")
	if err := http.ListenAndServe(addr, mux); err != nil {
		log.Fatal(err)
	}
}
