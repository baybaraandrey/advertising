package web

import "github.com/pkg/errors"

// FieldError ...
type FieldError struct {
	Field string `json:"field"`
	Error string `json:"error"`
}

// ErrorResponse ...
type ErrorResponse struct {
	Error  string       `json:"error"`
	Fields []FieldError `json:"fields,omitempty"`
}

// Error ...
type Error struct {
	Err    error
	Status int
	Fields []FieldError
}

// NewRequestError ...
func NewRequestError(err error, status int) error {
	return &Error{err, status, nil}
}

// Error ...
func (err *Error) Error() string {
	return err.Err.Error()
}

// shutdown ...
type shutdown struct {
	Message string
}

// NewShutdownError ...
func NewShutdownError(message string) error {
	return &shutdown{Message: message}
}

func (s *shutdown) Error() string {
	return s.Message
}

func IsShutdown(err error) bool {
	if _, ok := errors.Cause(err).(*shutdown); ok {
		return true
	}
	return false
}
