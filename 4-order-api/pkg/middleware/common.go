// Package middleware
package middleware

import "net/http"

type ResponseWrapper struct {
	http.ResponseWriter
	StatusCode int
}

func (wrapper *ResponseWrapper) WriteHeader(statusCode int) {
	wrapper.ResponseWriter.WriteHeader(statusCode)
	wrapper.StatusCode = statusCode
}
