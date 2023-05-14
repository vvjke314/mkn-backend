package app

import (
	"encoding/json"
	"mkn-backend/internal/pkg/ds"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

// UpdateProject godoc
// @Summary      Update project
// @Description  Updates a specific project according to the entered parameters and returns all owned projects
// @Tags         change
// @Produce      json
// @Success 200 {object} []ds.Project
// @Security BearerAuth
// @Param project_id path string true "Project ID"
// @Param data body ds.UpdateProjectRequest true "New project information"
// @Failure 403 {object} errorResponse
// @Failure 404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router      /project/{project_id} [put]
func (a *Application) UpdateProject(c *gin.Context) {
	userId, err := a.GetUserIdByJWT(c)
	if err != nil {
		log.Println(err)
		newErrorResponse(c, http.StatusUnauthorized, "No such authoriuzed user")
		return
	}

	req := &ds.UpdateProjectRequest{}

	err = json.NewDecoder(c.Request.Body).Decode(req)
	if err != nil {
		log.Println(err)
		newErrorResponse(c, http.StatusBadRequest, "Request body passed incorrectly")
		return
	}

	projectId := c.Param("project_id")

	project, err := a.repo.GetProjectById(projectId)
	if err != nil {
		log.Println(err)
		newErrorResponse(c, http.StatusBadRequest, "Invalid project id")
		return
	}

	if userId != project.OwnerId.String() {
		newErrorResponse(c, http.StatusForbidden, "You cannot modify projects that are not yours")
		return
	}

	if req.Color != "" {
		project.Color = req.Color
	}
	if req.Title != "" {
		project.Title = req.Title
	}
	if req.Description != "" {
		project.Description = req.Description
	}

	project.LastEdited = time.Now().UTC()

	err = a.repo.UpdateProject(userId, projectId, project)
	if err != nil {
		log.Println(err)
		newErrorResponse(c, http.StatusInternalServerError, "Can't update project")
		return
	}

	projects, err := a.repo.GetAllOwnedProjects(userId)
	if err != nil {
		log.Println(err)
		newErrorResponse(c, http.StatusInternalServerError, "Can;t return all projects")
		return
	}

	c.JSON(http.StatusOK, projects)
}

// DeleteProject godoc
// @Summary      Deletes project
// @Description  Deletes a specific project and returns all owned projects
// @Tags         delete
// @Produce      json
// @Param project_id path string true "Project ID"
// @Security BearerAuth
// @Success 200 {object} []ds.Project
// @Failure 403 {object} errorResponse
// @Failure 404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router      /project/{project_id} [delete]
func (a *Application) DeleteProject(c *gin.Context) {
	userId, err := a.GetUserIdByJWT(c)
	if err != nil {
		log.Println(err)
		newErrorResponse(c, http.StatusUnauthorized, "No such authoriuzed user")
		return
	}

	projectId := c.Param("project_id")

	project, err := a.repo.GetProjectById(projectId)
	if err != nil {
		log.Println(err)
		newErrorResponse(c, http.StatusBadRequest, "Invalid project id")
		return
	}

	if userId != project.OwnerId.String() {
		newErrorResponse(c, http.StatusForbidden, "You cannot modify projects that are not yours")
		return
	}

	err = a.repo.DeleteProject(projectId)
	if err != nil {
		log.Println(err)
		newErrorResponse(c, http.StatusInternalServerError, "")
		return
	}

	projects, err := a.repo.GetAllOwnedProjects(userId)
	if err != nil {
		log.Println(err)
		newErrorResponse(c, http.StatusInternalServerError, "Can;t return all projects")
		return
	}

	c.JSON(http.StatusOK, projects)
}

// CreateSection godoc
// @Summary      Creates section
// @Description  Creates a section in the project
// @Tags         add
// @Produce      json
// @Param project_id path string true "Project ID"
// @Success 200 {object} []ds.Section
// @Failure 403 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router      /project/{project_id}/section [post]
func (a *Application) CreateSection(c *gin.Context) {

}

// GetCollaborators godoc
// @Summary      Returns collaborators
// @Description  Returns all collaborators of the project
// @Tags         info
// @Produce      json
// @Param project_id path string true "Project ID"
// @Success      200 {object} []ds.Collaboration
// @Failure 403 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router      /project/{project_id}/collaborators [get]
func (a *Application) GetCollaborators(c *gin.Context) {

}

// GetAllSections godoc
// @Summary      Returns all sections
// @Description  Returns all sections of the current project
// @Tags         info
// @Produce      json
// @Param project_id path string true "Project ID"
// @Success      200 {object} []ds.Section
// @Failure 403 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router      /project/{project_id}/sections [get]
func (a *Application) GetAllSections(c *gin.Context) {

}

// AddCollaborator godoc
// @Summary      Adds collaborators
// @Description  Adds a collaborator to the current project
// @Tags         add
// @Produce      json
// @Param project_id path string true "Project ID"
// @Success 200 {object} []ds.Collaboration
// @Failure 403 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router      /project/{project_id}/collaborator [post]
func (a *Application) AddCollaborator(c *gin.Context) {

}

// DeleteCollaborator godoc
// @Summary      Deletes collaborator
// @Description  Removes a collaborator from the current project
// @Tags         delete
// @Produce      json
// @Param project_id path string true "Project ID"
// @Success 200 {object} []ds.Collaboration
// @Failure 403 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router      /project/{project_id}/collaborator [delete]
func (a *Application) DeleteCollaborator(c *gin.Context) {

}
