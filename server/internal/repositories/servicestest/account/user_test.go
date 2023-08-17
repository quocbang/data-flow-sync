package account

import (
	"context"
)

func (s *Suite) TestSendMail() {
	assert := s.Assertions
	ctx := context.Background()
	err := s.GetDm().Account().SendMail(ctx, "bangquoc9@gmail.com")
	assert.NoError(err)
}
