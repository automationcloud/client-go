package client

import (
	"net/http"
	"strings"
)

type ValidationError struct {
	Messages []string
}

func (v ValidationError) Error() string {
	return "validation failed: " + strings.Join(v.Messages, ", ")
}

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
