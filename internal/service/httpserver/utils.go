package httpserver

import (
	"errors"
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sportgroup-hq/api/internal/models"
)

func (s *Server) error(ctx *gin.Context, err error) {
	switch {
	case errors.Is(err, models.NotFoundError):
		ctx.AbortWithStatus(http.StatusNotFound)
	default:
		slog.ErrorContext(ctx, err.Error())
		ctx.AbortWithStatus(http.StatusInternalServerError)
	}
}
