package middlewares

import (
	"net/http"

	"github.com/uptrace/bunrouter"
)

func ErrorMiddleware(next bunrouter.HandlerFunc) bunrouter.HandlerFunc {
	return func(w http.ResponseWriter, req bunrouter.Request) error {
		err := next(w, req)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return err
	}
}
