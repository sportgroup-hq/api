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

type JoinGroupRequest struct {
	Code string `json:"code" binding:"required,len=6"`
}

func (s *Server) getGroupsHandler(ctx *gin.Context) {
	groups, err := s.groupSrv.GetUserGroups(ctx, ctx.MustGet(userIDKey).(uuid.UUID))
	if err != nil {
		s.error(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, groups)
}

func (s *Server) createGroupHandler(ctx *gin.Context) {
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

func (s *Server) updateGroupHandler(ctx *gin.Context) {
	var updateGroupRequest models.UpdateGroupRequest

	if err := ctx.MustBindWith(&updateGroupRequest, binding.JSON); err != nil {
		return
	}

	userID := ctx.MustGet(userIDKey).(uuid.UUID)

	groupID, err := uuid.Parse(ctx.Param("group_id"))
	if err != nil {
		s.error(ctx, models.ErrPathMalformed)
		return
	}

	updateGroupRequest.ID = groupID

	group, err := s.groupSrv.UpdateGroup(ctx, userID, updateGroupRequest)
	if err != nil {
		s.error(ctx, err)
		return
	}

	ctx.JSON(http.StatusCreated, group)
}

func (s *Server) joinGroupHandler(ctx *gin.Context) {
	var jgr JoinGroupRequest

	if err := ctx.MustBindWith(&jgr, binding.JSON); err != nil {
		return
	}

	userID := ctx.MustGet(userIDKey).(uuid.UUID)

	if err := s.groupSrv.JoinGroup(ctx, userID, jgr.Code); err != nil {
		s.error(ctx, err)
		return
	}

	ctx.AbortWithStatus(http.StatusOK)
}

func (s *Server) deleteGroupHandler(ctx *gin.Context) {
	groupID, err := uuid.Parse(ctx.Param("group_id"))
	if err != nil {
		s.error(ctx, models.ErrPathMalformed)
		return
	}

	userID := ctx.MustGet(userIDKey).(uuid.UUID)

	if err = s.groupSrv.DeleteGroup(ctx, userID, groupID); err != nil {
		s.error(ctx, err)
		return
	}

	ctx.AbortWithStatus(http.StatusOK)
}

func (s *Server) leaveGroupHandler(ctx *gin.Context) {
	groupID, err := uuid.Parse(ctx.Param("group_id"))
	if err != nil {
		s.error(ctx, models.ErrPathMalformed)
		return
	}

	userID := ctx.MustGet(userIDKey).(uuid.UUID)

	if err = s.groupSrv.LeaveGroup(ctx, userID, groupID); err != nil {
		s.error(ctx, err)
		return
	}

	ctx.AbortWithStatus(http.StatusOK)
}

func (r CreateGroupRequest) toGroup() *models.Group {
	return &models.Group{
		Name:  r.Name,
		Sport: r.Sport,
	}
}
