package diff

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindDiff(t *testing.T) {
	assertion := assert.New(t)

	type Test struct {
		Name  string   `json:"name,omitempty"`
		Age   int      `json:"age,omitempty"`
		Habit []string `json:"habit,omitempty"`
	}

	old := Test{
		Name: "TEST_USER",
		Age:  20,
		Habit: []string{
			"listen to lofi hip hop music",
			"coding at night",
			"ish",
		},
	}

	new := Test{
		Name: "TEST_USER",
		Age:  23,
		Habit: []string{
			"listen to lofi hip hop music",
			"wake up at 6 am",
			"sleep at 10 pm",
			"ish",
		},
	}

	// good case
	{
		// Arrange
		oldByte, err := json.Marshal(old)
		assertion.NoError(err)
		newByte, err := json.Marshal(new)
		assertion.NoError(err)

		// Act
		added, deleted, err := FindDiff[Test](oldByte, newByte)

		// Assert
		assertion.NoError(err)
		expectedAdded, err := json.Marshal(Test{
			Age: 23,
			Habit: []string{
				"",
				"wake up at 6 am",
				"sleep at 10 pm",
				"",
			},
		})
		assertion.NoError(err)
		expectedDeleted, err := json.Marshal(Test{
			Age: 20,
			Habit: []string{
				"",
				"coding at night",
				"",
			},
		})
		assertion.NoError(err)
		assertion.Equal(expectedAdded, added)
		assertion.Equal(expectedDeleted, deleted)
	}

	// bad case: wrong type during find diff
	{
		// Arrange
		oldByte, err := json.Marshal(old)
		assertion.NoError(err)
		newByte, err := json.Marshal(new)
		assertion.NoError(err)

		// Act
		_, _, err = FindDiff[map[string]interface{}](oldByte, newByte)

		// Assert
		assertion.Error(err)
		expected := fmt.Errorf("wrong type during find diff, expect type is [struct or slice] but found [map]")
		assertion.Equal(expected.Error(), err.Error())
	}
}
