package suite

import (
	"context"
	"io"
	"log"
	"net/http"
	"net/http/httptest"

	"github.com/quocbang/data-flow-sync/server/internal/mocks"
	"github.com/quocbang/data-flow-sync/server/internal/services"
	"github.com/quocbang/data-flow-sync/server/internal/servicestest/internal/setupmock"
	"github.com/quocbang/data-flow-sync/server/utils/random"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"go.uber.org/zap"
)

type BasicSuite struct {
	*suite.Suite
	Context         context.Context
	HttpTestRequest func(method string, target string, body io.Reader) *http.Request
	MockRepository  func() *mocks.Repositories
	NewMockServer   func(m *mock.Mock, msOpts setupmock.MockServerOptions) services.Services
}

func NewSuite() *BasicSuite {
	field := []zap.Field{
		zap.String("random seed", random.RandomString(30)),
	}
	logger.Info("start service test", field...)
	return &BasicSuite{
		Suite:           &suite.Suite{},
		HttpTestRequest: httpTestRequest,
		MockRepository:  setupmock.NewMockRepositories,
		NewMockServer:   setupmock.NewMockServer,
	}
}

func (b *BasicSuite) SetupSuite() {
	b.Context = context.Background()
}

func (b *BasicSuite) TearDownSuite() {

}

var logger *zap.Logger

func init() {
	var err error
	logger, err = zap.NewDevelopment()
	if err != nil {
		log.Fatal(err)
	}
}

// httpTestRequest is define http test request.
func httpTestRequest(method string, target string, body io.Reader) *http.Request {
	return httptest.NewRequest(method, target, body)
}
