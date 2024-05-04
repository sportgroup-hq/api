package auth

import (
	"context"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/sportgroup-hq/api/internal/models"
)

func (s *Service) GenerateToken(ctx context.Context, user *models.User) (*models.UserCredentials, error) {
	expiresIn := time.Now().Add(time.Hour * 24 * 365 * 40)
	token, err := s.generateToken(jwt.MapClaims{
		"id":  user.ID,
		"exp": expiresIn.Unix(),
	})
	if err != nil {
		return nil, fmt.Errorf("generate token: %w", err)
	}

	credentials := &models.UserCredentials{
		UserID:      user.ID,
		AccessToken: token,
		ExpiresAt:   expiresIn,
		CreatedAt:   time.Now(),
	}

	//if err = s.repo.Auth().CreateCredentials(ctx, credentials); err != nil {
	//	return nil, fmt.Errorf("create credentials: %w", err)
	//}

	panic("implement me")

	return credentials, nil
}

func (s *Service) Login(ctx context.Context, accessToken string) (*models.User, error) {
	claims := jwt.MapClaims{}
	token, err := jwt.ParseWithClaims(accessToken, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(s.cfg.JwtSecret), nil
	})
	if err != nil {
		return nil, fmt.Errorf("parse token: %w", err)
	}

	if !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	userID, ok := claims["id"].(string)
	if !ok {
		return nil, fmt.Errorf("invalid token")
	}

	//user, err := s.repo.User().GetUserByID(ctx, uuid.MustParse(userID))
	//if err != nil {
	//	return nil, fmt.Errorf("get user by id: %w", err)
	//}
	//
	//return user, nil
	_ = userID
	panic("implement me")
}

func (s *Service) generateToken(claims jwt.MapClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(s.cfg.JwtSecret))
}
