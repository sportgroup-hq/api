package httpserver

import (
	"errors"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type Claims struct {
	jwt.RegisteredClaims
	Subject uuid.UUID `json:"sub"`
	Type    string    `json:"typ"`
}

func (s *Server) authMiddleware(ctx *gin.Context) {
	var claims Claims

	authHeader := ctx.GetHeader("Authorization")

	parts := strings.Split(authHeader, " ")

	if len(parts) != 2 || parts[0] != "Bearer" {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	token := parts[1]

	parsedToken, err := jwt.ParseWithClaims(token, &claims, s.jwtSecretFunc)
	if err != nil || !parsedToken.Valid {
		if errors.Is(err, jwt.ErrTokenExpired) {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "token expired"})
			return
		}
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	if claims.Type != "access" {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
		return
	}

	ctx.Set(userIDKey, claims.Subject)

	ctx.Next()
}

func (s *Server) jwtSecretFunc(_ *jwt.Token) (interface{}, error) {
	return []byte(s.cfg.JWT.Secret), nil
}
