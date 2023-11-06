package api

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os/signal"
	"syscall"
	"time"

	"github.com/samber/do"

	"matcha/internal/config"
	"matcha/internal/core/repository"
	"matcha/internal/pkg/broom"
	"matcha/internal/pkg/db"
)

func Run(ctx context.Context) error {
	defer broom.Clean()
	notifyCtx, notifyCancel := signal.NotifyContext(ctx, syscall.SIGINT, syscall.SIGTERM)
	defer notifyCancel()
	serverCtx, serverCancel := context.WithCancelCause(ctx)
	defer serverCancel(nil)

	i, err := NewContainer(config.Get())
	if err != nil {
		return err
	}
	server := do.MustInvoke[*http.Server](i)
	go func() {
		if err := server.ListenAndServe(); !errors.Is(err, http.ErrServerClosed) {
			serverCancel(err)
		}
		serverCancel(nil)
	}()

	select {
	case <-notifyCtx.Done():
		shutdownCtx, shutdownCancel := context.WithTimeout(context.Background(), 1*time.Minute)
		defer shutdownCancel()
		if err := server.Shutdown(shutdownCtx); err != nil {
			return fmt.Errorf("failed to shutdown gracefully: %v", err)
		}
	case <-serverCtx.Done():
		if cause := context.Cause(serverCtx); !errors.Is(cause, context.Canceled) {
			return cause
		}
	}
	return nil
}

func NewContainer(cfg config.Config) (*do.Injector, error) {
	sqldb, err := db.Open(cfg.DBConnString())
	if err != nil {
		return nil, err
	}
	broom.Gather(func() { _ = sqldb.Close })

	i := do.New()
	do.Provide(i, NewServer)
	do.ProvideNamedValue(i, "ServerAddress", cfg.ServerAddress)
	do.Provide(i, NewResolver)
	do.Provide(i, repository.NewManager)
	do.ProvideValue(i, sqldb)
	return i, nil
}
