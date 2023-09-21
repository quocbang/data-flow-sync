package account

import (
	"context"
	"fmt"
	"net/http"
	"reflect"
	"time"

	"bou.ke/monkey"
	"github.com/golang-jwt/jwt"
	"github.com/stretchr/testify/mock"
	"golang.org/x/crypto/bcrypt"

	"github.com/quocbang/data-flow-sync/server/internal/mocks"
	"github.com/quocbang/data-flow-sync/server/internal/repositories"
	repoErrors "github.com/quocbang/data-flow-sync/server/internal/repositories/errors"
	m "github.com/quocbang/data-flow-sync/server/internal/repositories/orm/models"
	a "github.com/quocbang/data-flow-sync/server/internal/services/account"
	"github.com/quocbang/data-flow-sync/server/internal/servicestest/internal/setupmock"
	"github.com/quocbang/data-flow-sync/server/swagger/restapi/operations/account"
)

func (s *Suite) TestLogin() {
	assertion := s.Assertions
	username := "test_user"
	password := "test_password"
	params := func() account.LoginParams {
		return account.LoginParams{
			HTTPRequest: s.HttpTestRequest(http.MethodPost, "/api/user/login", nil),
			Login: account.LoginBody{
				Username: &username,
				Password: &password,
			},
		}
	}
	s.Context = params().HTTPRequest.Context()
	s.Context = context.WithValue(s.Context, a.SecretAccessKey, "")

	{ // login successfully
		// Arrange
		// mock repository
		goodParams := params()
		mockRepo := s.MockRepository()                                // repositories mock struct
		mockRepo.EXPECT().Account().ReturnArguments = mock.Arguments{ // service Account
			func() repositories.AccountServices { // func return AccountService interface that has multiple methods,
				account := mocks.AccountServices{} // and each method has mocks.AccountServices{} struct
				account.EXPECT().GetAccount(s.Context, "test_user").ReturnArguments = mock.Arguments{
					m.Account{}, nil,
				}
				return &account
			}(),
		}

		// patch compare password and generate jwt
		monkey.Patch(bcrypt.CompareHashAndPassword, func([]byte, []byte) error {
			return nil
		})
		monkey.PatchInstanceMethod(reflect.TypeOf(m.Account{}), "GenerateJWT", func(m.Account, context.Context, time.Duration, string) (string, error) {
			return "token_for_test", nil
		})
		defer monkey.UnpatchAll()

		mockServer := s.NewMockServer(setupmock.WithMockRepositories(&mockRepo.Mock)) // return mock Repositories struct for API service interface

		// Act
		response := mockServer.Account.Login(goodParams).(*account.LoginOK)

		// Assert
		assertion.Equal("token_for_test", response.Payload.Token)
	}
	{ // Internal Error: wrong password
		// Arrange
		wrongPassword := "wrong_password"
		wrongPasswordParams := params()
		wrongPasswordParams.Login.Password = &wrongPassword
		mockRepo := s.MockRepository()                                // repositories mock struct
		mockRepo.EXPECT().Account().ReturnArguments = mock.Arguments{ // service Account
			func() repositories.AccountServices { // func return AccountService interface that has multiple methods,
				account := mocks.AccountServices{} // and each method has mocks.AccountServices{} struct
				account.EXPECT().GetAccount(s.Context, "test_user").ReturnArguments = mock.Arguments{
					m.Account{}, nil,
				}
				return &account
			}(),
		}

		// patch compare password and generate jwt
		monkey.Patch(bcrypt.CompareHashAndPassword, func([]byte, []byte) error {
			return bcrypt.ErrMismatchedHashAndPassword
		})

		defer monkey.UnpatchAll()

		mockServer := s.NewMockServer(setupmock.WithMockRepositories(&mockRepo.Mock))

		// Act
		response := mockServer.Account.Login(wrongPasswordParams).(*account.LoginDefault)

		// Assert
		assertion.Equal(int64(repoErrors.Code_WRONG_PASSWORD), response.Payload.Code)
		assertion.Equal("wrong password", response.Payload.Details)
	}
	{ // failed to generate jwt
		// Arrange
		// mock repository
		mockRepo := s.MockRepository()                                // repositories mock struct
		mockRepo.EXPECT().Account().ReturnArguments = mock.Arguments{ // service Account
			func() repositories.AccountServices { // func return AccountService interface that has multiple methods,
				account := mocks.AccountServices{} // and each method has mocks.AccountServices{} struct
				account.EXPECT().GetAccount(s.Context, "test_user").ReturnArguments = mock.Arguments{
					m.Account{}, nil,
				}
				return &account
			}(),
		}

		// patch compare password and generate jwt
		monkey.Patch(bcrypt.CompareHashAndPassword, func([]byte, []byte) error {
			return nil
		})
		monkey.PatchInstanceMethod(reflect.TypeOf(m.Account{}), "GenerateJWT", func(m.Account, context.Context, time.Duration, string) (string, error) {
			return "", fmt.Errorf("failed to generate token")
		})
		defer monkey.UnpatchAll()

		mockServer := s.NewMockServer(setupmock.WithMockRepositories(&mockRepo.Mock))

		// Act
		response := mockServer.Account.Login(params()).(*account.LoginDefault)

		// Assert
		assertion.Equal(int64(0), response.Payload.Code)
		assertion.Equal("failed to generate token", response.Payload.Details)
	}
	{ // failed to get account
		// Arrange
		// mock repository
		mockRepo := s.MockRepository()                                // repositories mock struct
		mockRepo.EXPECT().Account().ReturnArguments = mock.Arguments{ // service Account
			func() repositories.AccountServices { // func return AccountService interface that has multiple methods,
				account := mocks.AccountServices{} // and each method has mocks.AccountServices{} struct
				account.EXPECT().GetAccount(s.Context, "test_user").ReturnArguments = mock.Arguments{
					m.Account{}, repoErrors.Error{
						Code:    0,
						Details: "account not found",
					},
				}
				return &account
			}(),
		}

		mockServer := s.NewMockServer(setupmock.WithMockRepositories(&mockRepo.Mock))

		// Act
		response := mockServer.Account.Login(params()).(*account.LoginDefault)

		// Assert
		assertion.Equal(int64(0), response.Payload.Code)
		assertion.Equal("account not found", response.Payload.Details)
	}
}

func (s *Suite) TestLogOut() {
	assertion := s.Assertions

	params := func() account.LogoutParams {
		return account.LogoutParams{
			HTTPRequest: s.HttpTestRequest(http.MethodPost, "/api/user/logout", nil),
		}
	}

	s.Context = params().HTTPRequest.Context()
	s.Context = context.WithValue(s.Context, a.SecretAccessKey, "")

	{ // logout successfully
		// Arrange
		// patch VerifyToken
		p := monkey.Patch(m.VerifyToken, func(string, string) (*m.JwtCustomClaims, error) {
			return &m.JwtCustomClaims{
				StandardClaims: jwt.StandardClaims{
					ExpiresAt: time.Now().Unix(),
				},
			}, nil
		})
		defer p.Unpatch()

		// mock redis
		goodParams := params()
		goodParams.HTTPRequest.Header.Set(string(a.AuthorizationKey), "token_for_tester")
		mockRedis := s.MockRedis()
		mockRedis.EXPECT().AddToBackList(s.Context, "token_for_tester", time.Until(time.Unix(time.Now().Unix(), 0))).
			ReturnArguments = mock.Arguments{nil}

		mockServer := s.NewMockServer(setupmock.WithMockRedisServer(&mockRedis.Mock))

		// Act
		response := mockServer.Account.Logout(goodParams, nil)

		// Assert
		_, ok := response.(*account.LogoutOK)
		assertion.True(ok)
	}
	// bad cases
	{ // failed to verify token
		// Arrange
		// patch VerifyToken
		p := monkey.Patch(m.VerifyToken, func(string, string) (*m.JwtCustomClaims, error) {
			return &m.JwtCustomClaims{}, fmt.Errorf("failed to verify token")
		})
		defer p.Unpatch()

		mockServer := s.NewMockServer()

		// Act
		response := mockServer.Account.Logout(params(), nil).(*account.LogoutDefault)

		// Assert
		assertion.Equal(int64(0), response.Payload.Code)
		assertion.Equal("failed to verify token", response.Payload.Details)
	}
	{ // failed to add token in black list
		// Arrange
		// patch VerifyToken
		p := monkey.Patch(m.VerifyToken, func(string, string) (*m.JwtCustomClaims, error) {
			return &m.JwtCustomClaims{
				StandardClaims: jwt.StandardClaims{
					ExpiresAt: time.Now().Unix(),
				},
			}, nil
		})
		defer p.Unpatch()

		// mock redis
		goodParams := params()
		goodParams.HTTPRequest.Header.Set(string(a.AuthorizationKey), "token_for_tester")
		mockRedis := s.MockRedis()
		mockRedis.EXPECT().AddToBackList(s.Context, "token_for_tester", time.Until(time.Unix(time.Now().Unix(), 0))).
			ReturnArguments = mock.Arguments{fmt.Errorf("failed to add token to black list")}

		mockServer := s.NewMockServer(setupmock.WithMockRedisServer(&mockRedis.Mock))

		// Act
		response := mockServer.Account.Logout(goodParams, nil).(*account.LogoutDefault)

		// Assert
		assertion.Equal(int64(0), response.Payload.Code)
		assertion.Equal("failed to add token to black list", response.Payload.Details)
	}
}
