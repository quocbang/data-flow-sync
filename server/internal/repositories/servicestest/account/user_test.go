package account

import (
	"context"
	"reflect"
	"time"

	"github.com/redis/go-redis/v9"

	"bou.ke/monkey"
	"github.com/quocbang/data-flow-sync/server/internal/repositories"
	cusErr "github.com/quocbang/data-flow-sync/server/internal/repositories/errors"
)

func (s *Suite) TestGetAccount() {
	assert := s.Assertions
	ctx := context.Background()

	// good cases
	{ // get user by name
		// Arrange
		_, err := s.GetDm().Account().SignUp(ctx, repositories.SignUpAccountRequest{
			CreateAccountRequest: repositories.CreateAccountRequest{
				UserID: "james",
				Email:  "james@gmail.com",
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
		acc, err := s.GetDm().Account().GetAccount(ctx, "james@gmail.com")

		// Assert
		assert.NoError(err)
		assert.Equal(int32(acc.Roles), int32(0))
	}
	// bad case
	{ // not found
		// Arrange

		// Act
		_, err := s.GetDm().Account().GetAccount(ctx, "noexist@gmail.com")

		// Assert
		assert.Error(err)
		expect := cusErr.Error{
			Code:    cusErr.Code_ACCOUNT_NOT_FOUND,
			Details: "account not found",
		}
		assert.Equal(expect.Error(), err.Error())
	}

}

func (s *Suite) TestSignUp() {
	assert := s.Assertions
	ctx := context.Background()

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
	monkey.PatchInstanceMethod(reflect.TypeOf(&redis.StringCmd{}), "Result", func(*redis.StringCmd) (string, error) {
		return "111111", nil
	})

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
