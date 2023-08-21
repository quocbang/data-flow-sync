package account

import (
	"context"
	"time"

	"github.com/quocbang/data-flow-sync/server/internal/repositories"
)

func (s *Suite) TestSendMail() {
	assert := s.Assertions
	ctx := context.Background()
	err := s.GetDm().Account().SendMail(ctx, repositories.SendMailRequest{
		UserID: "thai",
		Email:  "1194030275.dnu@gmail.com",
	})
	assert.NoError(err)
	otp, err := s.GetRedis().Get(ctx, "thai").Result()
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
			Email:    "1194030275.dnu@gmail.com",
			Password: "test_password",
		},
	})
	assert.NoError(err)

	// Assert
	_, err = s.GetDm().Account().SignIn(ctx, repositories.SignInRequest{
		UserID:   "james",
		Password: "test_password",
		Options: repositories.Option{
			TokenLifeTime: 2 * time.Minute,
		},
	})
	assert.NoError(err)

}

func (s *Suite) TestUpdateRole() {
	// assert := s.Assertions
	// ctx := context.Background()

	// // Arrange
	// s.GetDm().Account().SignUp()
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
