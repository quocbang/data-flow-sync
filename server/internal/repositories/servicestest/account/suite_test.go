package account

import (
	"testing"

	s "github.com/stretchr/testify/suite"

	"github.com/quocbang/data-flow-sync/server/internal/repositories/orm/models"
	"github.com/quocbang/data-flow-sync/server/internal/repositories/servicestest/internal/suite"
)

type Suite struct {
	suite.SuiteConfig
}

// NewSuite create all necessary.
func NewSuite() *Suite {
	s := suite.NewSuiteTest(suite.NewSuiteParameters{
		RelativeModels:       []models.Models{&models.Account{}},
		ClearDataForEachTest: true,
	})
	return &Suite{
		SuiteConfig: *s,
	}
}
func TestSuite(t *testing.T) {
	// Run the test suite using suite.Run.
	s.Run(t, NewSuite())
}
