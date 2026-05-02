package server

import (
	"github.com/gin-gonic/gin"

	"github.com/mihaics/demo-sandbox/internal/handlers"
)

// New builds the HTTP handler with all routes registered. It does not
// start the server; the caller wraps it in an *http.Server so signal
// handling and graceful shutdown stay in one place.
func New() *gin.Engine {
	r := gin.Default()
	r.GET("/healthz", handlers.Healthz)
	return r
}
