package setupmock

import (
	"time"

	"github.com/quocbang/data-flow-sync/server/api"
	"github.com/quocbang/data-flow-sync/server/internal/mocks"
	"github.com/quocbang/data-flow-sync/server/internal/services"
	"github.com/stretchr/testify/mock"
)

type MockServerOptions struct {
	TokenLifeTime time.Duration
	MRExpiryTime  int64
}

func NewMockRepositories() *mocks.Repositories {
	return &mocks.Repositories{}
}

func NewMockMailServer() *mocks.MailServer {
	return &mocks.MailServer{}
}

func NewMockRedis() *mocks.RedisConn {
	return &mocks.RedisConn{}
}

type Option func(*api.ServiceConfig)

// NewMockServer initialize mock server.
func NewMockServer(opts ...Option) services.Services {
	var config = api.ServiceConfig{}

	for _, opt := range opts {
		opt(&config)
	}
	return *api.NewHandleService(config)
}

func WithMockRepositories(m *mock.Mock) Option {
	return func(sc *api.ServiceConfig) {
		sc.Repo = &mocks.Repositories{
			Mock: mock.Mock{
				Calls:         m.Calls,
				ExpectedCalls: m.ExpectedCalls,
			},
		}
	}
}

func WithMockMailServer(m *mock.Mock) Option {
	return func(sc *api.ServiceConfig) {
		sc.Smtp = &mocks.MailServer{
			Mock: mock.Mock{
				Calls:         m.Calls,
				ExpectedCalls: m.ExpectedCalls,
			},
		}
	}
}

func WithMockRedisServer(m *mock.Mock) Option {
	return func(sc *api.ServiceConfig) {
		sc.Redis = &mocks.RedisConn{
			Mock: mock.Mock{
				Calls:         m.Calls,
				ExpectedCalls: m.ExpectedCalls,
			},
		}
	}
}

func WithRequirementOptions(m MockServerOptions) Option {
	return func(sc *api.ServiceConfig) {
		sc.MRExpiryTime = m.MRExpiryTime
		sc.TokenLifeTime = m.TokenLifeTime
	}
}
