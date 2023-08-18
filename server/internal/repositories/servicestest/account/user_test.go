package account

import (
	"context"
)

func (s *Suite) TestSendMail() {
	assert := s.Assertions
	ctx := context.Background()
	err := s.GetDm().Account().SendMail(ctx, "1194030275.dnu@gmail.com")
	assert.NoError(err)
}
