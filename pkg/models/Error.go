package models

type ErrorCode int

const (
	Success ErrorCode = iota
	InvalidParams
)

type Error struct {
	ErrorCode    ErrorCode `json:"err_code"`
	ErrorMessage string    `json:"err_msg"`
}
