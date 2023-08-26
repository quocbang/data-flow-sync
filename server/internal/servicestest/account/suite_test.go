package account

import (
	"testing"

	s "github.com/stretchr/testify/suite"

	suiteutils "github.com/quocbang/data-flow-sync/server/internal/servicestest/internal/suite"
)

type Suite struct {
	suiteutils.BasicSuite
}

func NewSuite() *Suite {
	s := suiteutils.NewSuite()
	return &Suite{BasicSuite: *s}
}

func TestSuite(t *testing.T) {
	s.Run(t, NewSuite())
}
