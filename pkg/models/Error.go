package models

import (
	"fmt"
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

type Error struct {
	Code int32
	Text string
}

func (e Error) Error() string {
	return fmt.Sprintf("code: %d, msg is: %s", e.Code, e.Text)
}
