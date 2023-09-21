package models

import (
	"encoding/json"
	"reflect"
)

// ScanJSON asserts src as []byte or string type and unmarshal to ptr.
func ScanJSON(src, ptr any) error {
	var data []byte
	switch v := src.(type) {
	case []byte:
		data = make([]byte, len(v))
		copy(data, v)
	case string:
		data = []byte(v)
	default:
		return &scanBadTypeError{
			structName: reflect.TypeOf(ptr).String(),
			src:        src,
		}
	}
	return json.Unmarshal(data, ptr)
}
