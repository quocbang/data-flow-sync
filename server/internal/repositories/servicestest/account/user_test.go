package account

import (
	"context"
	"testing"
)

func (s Suite) Test_SendMail(t *testing.T) {
	ctx := context.Background()
	s.GetDm().Account().SendMail(ctx, "leducthai.access.vcmi@gmail.com")
}
