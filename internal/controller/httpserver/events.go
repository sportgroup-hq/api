package httpserver

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/sportgroup-hq/api/internal/models"
)

func (s *Server) createEvent(ctx *gin.Context) {
	groupID, err := uuid.Parse(ctx.Param(GroupIDParam))
	if err != nil {
		s.error(ctx, models.ErrPathMalformed)
		return
	}

	userID := ctx.MustGet(userIDKey).(uuid.UUID)

	var event models.CreateEventRequest

	if err = ctx.ShouldBind(&event); err != nil {
		s.error(ctx, err)
		return
	}

	createdEvent, err := s.eventSrv.CreateEvent(ctx, userID, groupID, &event)
	if err != nil {
		s.error(ctx, err)
		return
	}

	ctx.JSON(http.StatusCreated, createdEvent)
}

func (s *Server) getEvents(ctx *gin.Context) {
	groupID, err := uuid.Parse(ctx.Param(GroupIDParam))
	if err != nil {
		s.error(ctx, models.ErrPathMalformed)
		return
	}

	userID := ctx.MustGet(userIDKey).(uuid.UUID)

	events, err := s.eventSrv.GetEventsByGroup(ctx, userID, groupID)
	if err != nil {
		s.error(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, events)
}

func (s *Server) getEvent(ctx *gin.Context) {
	groupID, err := uuid.Parse(ctx.Param(GroupIDParam))
	if err != nil {
		s.error(ctx, models.ErrPathMalformed)
		return
	}

	eventID, err := uuid.Parse(ctx.Param(EventIDParam))
	if err != nil {
		s.error(ctx, models.ErrPathMalformed)
		return
	}

	userID := ctx.MustGet(userIDKey).(uuid.UUID)

	event, err := s.eventSrv.GetEventByID(ctx, userID, groupID, eventID)
	if err != nil {
		s.error(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, event)
}

func (s *Server) updateEventHandler(ctx *gin.Context) {
	//groupID, err := uuid.Parse(ctx.Param(GroupIDParam))
	//if err != nil {
	//	s.error(ctx, models.ErrPathMalformed)
	//	return
	//}
	//
	//eventID, err := uuid.Parse(ctx.Param(EventIDParam))
	//if err != nil {
	//	s.error(ctx, models.ErrPathMalformed)
	//	return
	//}
	//
	//userID := ctx.MustGet(userIDKey).(uuid.UUID)
	//
	//var updateEvent models.UpdateEventRequest
	//
	//if err = ctx.ShouldBind(&updateEvent); err != nil {
	//	s.error(ctx, err)
	//	return
	//}
	//
	//updatedEvent, err := s.eventSrv.UpdateEvent(ctx, userID, groupID, eventID, &updateEvent)
	//if err != nil {
	//	s.error(ctx, err)
	//	return
	//}
	//
	//ctx.JSON(http.StatusOK, updatedEvent)
}

func (s *Server) deleteEvent(ctx *gin.Context) {
	groupID, err := uuid.Parse(ctx.Param(GroupIDParam))
	if err != nil {
		s.error(ctx, models.ErrPathMalformed)
		return
	}

	eventID, err := uuid.Parse(ctx.Param(EventIDParam))
	if err != nil {
		s.error(ctx, models.ErrPathMalformed)
		return
	}

	userID := ctx.MustGet(userIDKey).(uuid.UUID)

	err = s.eventSrv.DeleteEvent(ctx, userID, groupID, eventID)
	if err != nil {
		s.error(ctx, err)
		return
	}

	ctx.Status(http.StatusNoContent)
}

func (s *Server) setRecordValue(ctx *gin.Context) {
	groupID, err := uuid.Parse(ctx.Param(GroupIDParam))
	if err != nil {
		s.error(ctx, models.ErrPathMalformed)
		return
	}

	eventID, err := uuid.Parse(ctx.Param(EventIDParam))
	if err != nil {
		s.error(ctx, models.ErrPathMalformed)
		return
	}

	recordID, err := uuid.Parse(ctx.Param(RecordIDParam))
	if err != nil {
		s.error(ctx, models.ErrPathMalformed)
		return
	}

	userID := ctx.MustGet(userIDKey).(uuid.UUID)

	var recordValue models.EventRecordValue

	if err = ctx.ShouldBind(&recordValue); err != nil {
		s.error(ctx, err)
		return
	}

	recordValue.RecordID = recordID

	err = s.eventSrv.SetEventRecordValue(ctx, userID, groupID, eventID, recordID, &recordValue)
	if err != nil {
		s.error(ctx, err)
		return
	}

	ctx.Status(http.StatusNoContent)
}

func (s *Server) createEventComment(ctx *gin.Context) {
	groupID, err := uuid.Parse(ctx.Param(GroupIDParam))
	if err != nil {
		s.error(ctx, models.ErrPathMalformed)
		return
	}

	eventID, err := uuid.Parse(ctx.Param(EventIDParam))
	if err != nil {
		s.error(ctx, models.ErrPathMalformed)
		return
	}

	userID := ctx.MustGet(userIDKey).(uuid.UUID)

	var comment models.CreateCommentRequest

	if err = ctx.ShouldBind(&comment); err != nil {
		s.error(ctx, err)
		return
	}

	if err = s.eventSrv.CreateEventComment(ctx, userID, groupID, eventID, &comment); err != nil {
		s.error(ctx, err)
		return
	}

	ctx.AbortWithStatus(http.StatusCreated)
}

func (s *Server) getEventComments(ctx *gin.Context) {
	groupID, err := uuid.Parse(ctx.Param(GroupIDParam))
	if err != nil {
		s.error(ctx, models.ErrPathMalformed)
		return
	}

	eventID, err := uuid.Parse(ctx.Param(EventIDParam))
	if err != nil {
		s.error(ctx, models.ErrPathMalformed)
		return
	}

	userID := ctx.MustGet(userIDKey).(uuid.UUID)

	comments, err := s.eventSrv.GetEventComments(ctx, userID, groupID, eventID)
	if err != nil {
		s.error(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, comments)
}
