package models

import (
	"fmt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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
	GrpcError    = 4
)

// Grpc error
type GrpcStatusError interface {
	GRPCStatus() *status.Status
	Error() string
}

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
	case GrpcStatusError:
		s := v.GRPCStatus()
		if s.Code() == codes.Unknown {
			code = 0
		} else if s.Code() < 20 {
			code = 4
		}
		txt = v.Error()
	default:
		txt = v.Error()
		code = 0
	}

	if len(extMsg) != 0 {
		txt = strings.Join(extMsg, " : ") + " : " + txt
	}

	return &Error{Code: code, Text: txt}
}
