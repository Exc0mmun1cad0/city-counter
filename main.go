package main

import (
	"context"
	"test-app/cache"
)

func main() {
	ctx := context.Background()

	cache, _ := cache.NewRedisCache(ctx, "localhost:6379", "", 0)

	_ = cache
}
