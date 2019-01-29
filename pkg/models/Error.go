package models

import (
	"fmt"
	"path/filepath"
	"runtime"
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
	Code  int32
	Text  string
	Where string
}

func (e Error) Error() string {
	return fmt.Sprintf("code: %d, msg: %s, where: %s", e.Code, e.Text, e.Where)
}

func Caller(calldepth int, short bool) string {
	_, file, line, ok := runtime.Caller(calldepth + 1)
	if !ok {
		file = "???"
		line = 0
	} else if short {
		file = filepath.Base(file)
	}

	return fmt.Sprintf("%s:%d", file, line)
}
