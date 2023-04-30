package app

import (
	"crypto/sha1"
	"encoding/hex"
	"encoding/json"
	"errors"
	"mkn-backend/internal/pkg/ds"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
)

type SignUpReqBody struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

// Login godoc
// @Summary      Login user
// @Description  Login user
// @Tags         auth
// @Produce      json
// @Success      200 {object} ds.User
// @Failure 500 {object} errorResponse
// @Router      /login [post]
func (a *Application) Login(c *gin.Context) {

}

// Logout godoc
// @Summary      Logout user
// @Description  Logout user
// @Tags         auth
// @Produce      json
// @Success      200 {object} ds.User
// @Failure 403 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router      /logout [get]
func (a *Application) Logout(c *gin.Context) {

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
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	if req.Password == "" {
		c.AbortWithError(http.StatusBadRequest, errors.New("Password is empty"))
		return
	}

	if req.Username == "" {
		c.AbortWithError(http.StatusBadRequest, errors.New("Username is empty"))
		return
	}

	if req.Email == "" {
		c.AbortWithError(http.StatusBadRequest, errors.New("Email is empty"))
	}

	user, err := a.repo.SignUp(&ds.User{
		Id:        uuid.New(),
		Username:  req.Username,
		IsManager: 0,
		Email:     req.Email,
		Password:  generateHashString(req.Password),
	})

	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, user)
}

func generateHashString(s string) string {
	h := sha1.New()
	h.Write([]byte(s))
	log.Println(h)
	return hex.EncodeToString(h.Sum(nil))
}
