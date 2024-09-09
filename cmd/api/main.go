package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	rediscache "test-app/cache/redis"
	"test-app/config"
	app "test-app/container"
	"test-app/storage/postgres"

	_ "github.com/lib/pq"
)

func main() {
	err := run()
	if err != nil {
		log.Println(err)
	}
}

func run() error {
	ctx := context.Background()

	// loading configuration file for HTTP server, Postgres and Redis
	cfg := config.MustLoad()

	// initializing Redis cache storage
	cache, err := rediscache.NewRedisCache(
		ctx,
		cfg.Redis.Addr,
		cfg.Redis.Password,
		cfg.Redis.DB,
	)
	if err != nil {
		return fmt.Errorf("cannot initialize cache: %w", err)
	}

	// initializing Postgres database storage
	storage, err := postgres.NewPostgresStorage(
		cfg.Postgres.Host,
		cfg.Postgres.User,
		cfg.Postgres.DBName,
		cfg.Postgres.Password,
		cfg.Postgres.SSLMode,
	)
	if err != nil {
		return fmt.Errorf("cannot initialize db: %w", err)
	}

	// putting all together to container
	app := app.NewApp(storage, cache)

	// initializing routing
	router := http.NewServeMux()
	router.HandleFunc("/", app.GetStats)

	// initializing HTTP server
	srv := &http.Server{
		Addr:        cfg.HTTPServer.Addr,
		ReadTimeout: cfg.HTTPServer.ReadTimeout,
		IdleTimeout: cfg.HTTPServer.IdleTimeout,
		Handler:     router,
	}

	err = srv.ListenAndServe()
	if err != nil {
		return fmt.Errorf("cannot run HTTP server: %w", err)
	}

	return nil
}
