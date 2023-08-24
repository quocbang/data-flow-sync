package account

import (
	"context"
	"time"

	"bou.ke/monkey"
	"github.com/quocbang/data-flow-sync/server/internal/repositories"
	"github.com/quocbang/data-flow-sync/server/internal/repositories/services/account"
)

func (s *Suite) TestSendMail() {
	assert := s.Assertions
	ctx := context.Background()

	// Arrange
	_, err := s.GetDm().Account().SignUp(ctx, repositories.SignUpAccountRequest{
		CreateAccountRequest: repositories.CreateAccountRequest{
			UserID:   "james",
			Email:    "mori@kenda.com.tw",
			Password: "test_password",
		},
	})
	assert.NoError(err)

	// Act
	err = s.GetDm().Account().SendMail(ctx, repositories.SendMailRequest{
		Email: "mori@kenda.com.tw",
	})

	// Assert
	assert.NoError(err)
	otp, err := s.GetRedis().Get(ctx, "mori@kenda.com.tw").Result()
	assert.NoError(err)
	assert.NotEqual("", otp)
}

func (s *Suite) TestSignUp() {
	assert := s.Assertions
	ctx := context.Background()

	// Arrange
	// Act
	_, err := s.GetDm().Account().SignUp(ctx, repositories.SignUpAccountRequest{
		CreateAccountRequest: repositories.CreateAccountRequest{
			UserID:   "james",
			Email:    "mori@kenda.com.tw",
			Password: "test_password",
		},
	})
	assert.NoError(err)

	// Assert
	_, err = s.GetDm().Account().SignIn(ctx, repositories.SignInRequest{
		Identifier: "james",
		Password:   "test_password",
		Options: repositories.Option{
			TokenLifeTime: 2 * time.Minute,
		},
	})
	assert.NoError(err)

}

func (s *Suite) TestVerifyAccount() {
	assert := s.Assertions
	ctx := context.Background()
	p := monkey.Patch(account.OptCreator, func() string {
		return "111111"
	})
	defer p.Unpatch()

	// Arrange
	_, err := s.GetDm().Account().SignUp(ctx, repositories.SignUpAccountRequest{
		CreateAccountRequest: repositories.CreateAccountRequest{
			UserID:   "james",
			Email:    "mori@kenda.com.tw",
			Password: "test_password",
		},
	})
	assert.NoError(err)

	// Act
	_, err = s.GetDm().Account().VerifyAccount(ctx, repositories.VerifyAccountRequest{
		Otp:   "111111",
		Email: "mori@kenda.com.tw",
		Option: repositories.Option{
			TokenLifeTime: 2 * time.Minute,
		},
	})

	// Assert
	assert.NoError(err)
}

// TODO: wait SignUp method
func (s *Suite) TestSignIn() {

	// login successfully
	{
		// Arrange

		// Act

		// Assert
	}
}

// TODO: wait SignIn method
func (s *Suite) TestSignOut() {

}

// TODO: wait SignIn method
func (s *Suite) TestAuthorization() {

}

func (s *Suite) TestDeleteAccount() {

}
