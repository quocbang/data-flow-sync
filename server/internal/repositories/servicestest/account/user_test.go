package account

import (
	"context"
	"fmt"
	"io"
	"time"

	"bou.ke/monkey"
	"github.com/quocbang/data-flow-sync/server/internal/repositories"
	cusErr "github.com/quocbang/data-flow-sync/server/internal/repositories/errors"
	"github.com/quocbang/data-flow-sync/server/internal/repositories/services/account"
)

func (s *Suite) TestGetAccount() {
	assert := s.Assertions
	ctx := context.Background()
	// patch send mail method to reduce the real mail sending
	monkey.Patch(fmt.Fprintln, func(w io.Writer, a ...any) (int, error) {
		return 0, nil
	})
	monkey.Patch(account.OptCreator, func() string {
		return "111111"
	})
	defer monkey.UnpatchAll()

	// good cases
	{ // get user by name
		// Arrange
		_, err := s.GetDm().Account().SignUp(ctx, repositories.SignUpAccountRequest{
			CreateAccountRequest: repositories.CreateAccountRequest{
				UserID: "james",
				Email:  "james@gmai.com",
			},
		})
		assert.NoError(err)

		// Act
		acc, err := s.GetDm().Account().GetAccount(ctx, "james")

		// Assert
		assert.NoError(err)
		assert.Equal(int32(acc.Roles), int32(0))
	}
	{ // get user by email
		// Arrange

		// Act
		acc, err := s.GetDm().Account().GetAccount(ctx, "james@gmai.com")

		// Assert
		assert.NoError(err)
		assert.Equal(int32(acc.Roles), int32(0))
	}
	// bad case
	{ // not found
		// Arrange

		// Act
		_, err := s.GetDm().Account().GetAccount(ctx, "noexist@gmai.com")

		// Assert
		assert.Error(err)
		expect := cusErr.Error{
			Code:    cusErr.Code_ACCOUNT_NOT_FOUND,
			Details: "account not found",
		}
		assert.Equal(expect.Error(), err.Error())
	}

}

func (s *Suite) TestSendMail() {
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
	err = s.GetDm().Account().SendMail(ctx, repositories.SendMailRequest{
		Email: "mori@kenda.com.tw",
	})

	// Assert
	assert.NoError(err)
	otp, err := s.GetRedis().Get(ctx, "mori@kenda.com.tw").Result()
	assert.NoError(err)
	assert.Equal("111111", otp)
}

func (s *Suite) TestSignUp() {
	assert := s.Assertions
	ctx := context.Background()
	// patch send mail method to reduce the real mail sending
	monkey.Patch(fmt.Fprintln, func(w io.Writer, a ...any) (int, error) {
		return 0, nil
	})
	monkey.Patch(account.OptCreator, func() string {
		return "111111"
	})
	defer monkey.UnpatchAll()

	// Arrange
	// Act
	_, err := s.GetDm().Account().SignUp(ctx, repositories.SignUpAccountRequest{
		CreateAccountRequest: repositories.CreateAccountRequest{
			UserID:   "james",
			Email:    "james@gmail.com",
			Password: "test_password",
		},
	})
	assert.NoError(err)

	// Assert
	newAccount, err := s.GetDm().Account().GetAccount(ctx, "james")
	assert.NoError(err)
	assert.Equal(int32(newAccount.Roles), int32(0))

}

func (s *Suite) TestVerifyAccount() {
	assert := s.Assertions
	ctx := context.Background()
	// patch send mail method to reduce the real mail sending
	monkey.Patch(fmt.Fprintln, func(w io.Writer, a ...any) (int, error) {
		return 0, nil
	})
	monkey.Patch(account.OptCreator, func() string {
		return "111111"
	})
	defer monkey.UnpatchAll()

	// Arrange
	_, err := s.GetDm().Account().SignUp(ctx, repositories.SignUpAccountRequest{
		CreateAccountRequest: repositories.CreateAccountRequest{
			UserID:   "james",
			Email:    "james@gmail.com",
			Password: "test_password",
		},
	})
	assert.NoError(err)

	// Act
	_, err = s.GetDm().Account().VerifyAccount(ctx, repositories.VerifyAccountRequest{
		Otp:   "111111",
		Email: "james@gmail.com",
		Option: repositories.Option{
			TokenLifeTime: 2 * time.Minute,
		},
	})

	// Assert
	assert.NoError(err)
	newAccount, err := s.GetDm().Account().GetAccount(ctx, "james")
	assert.NoError(err)
	assert.Equal(int32(newAccount.Roles), int32(1))
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
