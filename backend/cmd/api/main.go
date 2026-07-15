package main

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"example.com/go-project_1/backend/internal/config"
	"example.com/go-project_1/backend/internal/database/postgres"
	"example.com/go-project_1/backend/internal/handlers"
	appmiddleware "example.com/go-project_1/backend/internal/middleware"
	"example.com/go-project_1/backend/internal/routes"
	"example.com/go-project_1/backend/internal/services"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("load config: %v", err)
	}

	ctx := context.Background()

	if cfg.DatabaseURL != "" {
		pool, err := postgres.New(ctx, cfg.DatabaseURL)
		if err != nil {
			log.Fatalf("connect postgres: %v", err)
		}
		defer pool.Close()
	}

	healthHandler := handlers.NewHealthHandler()
	authHandler := handlers.NewAuthHandler(services.NewAuthService(nil, nil))

	server := &http.Server{
		Addr:         ":" + cfg.Port,
		Handler:      appmiddleware.Chain(routes.New(routes.Dependencies{Health: healthHandler, Auth: authHandler})),
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	serverErr := make(chan error, 1)
	go func() {
		log.Printf("api listening on %s", server.Addr)
		serverErr <- server.ListenAndServe()
	}()

	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, syscall.SIGINT, syscall.SIGTERM)

	select {
	case err := <-serverErr:
		if err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("server error: %v", err)
		}
	case sig := <-shutdown:
		log.Printf("received signal %s, shutting down", sig)
		shutdownCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		if err := server.Shutdown(shutdownCtx); err != nil {
			log.Fatalf("shutdown: %v", err)
		}
	}
}