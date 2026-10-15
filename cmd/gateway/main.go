package main

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
	"os"

	"github.com/eduardotashiro/paycrash/internal/config"
)

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelInfo,
	}))
	slog.SetDefault(logger)

	conf, err := config.Load()
	if err != nil {
		logger.Error("failed to load config", slog.String("error", err.Error()))
		os.Exit(1)
	}

	http.HandleFunc("/healthcheck", healthHandler)

	slog.Info("server starting")
	slog.String("port", conf.Port)
	slog.String("port", conf.Env)

	if err := http.ListenAndServe(conf.Port, nil); err != nil {
		fmt.Printf("Error: %v\n", err)
	}
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	response := map[string]string{
		"status": "up",
	}
	json.NewEncoder(w).Encode(response)
}
