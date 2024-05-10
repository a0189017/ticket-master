package errors

import (
	"encoding/json"

	"github.com/pkg/errors"
)

type ErrorInfo struct {
	Message string `json:"message"`
	Code    string `json:"code"`
	Err     error  `json:"-"`
}

type SystemError struct {
	*ErrorInfo
	StackTrace errors.StackTrace `json:"stack,omitempty" swaggertype:"object"`
}

func (e SystemError) Error() string {
	b, _ := json.Marshal(e)
	return string(b)
}

func (e SystemError) Cause() error {
	return errors.Cause(e.Err)
}
