package httpserver

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/sportgroup-hq/api/internal/models"
)

func (s *Server) getGroups(ctx *gin.Context) {
	groups, err := s.groupSrv.GetGroupsByUser(ctx, ctx.MustGet(userIDKey).(uuid.UUID))
	if err != nil {
		s.error(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, groups)
}

func (s *Server) getGroupByID(ctx *gin.Context) {
	groupID, err := uuid.Parse(ctx.Param(GroupIDParam))
	if err != nil {
		s.error(ctx, models.ErrPathMalformed)
		return
	}

	userID := ctx.MustGet(userIDKey).(uuid.UUID)

	group, err := s.groupSrv.GetGroupByID(ctx, userID, groupID)
	if err != nil {
		s.error(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, group)
}

func (s *Server) createGroup(ctx *gin.Context) {
	var cgr models.CreateGroupRequest

	if err := ctx.ShouldBind(&cgr); err != nil {
		s.error(ctx, err)
		return
	}

	creatorID := ctx.MustGet(userIDKey).(uuid.UUID)

	group, err := s.groupSrv.CreateGroup(ctx, creatorID, cgr)
	if err != nil {
		s.error(ctx, err)
		return
	}

	ctx.JSON(http.StatusCreated, group)
}

func (s *Server) updateGroup(ctx *gin.Context) {
	var updateGroupRequest models.UpdateGroupRequest

	if err := ctx.ShouldBind(&updateGroupRequest); err != nil {
		s.error(ctx, err)
		return
	}

	userID := ctx.MustGet(userIDKey).(uuid.UUID)

	groupID, err := uuid.Parse(ctx.Param(GroupIDParam))
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

	ctx.JSON(http.StatusOK, group)
}

func (s *Server) joinGroup(ctx *gin.Context) {
	var jgr models.JoinGroupRequest

	if err := ctx.ShouldBind(&jgr); err != nil {
		s.error(ctx, err)
		return
	}

	userID := ctx.MustGet(userIDKey).(uuid.UUID)

	if err := s.groupSrv.JoinGroup(ctx, userID, jgr.Code); err != nil {
		s.error(ctx, err)
		return
	}

	ctx.AbortWithStatus(http.StatusOK)
}

func (s *Server) deleteGroup(ctx *gin.Context) {
	groupID, err := uuid.Parse(ctx.Param(GroupIDParam))
	if err != nil {
		s.error(ctx, models.ErrPathMalformed)
		return
	}

	userID := ctx.MustGet(userIDKey).(uuid.UUID)

	if err = s.groupSrv.DeleteGroup(ctx, userID, groupID); err != nil {
		s.error(ctx, err)
		return
	}

	ctx.AbortWithStatus(http.StatusNoContent)
}

func (s *Server) leaveGroup(ctx *gin.Context) {
	groupID, err := uuid.Parse(ctx.Param(GroupIDParam))
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

func (s *Server) getGroupMembers(ctx *gin.Context) {
	groupID, err := uuid.Parse(ctx.Param(GroupIDParam))
	if err != nil {
		s.error(ctx, models.ErrPathMalformed)
		return
	}

	userID := ctx.MustGet(userIDKey).(uuid.UUID)

	members, err := s.groupSrv.GetGroupMembers(ctx, userID, groupID)
	if err != nil {
		s.error(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, members)
}

func (s *Server) getGroupRecords(ctx *gin.Context) {
	groupID, err := uuid.Parse(ctx.Param(GroupIDParam))
	if err != nil {
		s.error(ctx, models.ErrPathMalformed)
		return
	}

	userID := ctx.MustGet(userIDKey).(uuid.UUID)

	records, err := s.groupSrv.GetGroupRecords(ctx, userID, groupID)
	if err != nil {
		s.error(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, records)
}
