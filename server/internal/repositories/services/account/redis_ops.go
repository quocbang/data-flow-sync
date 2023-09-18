package account

import (
	"context"
	"time"
)

func (s service) AddOTP(ctx context.Context, email string, otp string) error {
	_, err := s.rd.Set(ctx, email, otp, 5*time.Minute).Result()
	if err != nil {
		return err
	}
	return nil
}

func (s service) DelOTP(ctx context.Context, email string) error {
	return s.rd.Del(ctx, email).Err()
}
