package errors

import (
	"errors"
)

var (
	ErrDataExisted  = errors.New("data already existed")
	ErrDataNotFound = errors.New("data not found")
)

type Error struct {
	Code   Code
	Detail string
}

func (e Error) Error() string {
	msg := e.Code.String()
	if len(e.Detail) > 0 {
		msg += ": " + e.Detail
	}
	return msg
}
