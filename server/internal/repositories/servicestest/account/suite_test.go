package account

import (
	"testing"

	"github.com/quocbang/data-flow-sync/server/internal/repositories/orm/models"
	"github.com/quocbang/data-flow-sync/server/internal/repositories/servicestest/internal/suite"
	"github.com/stretchr/testify/suite"
)

type Suite struct {
	servicestest.SuiteConfig
}

func NewSuite() *Suite {
	s := servicestest.NewSuiteConfig(servicestest.SuiteParameters{
		RelativeModels: []models.Models{&models.Account{}},
		ClearData:      true,
		WithTimeStub:   true,
	})

	return &Suite{
		SuiteConfig: *s,
	}
}

func TestSuite(t *testing.T) {
	suite.Run(t, NewSuite())
}
