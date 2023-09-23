package diff

import (
	"testing"

	"github.com/quocbang/data-flow-sync/server/internal/repositories/errors"
	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v3"
)

func TestFindDiff(t *testing.T) {
	type TestFindDiff struct {
		Name string `yaml:"name,omitempty"`
		Age  int    `yaml:"age,omitempty"`
	}
	assertion := assert.New(t)

	// wrong type expected is struct
	{
		// Arrange
		// Act
		_, _, err := FindDiff[string]([]byte{}, []byte{})

		// Assert
		assertion.Error(err)
		expected := errors.Error{
			Details: "wrong type during  find diff, expect type is [struct] but found [string]",
		}
		assertion.Equal(expected, err)
	}

	// good case
	{
		// Arrange
		x := TestFindDiff{
			Name: "test_name_x",
			Age:  18,
		}
		y := TestFindDiff{
			Name: "test_name_y",
			Age:  18,
		}
		bx, err := yaml.Marshal(x)
		assertion.NoError(err)
		by, err := yaml.Marshal(y)
		assertion.NoError(err)
		// Act
		added, deleted, err := FindDiff[TestFindDiff](bx, by)

		// Assert
		assertion.NoError(err)
		addedExpected, err := yaml.Marshal(TestFindDiff{
			Name: "test_name_y",
		})
		assertion.NoError(err)
		deletedExpected, err := yaml.Marshal(TestFindDiff{
			Name: "test_name_x",
		})
		assertion.NoError(err)
		assertion.Equal(addedExpected, added)
		assertion.Equal(deletedExpected, deleted)
	}
}
