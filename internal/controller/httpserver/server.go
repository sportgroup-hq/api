package httpserver

import (
	"log/slog"

	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"github.com/sportgroup-hq/common-lib/validation"
)

func (s *Server) Start() error {
	r := gin.Default()

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		validation.Register(v)
	}

	r.Use(gzip.Gzip(gzip.DefaultCompression))

	addOpenAPIDocsRouter(r)

	api := r.Group("/api/v1")

	api.GET("/ping", pingHandler)

	authorized := api.Use(s.authMiddleware)

	// Users
	authorized.GET("/users/me", s.getMe)
	authorized.PATCH("/users/me", s.updateMe)

	// Groups
	authorized.GET("/groups", s.getGroups)
	authorized.GET("/groups/:group_id", s.getGroupByID)
	authorized.POST("/groups", s.createGroup)
	authorized.POST("/groups/join", s.joinGroup)
	authorized.DELETE("/groups/:group_id", s.deleteGroup)
	authorized.PATCH("/groups/:group_id", s.updateGroup)
	authorized.POST("/groups/:group_id/leave", s.leaveGroup)
	authorized.GET("/groups/:group_id/records", s.getGroupRecords)

	authorized.GET("/groups/:group_id/members", s.getGroupMembers)

	authorized.GET("/groups/:group_id/events", s.getEvents)
	authorized.POST("/groups/:group_id/events", s.createEvent)
	authorized.GET("/groups/:group_id/events/:event_id", s.getEvent)
	//authorized.PATCH("/groups/:group_id/events/:event_id", s.updateEvent)
	authorized.DELETE("/groups/:group_id/events/:event_id", s.deleteEvent)

	authorized.POST("/groups/:group_id/events/:event_id/records/:record_id", s.setRecordValue)

	authorized.GET("/groups/:group_id/events/:event_id/comments", s.getEventComments)
	authorized.POST("/groups/:group_id/events/:event_id/comments", s.createEventComment)

	slog.Info("Starting HTTP server on " + s.cfg.HTTP.Address + "...")

	return r.Run(s.cfg.HTTP.Address)
}
