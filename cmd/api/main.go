package main

import (
	"context"
	"fmt"
	rediscache "test-app/cache/redis"
	"test-app/config"
	"test-app/storage/postgres"

	_ "github.com/lib/pq"
)

func main() {
	run()

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
	db, err := postgres.NewPostgresStorage(
		cfg.Postgres.Host,
		cfg.Postgres.User,
		cfg.Postgres.DBName,
		cfg.Postgres.Password,
		cfg.Postgres.SSLMode,
	)
	if err != nil {
		return fmt.Errorf("cannot initialize db: %w", err)
	}

	_ = cache
	_ = db

	return nil
}
