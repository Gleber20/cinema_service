package http

import (
	"auth_service/pkg"
	"errors"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

const (
	authorizationHeader = "Authorization"
	UserIDCtx           = "userID"
)

func extractTokenFromHeader(c *gin.Context, headerKey string) (string, error) {
	header := c.GetHeader(headerKey)

	if header == "" {
		return "", errors.New("empty authorization header")
	}

	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 {
		return "", errors.New("invalid authorization header")
	}

	if len(headerParts[1]) == 0 {
		return "", errors.New("empty token")
	}

	return headerParts[1], nil
}

func CheckUserAuthentication(c *gin.Context) {
	token, err := extractTokenFromHeader(c, authorizationHeader)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	userID, isRefresh, err := pkg.ParseToken(token)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	if isRefresh {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "inappropriate token"})
		return
	}

	c.Set(UserIDCtx, userID)
}
