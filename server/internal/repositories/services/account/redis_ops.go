package account

import (
	"context"
	"time"
)

func (s service) addOTP(ctx context.Context, userID string, otp string) error {
	_, err := s.rd.Set(ctx, userID, otp, 5*time.Minute).Result()
	if err != nil {
		return err
	}
	return nil
}
