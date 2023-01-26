package handler

import (
	"github.com/gin-gonic/gin"
)

const (
	xPingHeader = "X-PING"
	xPongHeader = "X-PONG"
)

func (h *Handler) checkHeader(c *gin.Context) {
	header := c.GetHeader(xPingHeader)

	if header == "ping" {
		c.Header(xPongHeader, "pong")
	}
}
