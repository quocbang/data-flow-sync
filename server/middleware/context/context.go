package context

import (
	"context"

	"go.uber.org/zap"
)

type key int

const (
	loggerKey key = iota + 1
)

// SetLogger is store the zap logger in the context.
func SetLogger(ctx context.Context, logger *zap.Logger) context.Context {
	return context.WithValue(ctx, loggerKey, logger)
}
