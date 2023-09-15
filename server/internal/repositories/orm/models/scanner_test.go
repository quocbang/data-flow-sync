package models

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestScanJson(t *testing.T) {
	assertion := assert.New(t)
	type testStruct struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	}

	{
		// Arrange
		const (
			id   int    = 1
			name string = "TEST_NAME"
		)
		data := fmt.Sprintf(`{"id":%d, "name":"%s"}`, id, name)
		var replyOfString testStruct
		var replyOfByte testStruct

		// Act
		errOfString := ScanJSON(data, &replyOfString)
		errOfByte := ScanJSON([]byte(data), &replyOfByte)

		// Assert
		assertion.NoError(errOfString)
		assertion.NoError(errOfByte)
		expected := testStruct{
			ID:   id,
			Name: name,
		}
		assertion.Equal(expected, replyOfString)
		assertion.Equal(expected, replyOfByte)
	}

	{
		// Arrange
		var reply testStruct

		// Act
		err := ScanJSON(1, &reply)

		// Assert
		assertion.Error(err)
		assertion.Empty(reply)
		expected := "bad src type [int] for struct [*models.testStruct]"
		assertion.Equal(expected, err.Error())
	}
}
