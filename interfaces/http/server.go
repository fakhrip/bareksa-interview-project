package http

import (
	middlewares "bareksa-interview-project/interfaces/http/middlewares"
	routes "bareksa-interview-project/interfaces/http/routes"
	log "bareksa-interview-project/util/logger"
	"net/http"

	"github.com/klauspost/compress/gzhttp"
	"github.com/rs/cors"
	"github.com/uptrace/bunrouter"
)

func Start(isDebugMode bool, customLogger *log.CustomLogger) {
	router := bunrouter.New()

	c := cors.New(cors.Options{
		Debug: isDebugMode,
	})

	handler := http.Handler(router)
	handler = gzhttp.GzipHandler(handler)
	handler = c.Handler(handler)

	router.Use(middlewares.ErrorMiddleware).
		WithGroup("/api/users", routes.ApiRoutes())

	httpServer := &http.Server{
		Addr:    ":8888",
		Handler: handler,
	}

	customLogger.WriteLog(log.INFO, "Listening on http://localhost:8888")
	customLogger.WriteLog(log.INFO, httpServer.ListenAndServe())
}
