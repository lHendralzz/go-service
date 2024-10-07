package error

import (
	"net/http"

	"github.com/palantir/stacktrace"
)

const (
	ErrorLogin = stacktrace.ErrorCode(iota)
	ErrorGenerateJWTToken
	ErrorQuery
)

type ErrorMessage map[stacktrace.ErrorCode]Message

var ErrorMessages = ErrorMessage{
	ErrorLogin: {
		StatusCode: http.StatusUnauthorized,
		Message:    "Failed Login username or password is invalid",
	},
	ErrorGenerateJWTToken: {
		StatusCode: http.StatusInternalServerError,
		Message:    "Failed Generate JWT Token",
	},
	ErrorQuery: {
		StatusCode: http.StatusInternalServerError,
		Message:    "Failed Generate JWT Token",
	},
}

type Message struct {
	StatusCode int    `json:"status_code"`
	Message    string `json:"message"`
}
