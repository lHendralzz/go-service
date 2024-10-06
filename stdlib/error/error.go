package error

import (
	"net/http"

	"github.com/palantir/stacktrace"
)

const (
	ErrorFailedLogin = stacktrace.ErrorCode(iota)
)

type ErrorMessage map[stacktrace.ErrorCode]Message

var ErrorMessages = ErrorMessage{
	ErrorFailedLogin: {
		StatusCode: http.StatusUnauthorized,
		Message:    "Failed Login",
	},
}

type Message struct {
	StatusCode int    `json:"status_code"`
	Message    string `json:"message"`
}
