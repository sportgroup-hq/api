package httpserver

import (
	"github.com/sportgroup-hq/api/internal/service"
)

type Server struct {
	auth service.Auth
	user service.User
}

func New(authSrv service.Auth, userSrv service.User) *Server {
	return &Server{
		auth: authSrv,
		user: userSrv,
	}
}
