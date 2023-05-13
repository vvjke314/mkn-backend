package app

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

const (
	AuthHeader = "Authorization"
)

func (a *Application) UserIdentity(c *gin.Context) {
	jwtStr, err := a.GetJWT(c)
	if err != nil {
		newErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}
	userId, err := a.ParseToken(jwtStr)
	if err != nil {
		newErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}

	c.Set("userId", userId)
}

func (a *Application) GetJWT(c *gin.Context) (string, error) {
	header := c.GetHeader(AuthHeader)
	if header == "" {
		return "", errors.New("Empty auth method")
	}

	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 {
		return "", errors.New("Invalid auth header")
	}

	return headerParts[1], nil
}

func (a *Application) GetUserIdByJWT(c *gin.Context) (string, error) {
	jwtStr, err := a.GetJWT(c)
	if err != nil {
		return "", errors.Wrap(err, "Can't get JWT")
	}

	userId, err := a.redis.Get(*a.ctx, jwtStr).Result()
	if err != nil {
		return "", errors.Wrap(err, "Can't get user id from redis")
	}

	return userId, nil
}
