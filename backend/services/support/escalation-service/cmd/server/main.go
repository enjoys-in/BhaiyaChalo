package main

import (
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"os"

	"github.com/enjoys-in/BhaiyaChalo/services/support/escalation-service/internal/config"
	"github.com/enjoys-in/BhaiyaChalo/services/support/escalation-service/internal/event"
	"github.com/enjoys-in/BhaiyaChalo/services/support/escalation-service/internal/handler"
	"github.com/enjoys-in/BhaiyaChalo/services/support/escalation-service/internal/repository"
	"github.com/enjoys-in/BhaiyaChalo/services/support/escalation-service/internal/service"
)

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	cfg := config.Load()

	// TODO: init db connection
	repo := repository.NewRepository(nil)
	pub := event.NewKafkaPublisher(logger)
	svc := service.NewEscalationService(repo, pub)
	h := handler.NewEscalationHandler(svc)

	mux := http.NewServeMux()
	h.RegisterRoutes(mux)

	mux.HandleFunc("GET /healthz", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("ok"))
	})

	addr := fmt.Sprintf(":%d", cfg.Port)
	logger.Info("starting server", "addr", addr, "service", "escalation-service")
	if err := http.ListenAndServe(addr, mux); err != nil {
		log.Fatal(err)
	}
}
