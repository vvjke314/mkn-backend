package app

import (
	"encoding/json"
	"mkn-backend/internal/pkg/ds"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
)

// CreateProject godoc
// @Summary      Creates project
// @Description  Creates project
// @Tags         add
// @Produce      json
// @Param data body ds.Project true "Project data"
// @Security BearerAuth
// @Success      200 {object} ds.Project
// @Failure 403 {object} errorResponse
// @Failure 404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router      /project [post]
func (a *Application) CreateProject(c *gin.Context) {
	req := &ds.Project{}

	userId, err := a.GetUserIdByJWT(c)
	if err != nil {
		log.Println(err)
		newErrorResponse(c, http.StatusUnauthorized, "No such authoriuzed user")
		return
	}

	err = json.NewDecoder(c.Request.Body).Decode(req)
	if err != nil {
		log.Println(err)
		newErrorResponse(c, http.StatusBadRequest, "Bad request")
		return
	}

	req.Id = uuid.New()
	req.OwnerId, err = uuid.Parse(userId)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "Can't format uuid correctly")
		return
	}
	req.LastEdited = time.Now().UTC()

	pct, err := a.repo.CreateProject(req)
	if err != nil {
		log.Println(err)
		newErrorResponse(c, http.StatusBadRequest, "Can't create project")
		return
	}

	c.JSON(http.StatusOK, pct)
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
// @Router      /favorites/{project_id} [get]
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
	projects, err := a.repo.GetAllProjects()
	if err != nil {
		log.Println(err)
		newErrorResponse(c, http.StatusBadRequest, "Can't get all projects")
		return
	}

	c.JSON(http.StatusOK, projects)
}

// GetAllOwnedProjects godoc
// @Summary      Gets all owned project
// @Description  Gets all owned project
// @Tags         info
// @Produce      json
// @Security BearerAuth
// @Success      200 {object} []ds.Project
// @Failure 403 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router      /projects/owned [get]
func (a *Application) GetAllOwnedProjects(c *gin.Context) {
	userId, err := a.GetUserIdByJWT(c)
	if err != nil {
		log.Println(err)
		newErrorResponse(c, http.StatusUnauthorized, "No such authoriuzed user")
		return
	}

	projects, err := a.repo.GetAllOwnedProjects(userId)
	if err != nil {
		log.Println(err)
		newErrorResponse(c, http.StatusBadRequest, "Can't get all owned projects")
		return
	}

	c.JSON(http.StatusOK, projects)
}

// AddFavorite godoc
// @Summary      Add favorite project
// @Description  Add favorite project
// @Tags         add
// @Produce      json
// @Param project_id query string true "Project ID"
// @Security BearerAuth
// @Success      200 {object} []ds.FavoriteProject
// @Failure 403 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router      /favorite [post]
func (a *Application) AddFavorite(c *gin.Context) {
	userId, err := a.GetUserIdByJWT(c)
	if err != nil {
		log.Println(err)
		newErrorResponse(c, http.StatusUnauthorized, "No such authoriuzed user")
		return
	}

	projectId := c.Query("project_id")
	err = a.repo.AddFavorite(userId, projectId)
	if err != nil {
		log.Println(err)
		newErrorResponse(c, http.StatusUnauthorized, "Can't add favorite project")
		return
	}

	favoriteProjects, err := a.repo.GetFavoriteProjects(userId)
	if err != nil {
		log.Println(err)
		newErrorResponse(c, http.StatusUnauthorized, "Can't return favorite projects")
		return
	}

	c.JSON(http.StatusOK, favoriteProjects)
}

// DeleteFavorite godoc
// @Summary      Delete favorite project
// @Description  Delete favorite user project
// @Tags         delete
// @Produce      json
// @Param project_id path string true "Project ID"
// @Security BearerAuth
// @Success      200 {object} []ds.FavoriteProject
// @Failure 403 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router      /favorite/{project_id} [delete]
func (a *Application) DeleteFavorite(c *gin.Context) {
	userId, err := a.GetUserIdByJWT(c)
	if err != nil {
		log.Println(err)
		newErrorResponse(c, http.StatusUnauthorized, "No such authoriuzed user")
		return
	}
	projectId := c.Param("project_id")

	err = a.repo.DeleteFavorite(userId, projectId)
	if err != nil {
		log.Println(err)
		newErrorResponse(c, http.StatusBadRequest, "Can't delete user")
		return
	}

	favoriteProjects, err := a.repo.GetFavoriteProjects(userId)
	if err != nil {
		log.Println(err)
		newErrorResponse(c, http.StatusUnauthorized, "Can't return favorite projects")
		return
	}

	c.JSON(http.StatusOK, favoriteProjects)
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
