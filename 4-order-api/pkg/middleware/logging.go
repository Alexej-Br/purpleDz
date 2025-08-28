// Package middleware
package middleware

import (
	"net/http"

	"github.com/sirupsen/logrus"
)

func Log(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		wrapper := &ResponseWrapper{
			ResponseWriter: w,
			StatusCode:     http.StatusOK,
		}
		next.ServeHTTP(wrapper, r)
		lFields := logrus.Fields{
			"Method":      r.Method,
			"URl":         r.URL.Path,
			"Remote addr": r.RemoteAddr,
			"Status code": wrapper.StatusCode,
		}
		switch {
		case wrapper.StatusCode >= 500:
			logrus.WithFields(lFields).Error("Server's error")
		case wrapper.StatusCode >= 400:
			logrus.WithFields(lFields).Error("User's error")
		case wrapper.StatusCode >= 300:
			logrus.WithFields(lFields).Warn()
		case wrapper.StatusCode >= 200:
			logrus.WithFields(lFields).Debug("OK")
		}
	})
}
