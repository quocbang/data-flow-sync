package account

import (
	"testing"

	s "github.com/stretchr/testify/suite"

	"github.com/quocbang/data-flow-sync/server/internal/servicestest/internal/suite"
)

type Suite struct {
	suite.BasicSuite
}

func NewSuite() *Suite {
	s := suite.NewSuite()
	return &Suite{BasicSuite: *s}
}

func TestSuite(t *testing.T) {
	s.Run(t, NewSuite())
}
