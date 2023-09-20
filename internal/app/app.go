package app

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"golang.org/x/exp/slog"
	"golang.org/x/sync/errgroup"

	ht "github.com/samarec1812/audit-service/internal/app/ports/http"
	"github.com/samarec1812/audit-service/internal/app/repository/postgres/event"
	"github.com/samarec1812/audit-service/internal/app/service"
	"github.com/samarec1812/audit-service/internal/config"
	"github.com/samarec1812/audit-service/internal/pkg/logger"
	"github.com/samarec1812/audit-service/internal/pkg/postgres"
)

const (
	ctxTimeout = 30
)

func Run(cfg *config.Config) {

	log := logger.SetupLogger(cfg.Env)
	log.Info("starting application")
	log.Debug("debug message")

	db, err := postgres.Connect(cfg.DatabaseURL)
	if err != nil {
		log.Error("error connect database", err)
		os.Exit(1)
	}

	log.Info("database connect successful")
	eventRepo := event.NewEventRepository(db)

	app := service.NewApp(eventRepo)
	srv := ht.NewHTTPServer(cfg.HTTPServer, log, app)

	eg, ctx := errgroup.WithContext(context.Background())
	sigQuit := make(chan os.Signal, 1)
	signal.Ignore(syscall.SIGHUP, syscall.SIGPIPE)
	signal.Notify(sigQuit, syscall.SIGINT, syscall.SIGTERM)

	eg.Go(func() error {
		select {
		case s := <-sigQuit:
			log.Info("captured signal:", slog.String("signal", s.String()))
			return fmt.Errorf("captured signal: %v", s)
		case <-ctx.Done():
			return nil
		}
	})

	eg.Go(func() error {
		log.Info("starting http server, listening on:", slog.String("address", srv.Addr))
		defer log.Info("close http server listening on:", slog.String("address", srv.Addr))

		errCh := make(chan error)

		defer func() {
			shCtx, cancel := context.WithTimeout(context.Background(), ctxTimeout*time.Second)
			defer cancel()

			if err = srv.Shutdown(shCtx); err != nil {
				log.Error("can't close http server listening on %s: %s", slog.String("address", srv.Addr), slog.String("error", err.Error()))
				os.Exit(1)
			}

			close(errCh)
		}()

		go func() {
			if err = srv.ListenAndServe(); !errors.Is(err, http.ErrServerClosed) {
				errCh <- err
			}
		}()

		select {
		case <-ctx.Done():
			return ctx.Err()
		case err = <-errCh:
			return fmt.Errorf("http server can't listen and serve requests: %w", err)
		}
	})

	if err = eg.Wait(); err != nil {
		log.Info("gracefully shutting down the servers:", slog.String("error", err.Error()))
	}
}
