package httpserver

import (
	"errors"
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/sportgroup-hq/api/internal/models"
)

func (s *Server) error(ctx *gin.Context, err error) {
	if errors.As(err, &validator.ValidationErrors{}) {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	internalErr := models.Error(err)

	if errors.Is(internalErr, models.ErrUnknown) {
		slog.ErrorContext(ctx, err.Error())
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	ctx.AbortWithStatusJSON(internalErr.HTTPStatusCode, gin.H{
		"error": internalErr.Error(),
		"code":  internalErr.InternalCode,
	})
}
