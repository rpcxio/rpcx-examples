package error_customized

import (
	"encoding/json"
	"fmt"
)

// Error is customized error.
type Error struct {
	Code int    `json:"code,omitempty"`
	Msg  string `json:"msg,omitempty"`
}

// Error implements error interface.
func (e *Error) Error() string {
	return fmt.Sprintf(`{"code": %d, "msg": "%s"}`, e.Code, e.Msg)
}

func (e *Error) IsServiceError() bool {
	return true
}

// NewError creates a new Error.
func NewError(code int, msg string) *Error {
	return &Error{
		Code: code,
		Msg:  msg,
	}
}

// NewErrorString creates a new Error from string.
func MewErrorString(s string) (*Error, error) {
	var err Error
	e := json.Unmarshal([]byte(s), &err)

	return &err, e
}
