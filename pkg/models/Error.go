package models

import (
	"fmt"
	"strings"
)

const (
	Success = "success"
	Failure = "failure"
)

const (
	InvalidParam = 1
	AssertError  = 2
	DBError      = 3
)

type Error struct {
	Code int32
	Text string
}

func (e Error) Error() string {
	return fmt.Sprintf("code: %d, msg is: %s", e.Code, e.Text)
}

func NewError(code int32, msg string) *Error {
	return &Error{Code: code, Text: msg}
}

// error adaptor, may convert many error type to `error` type
func ErrorWrapper(err error, extMsg ...string) *Error {
	var txt string
	var code int32
	switch v := err.(type) {
	case Error:
		txt = v.Text
		code = v.Code
	default:
		txt = v.Error()
		code = 0
	}

	if len(extMsg) != 0 {
		txt = strings.Join(extMsg, " : ") + " : " + txt
	}

	return &Error{Code: code, Text: txt}
}
