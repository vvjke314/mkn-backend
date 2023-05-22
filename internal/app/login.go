package app

import (
	"crypto/sha1"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"mkn-backend/internal/pkg/ds"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/pkg/errors"
)

const (
	jwtPrefix = "Bearer "
	signedKey = "rfyueiwoopuih4y32kjh32"
	tokenTTL  = 12 * time.Hour
)

type tokenClaims struct {
	jwt.StandardClaims
	UserId        uuid.UUID `json:"id"`
	UserIsManager int       `json:"is_manager"`
}

type LoginReqBody struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// Login godoc
// @Summary      Login user
// @Description  Login user
// @Tags         auth
// @Produce      json
// @Param data body LoginReqBody true "User data"
// @Success      200 {object} ds.User
// @Failure 500 {object} errorResponse
// @Router      /login [post]
func (a *Application) Login(c *gin.Context) {
	req := &LoginReqBody{}

	err := json.NewDecoder(c.Request.Body).Decode(req)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "Can't decode body params")
		return
	}

	if req.Password == "" {
		newErrorResponse(c, http.StatusBadRequest, "Password is empty")
		return
	}

	if req.Username == "" {
		newErrorResponse(c, http.StatusBadRequest, "Username is empty")
		return
	}

	token, err := a.GenerateToken(req.Username, req.Password)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "Can't generate JWT")
		return
	}

	usr, err := a.repo.GetUser(req.Username, generateHashString(req.Password))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "No such user in database")
		return
	}

	a.redis.Set(*a.ctx, token, usr.Id.String(), tokenTTL)

	c.SetCookie("access_token", token, int(tokenTTL), "/", "localhost", false, false)
	c.JSON(http.StatusOK, usr)
}

// Logout godoc
// @Summary      Logout user
// @Description  Logout user
// @Tags         auth
// @Produce      json
// @Security BearerAuth
// @Success      200 {object} ds.User
// @Failure 403 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router      /logout [get]
func (a *Application) Logout(c *gin.Context) {
	jwtStr, err := a.GetJWT(c)
	if err != nil {
		newErrorResponse(c, http.StatusUnauthorized, "No JWT occured")
		return
	}

	userId, err := a.redis.Get(*a.ctx, jwtStr).Result()
	if err != nil {
		newErrorResponse(c, http.StatusUnauthorized, "No such token in repository")
		return
	}

	err = a.redis.Del(*a.ctx, jwtStr).Err()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, fmt.Sprintf("Can't delete key: %s", jwtStr))
		return
	}

	user, err := a.repo.GetUserById(userId)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "No such user in repository")
	}

	c.SetCookie("access_token", "", -1, "/", "localhost", false, false)
	c.JSON(http.StatusOK, user)
}

type SignUpReqBody struct {
	Username string `json:"username" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// SignUp godoc
// @Summary      Signup user
// @Description  Signup user
// @Tags         auth
// @Produce      json
// @Param data body SignUpReqBody true "User data"
// @Success      200 {object} ds.User
// @Failure 500 {object} errorResponse
// @Router      /signup [post]
func (a *Application) SignUp(c *gin.Context) {
	req := &SignUpReqBody{}

	err := json.NewDecoder(c.Request.Body).Decode(req)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "Can't decode body params")
		return
	}

	if req.Password == "" {
		newErrorResponse(c, http.StatusBadRequest, "Password is empty")
		return
	}

	if req.Username == "" {
		newErrorResponse(c, http.StatusBadRequest, "Username is empty")
		return
	}

	if req.Email == "" {
		newErrorResponse(c, http.StatusBadRequest, "Email is empty")
		return
	}

	user, err := a.repo.SignUp(&ds.User{
		Id:        uuid.New(),
		Username:  req.Username,
		IsManager: 0,
		Email:     req.Email,
		Password:  generateHashString(req.Password),
	})

	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "This nickname is already taken")
		return
	}

	c.JSON(http.StatusOK, user)
}

// GenerateToken
// Generate's JWT when user login
func (a *Application) GenerateToken(username, password string) (string, error) {
	user, err := a.repo.GetUser(username, generateHashString(password))
	if err != nil {
		return "", errors.Wrap(err, "Can't generate JWT")
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(tokenTTL).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		user.Id,
		user.IsManager,
	})

	return token.SignedString([]byte(signedKey))
}

// ParseToken
// Parses JWT
func (a *Application) ParseToken(accessToken string) (uuid.UUID, error) {
	token, err := jwt.ParseWithClaims(accessToken, &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("Invalid signing method")
		}

		return []byte(signedKey), nil
	})

	if err != nil {
		return uuid.Nil, err
	}

	claims, ok := token.Claims.(*tokenClaims)
	if !ok {
		return uuid.Nil, errors.New("Token claims are not  of type *tokenClaims")
	}

	return claims.UserId, nil
}

// generateHashString
// Using SHA1 to hash user's password
func generateHashString(s string) string {
	h := sha1.New()
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil))
}
