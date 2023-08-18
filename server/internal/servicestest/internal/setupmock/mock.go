package setupmock

import (
	"time"

	"github.com/quocbang/data-flow-sync/server/api"
	"github.com/quocbang/data-flow-sync/server/internal/mocks"
	"github.com/quocbang/data-flow-sync/server/internal/repositories"
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

// NewMockServer initialize mock server.
func NewMockServer(m *mock.Mock, opts MockServerOptions) services.Services {
	return *api.NewHandleService(api.ServiceConfig{
		Repo:          newMockRepositories(m),
		TokenLifeTime: opts.TokenLifeTime,
		MRExpiryTime:  opts.MRExpiryTime,
	})
}

func newMockRepositories(m *mock.Mock) repositories.Repositories {
	return &mocks.Repositories{
		Mock: mock.Mock{
			Calls:         m.Calls,
			ExpectedCalls: m.ExpectedCalls,
		},
	}
}
