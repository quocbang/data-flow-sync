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

// func (s *Suite) TestAuth() {
// 	assertion := s.Assertions
// 	token := "token_for_tester"
// 	s.Context = context.WithValue(s.Context, a.SecretAccessKey, "")

// 	{ // authorized
// 		// Arrange

// 		mockServer := s.NewMockServer(setupmock.WithMockRepositories(&mockRepo.Mock))

// 		// Act
// 		principal, err := mockServer.Account.Auth(token)

// 		// Assert
// 		assertion.NoError(err)
// 		assertion.Equal(&models.Principal{
// 			ID:                "tester",
// 			IsUnspecifiedUser: true,
// 			Role:              int64(roles.Roles_UNSPECIFIED),
// 		}, principal)
// 	}
// 	{ // token blocked
// 		// Arrange
// 		mockRepo := s.MockRepository()
// 		mockRepo.EXPECT().Account().ReturnArguments = mock.Arguments{
// 			func() repositories.AccountServices {
// 				account := mocks.AccountServices{}
// 				account.EXPECT().Authorization(s.Context, token).ReturnArguments = mock.Arguments{
// 					nil,
// 					repoErrors.Error{
// 						Code:    repoErrors.Code_TOKEN_BLOCKED,
// 						Details: "token was blocked",
// 					},
// 				}
// 				return &account
// 			}(),
// 		}
// 		mockServer := s.NewMockServer(setupmock.WithMockRepositories(&mockRepo.Mock))

// 		// Act
// 		_, err := mockServer.Account.Auth(token)

// 		// Assert
// 		assertion.Error(err)
// 		assertion.Equal(apiErrors.New(http.StatusUnauthorized, repoErrors.Error{
// 			Code:    repoErrors.Code_TOKEN_BLOCKED,
// 			Details: "token was blocked",
// 		}.Error()), err)
// 	}
// }

// func (s *Suite) TestSendMail() {
// 	s.Context = context.Background()
// 	assertions := s.Assertions
// 	testUserID := "james"
// 	testEmail := "james@gmail.com"
// 	params := func() account.SendMailParams {
// 		return account.SendMailParams{
// 			HTTPRequest: httptest.NewRequest(http.MethodPost, "http://example.com/api/user/send-mail", nil),
// 		}
// 	}

// 	{ // good case
// 		// Arrange
// 		mockMailServer := s.MockMailServer()
// 		mockMailServer.EXPECT().SendAccountVerification(s.Context, mailserver.MailVerifyRequest{
// 			Recipient: "james@gmail.com",
// 		}).ReturnArguments = mock.Arguments{"111111", nil}

// 		mockRepo := s.MockRepository()
// 		mockRepo.EXPECT().Account().ReturnArguments = mock.Arguments{
// 			func() repositories.AccountServices {
// 				acc := mocks.AccountServices{}
// 				acc.EXPECT().AddOTP(s.Context, "james@gmail.com", "111111").ReturnArguments =
// 					mock.Arguments{nil}

// 				return &acc
// 			}(),
// 		}

// 		mockServer := s.NewMockServer(setupmock.WithMockMailServer(&mockMailServer.Mock), setupmock.WithMockRepositories(&mockRepo.Mock))

// 		// Act
// 		response := mockServer.Account.SendMail(params(), &models.Principal{
// 			Email:             testEmail,
// 			ID:                testUserID,
// 			IsUnspecifiedUser: true,
// 			Role:              0,
// 		})

// 		// Assert
// 		_, ok := response.(*account.SendMailOK)
// 		assertions.True(ok)
// 	}
// 	// bad cases
// 	{ // internal fail
// 		// Arrange
// 		mockMailServer := s.MockMailServer()
// 		mockMailServer.EXPECT().SendAccountVerification(s.Context, mailserver.MailVerifyRequest{
// 			Recipient: "james@gmail.com",
// 		}).ReturnArguments = mock.Arguments{"", repoErrors.Error{
// 			Code:    0,
// 			Details: "internal error",
// 		}}

// 		mockServer := s.NewMockServer(setupmock.WithMockMailServer(&mockMailServer.Mock))

// 		// Act
// 		response := mockServer.Account.SendMail(params(), &models.Principal{
// 			Email:             testEmail,
// 			ID:                testUserID,
// 			IsUnspecifiedUser: true,
// 			Role:              0,
// 		})

// 		// Assert
// 		res := suiteutils.NewHttpResponseWriter()
// 		cusProducer := runtime.ProducerFunc(suiteutils.MyProducer)
// 		response.WriteResponse(res, cusProducer)

// 		expect := []byte(`{"details":"internal error"}`)
// 		assertions.Equal(string(expect), res.Body.String())
// 	}
// 	{ // identified user
// 		// Arrange
// 		mockMailServer := s.MockMailServer()
// 		mockMailServer.EXPECT().SendAccountVerification(s.Context, mailserver.MailVerifyRequest{
// 			Recipient: "james@gmail.com",
// 		}).ReturnArguments = mock.Arguments{"", repoErrors.Error{
// 			Code:    0,
// 			Details: "internal error",
// 		}}

// 		mockServer := s.NewMockServer(setupmock.WithMockMailServer(&mockMailServer.Mock))

// 		// Act
// 		response := mockServer.Account.SendMail(params(), &models.Principal{
// 			Email:             testEmail,
// 			ID:                testUserID,
// 			IsUnspecifiedUser: false,
// 			Role:              1,
// 		})

// 		// Assert
// 		res := suiteutils.NewHttpResponseWriter()
// 		cusProducer := runtime.ProducerFunc(suiteutils.MyProducer)
// 		response.WriteResponse(res, cusProducer)

// 		expect := []byte(`{"details":"user been verified"}`)
// 		assertions.Equal(string(expect), res.Body.String())
// 	}
// 	{ // failed to add to temporary storage
// 		// Arrange
// 		mockMailServer := s.MockMailServer()
// 		mockMailServer.EXPECT().SendAccountVerification(s.Context, mailserver.MailVerifyRequest{
// 			Recipient: "james@gmail.com",
// 		}).ReturnArguments = mock.Arguments{"111111", nil}

// 		mockRepo := s.MockRepository()
// 		mockRepo.EXPECT().Account().ReturnArguments = mock.Arguments{
// 			func() repositories.AccountServices {
// 				acc := mocks.AccountServices{}
// 				acc.EXPECT().AddOTP(s.Context, "james@gmail.com", "111111").ReturnArguments =
// 					mock.Arguments{
// 						repoErrors.Error{
// 							Code:    0,
// 							Details: "redis error",
// 						},
// 					}

// 				return &acc
// 			}(),
// 		}

// 		mockServer := s.NewMockServer(setupmock.WithMockMailServer(&mockMailServer.Mock), setupmock.WithMockRepositories(&mockRepo.Mock))

// 		// Act
// 		response := mockServer.Account.SendMail(params(), &models.Principal{
// 			Email:             testEmail,
// 			ID:                testUserID,
// 			IsUnspecifiedUser: true,
// 			Role:              0,
// 		})

// 		// Assert
// 		res := suiteutils.NewHttpResponseWriter()
// 		cusProducer := runtime.ProducerFunc(suiteutils.MyProducer)
// 		response.WriteResponse(res, cusProducer)

// 		expect := []byte(`{"details":"redis error"}`)
// 		assertions.Equal(string(expect), res.Body.String())
// 	}
// }

// func (s *Suite) TestVerifyAccount() {
// 	assertions := s.Assertions
// 	testUserID := "james"
// 	testEmail := "james@gmail.com"
// 	testOtp := "111111"
// 	testToken := "test.token"
// 	params := func() account.VerifyAccountParams {
// 		return account.VerifyAccountParams{
// 			HTTPRequest: httptest.NewRequest(http.MethodPost, "http://example.com/api/user/verify-account", nil),
// 			AccountVerify: account.VerifyAccountBody{
// 				Otp: testOtp,
// 			},
// 		}
// 	}

// 	// good case
// 	{
// 		// Arrange
// 		mockRepo := s.MockRepository()
// 		mockRepo.EXPECT().Account().ReturnArguments = mock.Arguments{
// 			func() repositories.AccountServices {
// 				acc := mocks.AccountServices{}
// 				acc.EXPECT().VerifyAccount(s.Context, repositories.VerifyAccountRequest{
// 					Otp:   testOtp,
// 					Email: testEmail,
// 				}).ReturnArguments = mock.Arguments{
// 					repositories.VerifyAccountReply{
// 						Token: testToken,
// 					}, nil,
// 				}

// 				return &acc
// 			}(),
// 		}

// 		mockServer := s.NewMockServer(setupmock.WithMockRepositories(&mockRepo.Mock))

// 		// Act
// 		response := mockServer.Account.VerifyAccount(params(), &models.Principal{
// 			Email:             testEmail,
// 			ID:                testUserID,
// 			IsUnspecifiedUser: true,
// 			Role:              0,
// 		})

// 		// Assert
// 		res := suiteutils.NewHttpResponseWriter()
// 		cusProducer := runtime.ProducerFunc(suiteutils.MyProducer)
// 		response.WriteResponse(res, cusProducer)

// 		expect := []byte(`{"token":"test.token"}`)
// 		assertions.Equal(string(expect), res.Body.String())
// 	}
// 	// bad cases
// 	{ // internal error
// 		// Arrange
// 		mockRepo := s.MockRepository()
// 		mockRepo.EXPECT().Account().ReturnArguments = mock.Arguments{
// 			func() repositories.AccountServices {
// 				acc := mocks.AccountServices{}
// 				acc.EXPECT().VerifyAccount(s.Context, repositories.VerifyAccountRequest{
// 					Otp:   testOtp,
// 					Email: testEmail,
// 				}).ReturnArguments = mock.Arguments{
// 					repositories.VerifyAccountReply{}, repoErrors.Error{
// 						Code:    0,
// 						Details: "internal error",
// 					},
// 				}

// 				return &acc
// 			}(),
// 		}

// 		mockServer := s.NewMockServer(setupmock.WithMockRepositories(&mockRepo.Mock))

// 		// Act
// 		response := mockServer.Account.VerifyAccount(params(), &models.Principal{
// 			Email:             testEmail,
// 			ID:                testUserID,
// 			IsUnspecifiedUser: true,
// 			Role:              0,
// 		})

// 		// Assert
// 		res := suiteutils.NewHttpResponseWriter()
// 		cusProducer := runtime.ProducerFunc(suiteutils.MyProducer)
// 		response.WriteResponse(res, cusProducer)

// 		expect := []byte(`{"details":"internal error"}`)
// 		assertions.Equal(string(expect), res.Body.String())
// 	}
// 	{ // verified user
// 		// Arrange
// 		mockRepo := s.MockRepository()
// 		mockRepo.EXPECT().Account().ReturnArguments = mock.Arguments{
// 			func() repositories.AccountServices {
// 				acc := mocks.AccountServices{}
// 				acc.EXPECT().VerifyAccount(s.Context, repositories.VerifyAccountRequest{
// 					Otp:   testOtp,
// 					Email: testEmail,
// 				}).ReturnArguments = mock.Arguments{
// 					repositories.VerifyAccountReply{}, nil,
// 				}

// 				return &acc
// 			}(),
// 		}

// 		mockServer := s.NewMockServer(setupmock.WithMockRepositories(&mockRepo.Mock))

// 		// Act
// 		response := mockServer.Account.VerifyAccount(params(), &models.Principal{
// 			Email:             testEmail,
// 			ID:                testUserID,
// 			IsUnspecifiedUser: false,
// 			Role:              1,
// 		})

// 		// Assert
// 		res := suiteutils.NewHttpResponseWriter()
// 		cusProducer := runtime.ProducerFunc(suiteutils.MyProducer)
// 		response.WriteResponse(res, cusProducer)

// 		expect := []byte(`{"details":"user been verified"}`)
// 		assertions.Equal(string(expect), res.Body.String())
// 	}
// }

// func (s *Suite) TestSignUp() {
// 	assertions := s.Assertions
// 	testUserID := "james"
// 	testEmail := "james@gmail.com"
// 	testPassword := "test_password"
// 	testToken := "test.token"
// 	s.Context = context.WithValue(s.Context, a.SecretAccessKey, "")
// 	params := func() account.SignUpParams {
// 		return account.SignUpParams{
// 			HTTPRequest: httptest.NewRequest(http.MethodPost, "http://example.com/api/user/verify-account", nil),
// 			SignUp: account.SignUpBody{
// 				Email:    testEmail,
// 				Name:     testUserID,
// 				Password: testPassword,
// 			},
// 		}
// 	}

// 	{ // good case
// 		// Arrange
// 		mockTxRepo := s.MockRepository()
// 		mockTxRepo.EXPECT().Account().ReturnArguments = mock.Arguments{
// 			func() repositories.AccountServices {
// 				acc := mocks.AccountServices{}
// 				acc.EXPECT().SignUp(s.Context, repositories.SignUpAccountRequest{
// 					CreateAccountRequest: repositories.CreateAccountRequest{
// 						UserID:   testUserID,
// 						Email:    testEmail,
// 						Password: testPassword,
// 					},
// 				}).ReturnArguments = mock.Arguments{
// 					repositories.SignInReply{
// 						Token: testToken,
// 					}, nil,
// 				}
// 				acc.EXPECT().AddOTP(s.Context, "james@gmail.com", "111111").ReturnArguments =
// 					mock.Arguments{nil}

// 				return &acc
// 			}(),
// 		}
// 		mockTxRepo.EXPECT().Commit().ReturnArguments = mock.Arguments{
// 			nil,
// 		}
// 		mockTxRepo.EXPECT().RollBack().ReturnArguments = mock.Arguments{
// 			fmt.Errorf("not in transaction"),
// 		}

// 		mockRepo := s.MockRepository()
// 		mockRepo.EXPECT().Begin(s.Context).ReturnArguments = mock.Arguments{
// 			mockTxRepo,
// 			nil,
// 		}

// 		mockMailServer := s.MockMailServer()
// 		mockMailServer.EXPECT().SendAccountVerification(s.Context, mailserver.MailVerifyRequest{
// 			Recipient: "james@gmail.com",
// 		}).ReturnArguments = mock.Arguments{"111111", nil}

// 		mockServer := s.NewMockServer(setupmock.WithMockRepositories(&mockRepo.Mock), setupmock.WithMockMailServer(&mockMailServer.Mock))

// 		// Act
// 		response := mockServer.Account.SignUp(params())

// 		// Assert
// 		res := suiteutils.NewHttpResponseWriter()
// 		cusProducer := runtime.ProducerFunc(suiteutils.MyProducer)
// 		response.WriteResponse(res, cusProducer)

// 		expect := []byte(`{"token":"test.token"}`)
// 		assertions.Equal(string(expect), res.Body.String())
// 	}
// 	// bad cases
// 	{ // sign up failed
// 		// Arrange
// 		mockTxRepo := s.MockRepository()
// 		mockTxRepo.EXPECT().Account().ReturnArguments = mock.Arguments{
// 			func() repositories.AccountServices {
// 				acc := mocks.AccountServices{}
// 				acc.EXPECT().SignUp(s.Context, repositories.SignUpAccountRequest{
// 					CreateAccountRequest: repositories.CreateAccountRequest{
// 						UserID:   testUserID,
// 						Email:    testEmail,
// 						Password: testPassword,
// 					},
// 				}).ReturnArguments = mock.Arguments{
// 					repositories.SignInReply{}, repoErrors.Error{
// 						Code:    0,
// 						Details: "internal error",
// 					},
// 				}
// 				acc.EXPECT().DelOTP(s.Context, "james@gmail.com").ReturnArguments =
// 					mock.Arguments{nil}

// 				return &acc
// 			}(),
// 		}
// 		mockTxRepo.EXPECT().RollBack().ReturnArguments = mock.Arguments{
// 			nil,
// 		}

// 		mockRepo := s.MockRepository()
// 		mockRepo.EXPECT().Begin(s.Context).ReturnArguments = mock.Arguments{
// 			mockTxRepo,
// 			nil,
// 		}

// 		mockServer := s.NewMockServer(setupmock.WithMockRepositories(&mockRepo.Mock))

// 		// Act
// 		response := mockServer.Account.SignUp(params())

// 		// Assert
// 		res := suiteutils.NewHttpResponseWriter()
// 		cusProducer := runtime.ProducerFunc(suiteutils.MyProducer)
// 		response.WriteResponse(res, cusProducer)

// 		expect := []byte(`{"details":"internal error"}`)
// 		assertions.Equal(string(expect), res.Body.String())
// 	}
// 	{ // send mail failed
// 		// Arrange
// 		mockTxRepo := s.MockRepository()
// 		mockTxRepo.EXPECT().Account().ReturnArguments = mock.Arguments{
// 			func() repositories.AccountServices {
// 				acc := mocks.AccountServices{}
// 				acc.EXPECT().SignUp(s.Context, repositories.SignUpAccountRequest{
// 					CreateAccountRequest: repositories.CreateAccountRequest{
// 						UserID:   testUserID,
// 						Email:    testEmail,
// 						Password: testPassword,
// 					},
// 				}).ReturnArguments = mock.Arguments{
// 					repositories.SignInReply{
// 						Token: testToken,
// 					}, nil,
// 				}
// 				acc.EXPECT().DelOTP(s.Context, "james@gmail.com").ReturnArguments =
// 					mock.Arguments{nil}

// 				return &acc
// 			}(),
// 		}
// 		mockTxRepo.EXPECT().RollBack().ReturnArguments = mock.Arguments{
// 			nil,
// 		}

// 		mockMailServer := s.MockMailServer()
// 		mockMailServer.EXPECT().SendAccountVerification(s.Context, mailserver.MailVerifyRequest{
// 			Recipient: "james@gmail.com",
// 		}).ReturnArguments = mock.Arguments{"", repoErrors.Error{
// 			Code:    0,
// 			Details: "send mail verification failed",
// 		}}

// 		mockRepo := s.MockRepository()
// 		mockRepo.EXPECT().Begin(s.Context).ReturnArguments = mock.Arguments{
// 			mockTxRepo,
// 			nil,
// 		}

// 		mockServer := s.NewMockServer(setupmock.WithMockRepositories(&mockRepo.Mock), setupmock.WithMockMailServer(&mockMailServer.Mock))

// 		// Act
// 		response := mockServer.Account.SignUp(params())

// 		// Assert
// 		res := suiteutils.NewHttpResponseWriter()
// 		cusProducer := runtime.ProducerFunc(suiteutils.MyProducer)
// 		response.WriteResponse(res, cusProducer)

// 		expect := []byte(`{"details":"send mail verification failed"}`)
// 		assertions.Equal(string(expect), res.Body.String())
// 	}
// 	{ // add otp to redis failed
// 		// Arrange
// 		mockTxRepo := s.MockRepository()
// 		mockTxRepo.EXPECT().Account().ReturnArguments = mock.Arguments{
// 			func() repositories.AccountServices {
// 				acc := mocks.AccountServices{}
// 				acc.EXPECT().SignUp(s.Context, repositories.SignUpAccountRequest{
// 					CreateAccountRequest: repositories.CreateAccountRequest{
// 						UserID:   testUserID,
// 						Email:    testEmail,
// 						Password: testPassword,
// 					},
// 				}).ReturnArguments = mock.Arguments{
// 					repositories.SignInReply{
// 						Token: testToken,
// 					}, nil,
// 				}
// 				acc.EXPECT().AddOTP(s.Context, "james@gmail.com", "111111").ReturnArguments =
// 					mock.Arguments{
// 						repoErrors.Error{
// 							Code:    0,
// 							Details: "failed to save otp",
// 						},
// 					}
// 				acc.EXPECT().DelOTP(s.Context, "james@gmail.com").ReturnArguments =
// 					mock.Arguments{nil}

// 				return &acc
// 			}(),
// 		}
// 		mockTxRepo.EXPECT().RollBack().ReturnArguments = mock.Arguments{
// 			nil,
// 		}

// 		mockRepo := s.MockRepository()
// 		mockRepo.EXPECT().Begin(s.Context).ReturnArguments = mock.Arguments{
// 			mockTxRepo,
// 			nil,
// 		}

// 		mockMailServer := s.MockMailServer()
// 		mockMailServer.EXPECT().SendAccountVerification(s.Context, mailserver.MailVerifyRequest{
// 			Recipient: "james@gmail.com",
// 		}).ReturnArguments = mock.Arguments{"111111", nil}

// 		mockServer := s.NewMockServer(setupmock.WithMockRepositories(&mockRepo.Mock), setupmock.WithMockMailServer(&mockMailServer.Mock))

// 		// Act
// 		response := mockServer.Account.SignUp(params())

// 		// Assert
// 		res := suiteutils.NewHttpResponseWriter()
// 		cusProducer := runtime.ProducerFunc(suiteutils.MyProducer)
// 		response.WriteResponse(res, cusProducer)

// 		expect := []byte(`{"details":"failed to save otp"}`)
// 		assertions.Equal(string(expect), res.Body.String())
// 	}
// }
