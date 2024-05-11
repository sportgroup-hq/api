package grpcserver

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/sportgroup-hq/api/internal/models"
	"github.com/sportgroup-hq/common-lib/api"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) GetUserByID(ctx context.Context, req *api.GetUserByIDRequest) (*api.GetUserByIDResponse, error) {
	userID, err := uuid.Parse(req.Id)
	if err != nil {
		return nil, fmt.Errorf("invalid user id: %w", err)
	}

	user, err := s.userSrv.GetUserByID(ctx, userID)
	if err != nil {
		return nil, fmt.Errorf("failed to get user: %w", err)
	}

	return &api.GetUserByIDResponse{User: models.UserToPB(user)}, nil
}

func (s *Server) GetUserByEmail(ctx context.Context, req *api.GetUserByEmailRequest) (*api.GetUserByEmailResponse, error) {
	if req.Email == "" {
		return nil, fmt.Errorf("email is required")
	}

	user, err := s.userSrv.GetUserByEmail(ctx, req.Email)
	if err != nil {
		return nil, fmt.Errorf("failed to get user: %w", err)
	}

	if user == nil {
		return nil, status.New(codes.NotFound, "user not found").Err()
	}

	return &api.GetUserByEmailResponse{User: models.UserToPB(user)}, nil
}

func (s *Server) UserExistsByEmail(ctx context.Context, req *api.UserExistsByEmailRequest) (*api.UserExistsByEmailResponse, error) {
	if req.Email == "" {
		return nil, fmt.Errorf("email is required")
	}

	exists, err := s.userSrv.UserExistsByEmail(ctx, req.Email)
	if err != nil {
		return nil, fmt.Errorf("failed to check user exists: %w", err)
	}

	return &api.UserExistsByEmailResponse{Exists: exists}, nil
}

func (s *Server) CreateUser(ctx context.Context, req *api.CreateUserRequest) (*api.CreateUserResponse, error) {
	user := &models.User{
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Email:     req.Email,
	}

	if req.Picture != nil {
		user.Picture = *req.Picture
	}

	user, err := s.userSrv.CreateUser(ctx, user)
	if err != nil {
		return nil, fmt.Errorf("failed to create user: %w", err)
	}

	return &api.CreateUserResponse{User: models.UserToPB(user)}, nil
}
