package httpserver

import (
	"log/slog"
	"time"

	"github.com/gin-contrib/cors"
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
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           24 * time.Hour,
	}))

	addOpenAPIDocsRouter(r)

	api := r.Group("/api/v1")

	api.GET("/ping", pingHandler)

	authorized := api.Use(s.authMiddleware)

	// Users
	authorized.GET("/users/me", s.getMeHandler)
	authorized.PATCH("/users/me", s.updateMeHandler)

	// Groups
	authorized.GET("/groups", s.getGroupsHandler)
	authorized.GET("/groups/:group_id", s.getGroupByIDHandler)
	authorized.POST("/groups", s.createGroupHandler)
	authorized.POST("/groups/join", s.joinGroupHandler)
	authorized.DELETE("/groups/:group_id", s.deleteGroupHandler)
	authorized.PATCH("/groups/:group_id", s.updateGroupHandler)
	authorized.POST("/groups/:group_id/leave", s.leaveGroupHandler)
	authorized.GET("/groups/:group_id/records", s.getGroupRecordsHandler)

	authorized.GET("/groups/:group_id/members", s.getGroupMembersHandler)

	authorized.GET("/groups/:group_id/events", s.getEventsHandler)
	authorized.POST("/groups/:group_id/events", s.createEventHandler)
	authorized.GET("/groups/:group_id/events/:event_id", s.getEventHandler)
	//authorized.PATCH("/groups/:group_id/events/:event_id", s.updateEventHandler)
	authorized.DELETE("/groups/:group_id/events/:event_id", s.deleteEventHandler)

	authorized.POST("/groups/:group_id/events/:event_id/records/:record_id", s.setRecordValueHandler) // ANSWER TO RECORD

	slog.Info("Starting HTTP server on " + s.cfg.HTTP.Address + "...")

	return r.Run(s.cfg.HTTP.Address)
}
