package middlewares

import (
	log "bareksa-interview-project/util/logger"
	"net/http"

	"github.com/uptrace/bunrouter"
)

func ErrorMiddleware(customLogger *log.CustomLogger) bunrouter.MiddlewareFunc {
	return func(next bunrouter.HandlerFunc) bunrouter.HandlerFunc {
		return func(w http.ResponseWriter, req bunrouter.Request) error {
			if err := next(w, req); err != nil {
				customLogger.WriteLog(log.WARN, err)
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}

			return nil
		}
	}
}
