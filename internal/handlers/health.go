package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Healthz reports liveness. Always returns 200; deeper readiness
// checks (DB, downstream services) belong on a separate endpoint.
func Healthz(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}
