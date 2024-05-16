package httpserver

import (
	"errors"
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sportgroup-hq/api/internal/models"
)

func (s *Server) error(c *gin.Context, err error) {
	switch {
	case errors.Is(err, models.NotFoundError):
		c.AbortWithStatus(http.StatusNotFound)
	default:
		slog.ErrorContext(c, err.Error())
		c.AbortWithStatus(http.StatusInternalServerError)
	}
}
