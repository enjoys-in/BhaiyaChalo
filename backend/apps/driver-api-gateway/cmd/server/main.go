package main

import (
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"os"

	"github.com/enjoys-in/BhaiyaChalo/apps/driver-api-gateway/internal/config"
	"github.com/enjoys-in/BhaiyaChalo/apps/driver-api-gateway/internal/middleware"
	"github.com/enjoys-in/BhaiyaChalo/apps/driver-api-gateway/internal/route"
)

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	cfg := config.Load()

	router := route.NewRouter()

	// Wrap with logging middleware
	h := middleware.LoggingMiddleware(logger)(router)

	addr := fmt.Sprintf(":%d", cfg.Port)
	logger.Info("starting gateway", "addr", addr, "service", "driver-api-gateway")
	if err := http.ListenAndServe(addr, h); err != nil {
		log.Fatal(err)
	}
}
