package cache

import (
	"context"
)

type Cache interface {
	Insert(ctx context.Context, key string, value int) error
	Get(ctx context.Context, key string) (int, error)
	Delete(ctx context.Context, key string) error
}
