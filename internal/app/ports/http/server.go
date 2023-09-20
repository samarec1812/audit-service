package http

import (
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"golang.org/x/exp/slog"

	"github.com/samarec1812/audit-service/internal/app/service"
	"github.com/samarec1812/audit-service/internal/config"
)

func NewHTTPServer(cfg config.HTTPServer, logger *slog.Logger, a service.App) *http.Server {
	handler := chi.NewRouter()

	s := &http.Server{
		Addr:         cfg.Address,
		Handler:      handler,
		ReadTimeout:  cfg.Timeout * time.Second,
		WriteTimeout: cfg.Timeout * time.Second,
		IdleTimeout:  cfg.IdleTimeout * time.Second,
	}

	AppRouter(handler, logger, a)

	return s
}
