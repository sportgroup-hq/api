package httpserver

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/sportgroup-hq/api/internal/models"
)

func (s *Server) getMe(ctx *gin.Context) {
	user, err := s.userSrv.GetUserByID(ctx, ctx.MustGet(userIDKey).(uuid.UUID))
	if err != nil {
		s.error(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, user)
}

func (s *Server) updateMe(ctx *gin.Context) {
	var uur models.UpdateUserRequest

	if err := ctx.ShouldBindJSON(&uur); err != nil {
		s.error(ctx, err)
		return
	}

	uur.ID = ctx.MustGet(userIDKey).(uuid.UUID)

	user, err := s.userSrv.UpdateUser(ctx, uur)
	if err != nil {
		s.error(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, user)
}
