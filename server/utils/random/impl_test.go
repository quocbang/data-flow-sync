package random

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRandomString(t *testing.T) {
	assertion := assert.New(t)

	reply := RandomString(20)
	assertion.NotNil(reply)
	assertion.Len(reply, 20)
}
