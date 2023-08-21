package account

import (
	"context"
	"time"
)

func (s service) addOTP(ctx context.Context, userID string, otp string) error {
	otpExpTime := 2 * time.Minute
	err := s.rd.Set(ctx, userID, otp, otpExpTime).Err()
	if err != nil {
		return err
	}
	return nil
}
