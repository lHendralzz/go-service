package error

import (
	"net/http"

	"github.com/palantir/stacktrace"
)

const (
	ErrorLogin = stacktrace.ErrorCode(iota)
	ErrorGenerateJWTToken
	ErrorQuery
	ErrorUnauthorized
	ErrorInvalidRequest
	ErrorBeginTransaction
	ErrorRedisLock
	ErrorLockedOrder
	ErrorReduceProduct
)

type ErrorMessage map[stacktrace.ErrorCode]Message

var ErrorMessages = ErrorMessage{
	ErrorLogin: {
		StatusCode: http.StatusUnauthorized,
		Message:    "Failed Login Email or password is invalid",
	},
	ErrorGenerateJWTToken: {
		StatusCode: http.StatusInternalServerError,
		Message:    "Failed Generate JWT Token",
	},
	ErrorQuery: {
		StatusCode: http.StatusInternalServerError,
		Message:    "Failed Run Query",
	},
	ErrorUnauthorized: {
		StatusCode: http.StatusUnauthorized,
		Message:    "Invalid Token",
	},

	ErrorInvalidRequest: {
		StatusCode: http.StatusBadRequest,
		Message:    "Invalid Request",
	},

	ErrorBeginTransaction: {
		StatusCode: http.StatusInternalServerError,
		Message:    "Failed Begin Transaciton",
	},
	ErrorRedisLock: {
		StatusCode: http.StatusInternalServerError,
		Message:    "Failed Redis",
	},

	ErrorLockedOrder: {
		StatusCode: http.StatusUnprocessableEntity,
		Message:    "Failed Checkout Order Locked",
	},
	ErrorReduceProduct: {
		StatusCode: http.StatusInternalServerError,
		Message:    "Failed Reduce Cost",
	},
}

type Message struct {
	StatusCode int    `json:"status_code"`
	Message    string `json:"message"`
}
