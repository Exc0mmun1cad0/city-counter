package cache

import (
	"context"
)

type Cache interface {
	Insert(ctx context.Context, key, value string) error
	Get(ctx context.Context, key string) (string, error)
	Delete(ctx context.Context, key string) error
}
