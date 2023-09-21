package redisconn

import (
	"context"
	"time"

	redis_db "github.com/quocbang/data-flow-sync/server/utils/redis_database"
	"github.com/redis/go-redis/v9"
)

type RDConn struct {
	rd *redis.Client
}

func (r RDConn) AddOTP(ctx context.Context, email string, otp string) error {
	// get to the desire database
	r.rd.Conn().Select(ctx, int(redis_db.Redis_DB_OTP))

	// add otp to database
	_, err := r.rd.Set(ctx, email, otp, 5*time.Minute).Result()
	if err != nil {
		return err
	}
	return nil
}

func (r RDConn) DelOTP(ctx context.Context, email string) error {
	// get to the desire database
	r.rd.Conn().Select(ctx, int(redis_db.Redis_DB_OTP))
	return r.rd.Del(ctx, email).Err()
}

func (r RDConn) GetOTPByEmail(ctx context.Context, email string) (string, error) {
	// get to the desire database
	r.rd.Conn().Select(ctx, int(redis_db.Redis_DB_OTP))
	value, err := r.rd.Get(ctx, email).Result()
	if err != nil {
		return "", err
	}
	return value, nil
}

// GetBlackList get data in black list.
func (r RDConn) GetBlackList(ctx context.Context, token string) (int64, error) { // return row number and error.
	// get to the desire database
	r.rd.Conn().Select(ctx, int(redis_db.Redis_DB_BLACKLIST))

	// check if token is in black list
	return r.rd.Exists(ctx, token).Result()
}

// AddToBackList add token to black list
func (r RDConn) AddToBackList(ctx context.Context, token string, expTime time.Duration) error {
	// get to the desire database
	r.rd.Conn().Select(ctx, int(redis_db.Redis_DB_BLACKLIST))

	// add token to black list
	return r.rd.Set(ctx, token, nil, expTime).Err()
}

func (r RDConn) Close() error {
	// close redis
	if err := r.rd.Close(); err != nil {
		return err
	}
	return nil
}
