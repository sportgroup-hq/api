package httpserver

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/google/uuid"
	"github.com/sportgroup-hq/api/internal/models"
)

type CreateGroupRequest struct {
	Name  string `json:"name" binding:"required"`
	Sport string `json:"sport" binding:"required"`
}

func (s *Server) getGroupsHandler(ctx *gin.Context) {
	groups, err := s.groupSrv.GetUserGroups(ctx, ctx.MustGet(userIDKey).(uuid.UUID))
	if err != nil {
		s.error(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, groups)
}

func (s *Server) createGroupsHandler(ctx *gin.Context) {
	var cgr CreateGroupRequest

	if err := ctx.MustBindWith(&cgr, binding.JSON); err != nil {
		return
	}

	creatorID := ctx.MustGet(userIDKey).(uuid.UUID)
	newGroup := cgr.toGroup()

	group, err := s.groupSrv.CreateGroup(ctx, creatorID, newGroup)
	if err != nil {
		s.error(ctx, err)
		return
	}

	ctx.JSON(http.StatusCreated, group)
}

func (r CreateGroupRequest) toGroup() *models.Group {
	return &models.Group{
		Name:  r.Name,
		Sport: r.Sport,
	}
}
