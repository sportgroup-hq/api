package httpserver

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/sportgroup-hq/api/internal/models"
)

func (s *Server) getMeHandler(ctx *gin.Context) {
	user, err := s.userSrv.GetUserByID(ctx, ctx.MustGet(userIDKey).(uuid.UUID))
	if err != nil {
		s.error(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, user)
}

func (s *Server) updateMeHandler(c *gin.Context) {
	var uur models.UpdateUserRequest

	if err := c.ShouldBindJSON(&uur); err != nil {
		s.error(c, err)
		return
	}

	uur.ID = c.MustGet(userIDKey).(uuid.UUID)

	user, err := s.userSrv.UpdateUser(c, uur)
	if err != nil {
		s.error(c, err)
		return
	}

	c.JSON(http.StatusOK, user)
}
