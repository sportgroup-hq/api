package httpserver

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (s *Server) getMeHandler(c *gin.Context) {
	user, err := s.userSrv.GetUserByID(c, c.MustGet(userIDKey).(uuid.UUID))
	if err != nil {
		s.error(c, err)
		return
	}

	c.JSON(http.StatusOK, user)
}
