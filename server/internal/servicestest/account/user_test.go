package account

import (
	"net/http"

	apiErrors "github.com/go-openapi/errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/quocbang/data-flow-sync/server/internal/mocks"
	"github.com/quocbang/data-flow-sync/server/internal/repositories"
	repoErrors "github.com/quocbang/data-flow-sync/server/internal/repositories/errors"
	m "github.com/quocbang/data-flow-sync/server/internal/repositories/orm/models"
	"github.com/quocbang/data-flow-sync/server/internal/servicestest/internal/setupmock"
	"github.com/quocbang/data-flow-sync/server/swagger/models"
	"github.com/quocbang/data-flow-sync/server/swagger/restapi/operations/account"
	"github.com/quocbang/data-flow-sync/server/utils/roles"
)

func (s *Suite) TestLogin() {
	assertion := assert.New(s.T())
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

	{ // login successfully
		// Arrange
		goodParams := params()
		goodMockRepo := s.MockRepository()                                        // repositories mock struct
		goodMockRepo.EXPECT().Account().Maybe().ReturnArguments = mock.Arguments{ // service Account
			func() repositories.AccountServices { // func return AccountService interface that has multiple methods,
				account := mocks.AccountServices{} // and each method has mocks.AccountServices{} struct
				account.EXPECT().SignIn(s.Context, repositories.SignInRequest{
					UserID:   username,
					Password: password,
				}).ReturnArguments = mock.Arguments{
					repositories.SignInReply{Token: "token_for_tester"},
					nil,
				}
				return &account
			}(),
		}
		mockServer := s.NewMockServer(&goodMockRepo.Mock, setupmock.MockServerOptions{}) // return mock Repositories struct for API service interface

		// Act
		response := mockServer.Account.Login(goodParams).(*account.LoginOK)

		// Assert
		assertion.Equal("token_for_tester", response.Payload.Token)
	}
	{ // Internal Error: wrong password
		// Arrange
		wrongPassword := "wrong_password"
		wrongPasswordParams := params()
		wrongPasswordParams.Login.Password = &wrongPassword
		badMockRepo := s.MockRepository()
		badMockRepo.EXPECT().Account().ReturnArguments = mock.Arguments{
			func() repositories.AccountServices {
				account := mocks.AccountServices{}
				account.EXPECT().SignIn(s.Context, repositories.SignInRequest{
					UserID:   username,
					Password: wrongPassword,
				}).ReturnArguments = mock.Arguments{
					repositories.SignInReply{},
					repoErrors.Error{
						Code:    repoErrors.Code_WRONG_PASSWORD,
						Details: "wrong password",
					},
				}
				return &account
			}(),
		}
		mockServer := s.NewMockServer(&badMockRepo.Mock, setupmock.MockServerOptions{})

		// Act
		response := mockServer.Account.Login(wrongPasswordParams).(*account.LoginDefault)

		// Assert
		assertion.Equal(int64(repoErrors.Code_WRONG_PASSWORD), response.Payload.Code)
		assertion.Equal("wrong password", response.Payload.Details)
	}
}

func (s *Suite) TestLogOut() {
	assertion := assert.New(s.T())

	params := func() account.LogoutParams {
		return account.LogoutParams{
			HTTPRequest: s.HttpTestRequest(http.MethodPost, "/api/user/logout", nil),
		}
	}

	{ // logout successfully
		// Arrange
		goodParams := params()
		goodParams.HTTPRequest.Header.Set("x-data-flow-sync-auth-key", "token_for_tester")
		mockRepo := s.MockRepository()
		mockRepo.EXPECT().Account().ReturnArguments = mock.Arguments{
			func() repositories.AccountServices {
				account := mocks.AccountServices{}
				account.EXPECT().SignOut(s.Context, "token_for_tester").ReturnArguments = mock.Arguments{
					nil,
				}
				return &account
			}(),
		}
		mockServer := s.NewMockServer(&mockRepo.Mock, setupmock.MockServerOptions{})

		// Act
		response := mockServer.Account.Logout(goodParams, nil)

		// Assert
		_, ok := response.(*account.LogoutOK)
		assertion.True(ok)
	}
}

func (s *Suite) TestAuth() {
	assertion := assert.New(s.T())
	token := "token_for_tester"

	{ // authorized
		// Arrange
		mockRepo := s.MockRepository()
		mockRepo.EXPECT().Account().ReturnArguments = mock.Arguments{
			func() repositories.AccountServices {
				account := mocks.AccountServices{}
				account.EXPECT().Authorization(s.Context, token).ReturnArguments = mock.Arguments{
					&m.JwtCustomClaims{
						UserID:            "tester",
						Role:              roles.Roles_UNSPECIFIED,
						IsUnspecifiedUser: true,
					},
					nil,
				}
				return &account
			}(),
		}
		mockServer := s.NewMockServer(&mockRepo.Mock, setupmock.MockServerOptions{})

		// Act
		principal, err := mockServer.Account.Auth(token)

		// Assert
		assertion.NoError(err)
		assertion.Equal(&models.Principal{
			ID:                "tester",
			IsUnspecifiedUser: true,
			Role:              int64(roles.Roles_UNSPECIFIED),
		}, principal)
	}
	{ // token blocked
		// Arrange
		mockRepo := s.MockRepository()
		mockRepo.EXPECT().Account().ReturnArguments = mock.Arguments{
			func() repositories.AccountServices {
				account := mocks.AccountServices{}
				account.EXPECT().Authorization(s.Context, token).ReturnArguments = mock.Arguments{
					nil,
					repoErrors.Error{
						Code:    repoErrors.Code_TOKEN_BLOCKED,
						Details: "token was blocked",
					},
				}
				return &account
			}(),
		}
		mockServer := s.NewMockServer(&mockRepo.Mock, setupmock.MockServerOptions{})

		// Act
		_, err := mockServer.Account.Auth(token)

		// Assert
		assertion.Error(err)
		assertion.Equal(apiErrors.New(http.StatusUnauthorized, repoErrors.Error{
			Code:    repoErrors.Code_TOKEN_BLOCKED,
			Details: "token was blocked",
		}.Error()), err)
	}
}
