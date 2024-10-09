package restHandler

import "net/http"

const (
	POST    string = http.MethodPost
	GET     string = http.MethodGet
	PUT     string = http.MethodPut
	DELETE  string = http.MethodDelete
	OPTIONS string = http.MethodOptions
	CONNECT string = http.MethodConnect
	HEAD    string = http.MethodHead
	PATCH   string = http.MethodPatch
)
