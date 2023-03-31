package app

import "github.com/gin-gonic/gin"

// CreateProject godoc
// @Summary      Creates project
// @Description  Creates project
// @Tags         add
// @Produce      json
// @Success      200 {integer} 1
// @Failure 403 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router      /project [post]
func (a *Application) CreateProject(c *gin.Context) {

}

// GetUpcomingNotifications godoc
// @Summary      Gets upcoming notifications
// @Description  Gets upcoming notifications
// @Tags         info
// @Produce      json
// @Success      200 {object} []ds.Notification
// @Failure 403 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router      /upcoming [get]
func (a *Application) GetUpcomingNotifications(c *gin.Context) {

}

// GetFavoriteProjects godoc
// @Summary      Gets favorite proj
// @Description  Gets favorite proj
// @Tags         info
// @Produce      json
// @Success      200 {object} []ds.Project
// @Failure 403 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router      /favorites [get]
func (a *Application) GetFavoriteProjects(c *gin.Context) {

}

// GetAllProjects godoc
// @Summary      Gets all proj
// @Description  Gets all proj
// @Tags         info
// @Produce      json
// @Success      200 {object} []ds.Project
// @Failure 403 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router      /projects [get]
func (a *Application) GetAllProjects(c *gin.Context) {

}

// GetAllOwnedProjects godoc
// @Summary      Gets all owned proj
// @Description  Gets all owned proj
// @Tags         info
// @Produce      json
// @Success      200 {object} []ds.Project
// @Failure 403 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router      /owned_projects [get]
func (a *Application) GetAllOwnedProjects(c *gin.Context) {

}
