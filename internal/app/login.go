package app

import "github.com/gin-gonic/gin"

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

// Logout godoc
// @Summary      Signup user
// @Description  Signup user
// @Tags         auth
// @Produce      json
// @Success      200 {object} ds.User
// @Failure 500 {object} errorResponse
// @Router      /signup [post]
func (a *Application) SignUp(c *gin.Context) {

}
