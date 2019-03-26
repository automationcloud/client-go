package client

import (
	"errors"
	"net/http"
	"strings"
)

var (
	ServerError = errors.New("server error")
	ClientError = errors.New("client error")
)

// ValidationError contains information about details of validation failure.
type ValidationError struct {
	Messages []string
}

// Error makes string representation of a validation error
func (v ValidationError) Error() string {
	return "validation failed: " + strings.Join(v.Messages, ", ")
}

// parse ValidationError from http response
func validationError(resp *http.Response) error {
	var b struct {
		Details struct {
			Messages []string `json:"messages"`
		} `json:"details"`
	}

	if err := readBody(resp, &b); err != nil {
		return err
	}
	return ValidationError{b.Details.Messages}
}
