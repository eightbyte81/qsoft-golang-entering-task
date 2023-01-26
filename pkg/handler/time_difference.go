package handler

import (
	"github.com/gin-gonic/gin"
	"math"
	"net/http"
)

func (h *Handler) getTimeDifference(c *gin.Context) {
	yearString := c.Param("year")

	difference, err := h.services.TimeDifference.GetTimeDifference(yearString)

	if err != nil {
		c.String(http.StatusBadRequest, "Error getting year value: %s", yearString)
		return
	}

	if difference < 0 {
		c.String(http.StatusOK, "Days left: %d", int(math.Abs(difference)))
		return
	} else {
		c.String(http.StatusOK, "Days gone: %d", int(difference))
		return
	}
}
