package rediscache

import (
	"context"
	"fmt"
	"testing"
)

// Just a temporary solution
// TODO: rewrite it as a normal gopher
func TestCache(t *testing.T) {
	ctx := context.Background()

	cache, err := NewRedisCache(ctx, "localhost:6379", "", 0)
	fmt.Println(err == nil, err)

	err = cache.Insert(ctx, "1", "2")
	fmt.Println(err == nil, err)

	val, err := cache.Get(ctx, "1")
	fmt.Println(err == nil, err)
	fmt.Println(val == "2", val)

	stCMD := cache.cache.FlushAll(ctx)
	fmt.Println(stCMD.Result())
}
