package grpcserver

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/sportgroup-hq/api/internal/models"
	"github.com/sportgroup-hq/common-lib/api"
)

func (s *Server) GetUser(ctx context.Context, req *api.GetUserRequest) (*api.GetUserResponse, error) {
	id, err := uuid.Parse(req.Id)
	if err != nil {
		return nil, fmt.Errorf("invalid id: %w", err)
	}

	user, err := s.UserSrv.GetUser(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("failed to get user: %w", err)
	}

	return &api.GetUserResponse{User: models.UserToPB(user)}, nil
}
