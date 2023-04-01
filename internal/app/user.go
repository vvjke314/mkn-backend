package app

import "github.com/gin-gonic/gin"

// CreateProject godoc
// @Summary      Creates project
// @Description  Creates project
// @Tags         add
// @Produce      json
// @Success      200 {object} []ds.Project
// @Failure 403 {object} errorResponse
// @Failure 404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router      /project [post]
func (a *Application) CreateProject(c *gin.Context) {

}

// GetUpcomingNotifications godoc
// @Summary      Gets upcoming notifications
// @Description  Returns upcoming notifications
// @Tags         info
// @Produce      json
// @Success      200 {object} []ds.Notification
// @Failure 403 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router      /upcoming [get]
func (a *Application) GetUpcomingNotifications(c *gin.Context) {

}

// GetFavoriteProjects godoc
// @Summary      Gets favorite projects
// @Description  Returns favorite projects
// @Tags         info
// @Produce      json
// @Success      200 {object} []ds.FavoriteProject
// @Failure 403 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router      /favorites [get]
func (a *Application) GetFavoriteProjects(c *gin.Context) {

}

// GetFavoriteProject godoc
// @Summary      Gets favorite projects
// @Description  Returns favorite projects
// @Tags         info
// @Produce      json
// @Param project_id path string true "Project ID"
// @Success      200 {object} ds.FavoriteProject
// @Failure 403 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router      /favorites/:project_id [get]
func (a *Application) GetFavoriteProject(c *gin.Context) {

}

// GetAllProjects godoc
// @Summary      Gets all projects
// @Description  Returns all projects
// @Tags         info
// @Produce      json
// @Success      200 {object} []ds.Project
// @Failure 403 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router      /projects [get]
func (a *Application) GetAllProjects(c *gin.Context) {

}

// GetAllOwnedProjects godoc
// @Summary      Gets all owned project
// @Description  Gets all owned project
// @Tags         info
// @Produce      json
// @Success      200 {object} []ds.Project
// @Failure 403 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router      /projects/owned [get]
func (a *Application) GetAllOwnedProjects(c *gin.Context) {

}

// AddFavorite godoc
// @Summary      Add favorite project
// @Description  Add favorite project
// @Tags         add
// @Produce      json
// @Param project_id query string true "Project ID"
// @Success      200 {object} []ds.FavoriteProject
// @Failure 403 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router      /favorite [post]
func (a *Application) AddFavorite(c *gin.Context) {

}

// DeleteFavorite godoc
// @Summary      Delete favorite project
// @Description  Delete favorite user project
// @Tags         delete
// @Produce      json
// @Param project_id path string true "Project ID"
// @Success      200 {object} []ds.FavoriteProject
// @Failure 403 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router      /favorite/:project_id [delete]
func (a *Application) DeleteFavorite(c *gin.Context) {

}

// ChangeEmail godoc
// @Summary      Changes user email
// @Description  Changes user email
// @Tags         change
// @Produce      json
// @Success      200 {object} []ds.FavoriteProject
// @Failure 403 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router      /email [put]
func (a *Application) ChangeEmail(c *gin.Context) {

}

// LastThreeProjects godoc
// @Summary      Returns last 3 projects
// @Description  Returns the last three projects by last edit time
// @Tags         info
// @Produce      json
// @Success      200 {object} []ds.Project
// @Failure 403 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router      /projects/latest [get]
func (a *Application) LastThreeProjects(c *gin.Context) {

}
