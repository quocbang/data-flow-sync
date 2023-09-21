package redis_conn

import (
	"context"
	"time"
)

type RedisConn interface {
	Close() error
	AddOTP(context.Context, string, string) error
	DelOTP(context.Context, string) error
	GetOTPByEmail(context.Context, string) (string, error)
	GetBlackList(context.Context, string) (int64, error)
	AddToBackList(context.Context, string, time.Duration) error
}
