package account

import (
	"context"

	"github.com/quocbang/data-flow-sync/server/internal/repositories"
	cusErr "github.com/quocbang/data-flow-sync/server/internal/repositories/errors"
)

func (s *Suite) TestGetAccount() {
	assert := s.Assertions
	ctx := context.Background()

	// good cases
	{ // get user by name
		// Arrange
		err := s.GetDm().Account().SignUp(ctx, repositories.SignUpAccountRequest{
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
	err := s.GetDm().Account().SignUp(ctx, repositories.SignUpAccountRequest{
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
