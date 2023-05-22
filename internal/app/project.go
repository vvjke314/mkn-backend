package app

import (
	"encoding/json"
	"mkn-backend/internal/pkg/ds"
	"mkn-backend/internal/pkg/grpcApi"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
)

// UpdateProject godoc
// @Summary      Update project
// @Description  Updates a specific project according to the entered parameters and returns all owned projects
// @Tags         project
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
		newErrorResponse(c, http.StatusBadRequest, "Invalid project ID")
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
// @Tags         project
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
		newErrorResponse(c, http.StatusBadRequest, "Invalid project ID")
		return
	}

	if userId != project.OwnerId.String() {
		newErrorResponse(c, http.StatusForbidden, "You cannot modify projects that are not yours")
		return
	}

	notifications, err := a.repo.DeleteProject(projectId)
	if err != nil {
		log.Println(err)
		newErrorResponse(c, http.StatusInternalServerError, "")
		return
	}

	for i := range notifications {
		a.grpcClient.CancelNotification(*a.ctx, &grpcApi.CancelNotificationRequest{NotificationId: notifications[i].Id.String()})
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
// @Tags         section
// @Produce      json
// @Security BearerAuth
// @Param project_id path string true "Project ID"
// @Param data body ds.CreateSectionRequest true "Section information"
// @Success 200 {object} []ds.Section
// @Failure 403 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router      /project/{project_id}/section [post]
func (a *Application) CreateSection(c *gin.Context) {
	req := &ds.CreateSectionRequest{}

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
		newErrorResponse(c, http.StatusBadRequest, "Invalid project ID")
		return
	}

	if project.OwnerId.String() != userId {
		newErrorResponse(c, http.StatusForbidden, "You cannot modify projects that are not yours")
		return
	}

	err = json.NewDecoder(c.Request.Body).Decode(req)
	if err != nil {
		log.Println(err)
		newErrorResponse(c, http.StatusBadRequest, "Bad request body")
		return
	}

	section := &ds.Section{
		Id:        uuid.New(),
		Color:     req.Color,
		Title:     req.Title,
		ProjectId: uuid.MustParse(projectId),
	}

	err = a.repo.CreateSection(*section)
	if err != nil {
		log.Println(err)
		newErrorResponse(c, http.StatusInternalServerError, "Can't create section")
		return
	}

	sections, err := a.repo.GetAllSections(projectId)
	if err != nil {
		log.Println(err)
		newErrorResponse(c, http.StatusInternalServerError, "Can't get all sections")
		return
	}

	c.JSON(http.StatusOK, sections)
}

// GetCollaborators godoc
// @Summary      Returns collaborators
// @Description  Returns all collaborators of the project
// @Tags         collaborations
// @Produce      json
// @Security BearerAuth
// @Param project_id path string true "Project ID"
// @Success      200 {object} []ds.User
// @Failure 403 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router      /project/{project_id}/collaborators [get]
func (a *Application) GetCollaborators(c *gin.Context) {
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
		newErrorResponse(c, http.StatusBadRequest, "Invalid project ID")
		return
	}

	if project.OwnerId.String() != userId {
		newErrorResponse(c, http.StatusForbidden, "You cannot modify projects that are not yours")
		return
	}

	users, err := a.repo.GetAllCollaborators(projectId)
	if err != nil {
		log.Println(err)
		newErrorResponse(c, http.StatusInternalServerError, "Can't get all collaborators")
		return
	}

	c.JSON(http.StatusOK, users)
}

// GetAllSections godoc
// @Summary      Returns all sections
// @Description  Returns all sections of the current project
// @Tags         section
// @Produce      json
// @Param project_id path string true "Project ID"
// @Security BearerAuth
// @Success      200 {object} []ds.Section
// @Failure 403 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router      /project/{project_id}/sections [get]
func (a *Application) GetAllSections(c *gin.Context) {
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
		newErrorResponse(c, http.StatusBadRequest, "Invalid project ID")
		return
	}

	if project.OwnerId.String() != userId && !a.repo.IsCollaborator(userId, projectId) {
		newErrorResponse(c, http.StatusForbidden, "You cannot watch projects that are not yours")
		return
	}

	sections, err := a.repo.GetAllSections(projectId)
	if err != nil {
		log.Println(err)
		newErrorResponse(c, http.StatusInternalServerError, "Can't get all sections")
		return
	}

	c.JSON(http.StatusOK, sections)
}

// AddCollaborator godoc
// @Summary      Adds collaborators
// @Description  Adds a collaborator to the current project and returns all collaborators of this project
// @Tags         collaborations
// @Produce      json
// @Security BearerAuth
// @Param collaborator_name query string true "Collaborator nickname"
// @Param project_id path string true "Project ID"
// @Success 200 {object} []ds.User
// @Failure 403 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router      /project/{project_id}/collaborator [post]
func (a *Application) AddCollaborator(c *gin.Context) {
	userId, err := a.GetUserIdByJWT(c)
	if err != nil {
		log.Println(err)
		newErrorResponse(c, http.StatusUnauthorized, "No such authoriuzed user")
		return
	}

	projectId := c.Param("project_id")
	collab := c.Query("collaborator_name")

	if collab == "" {
		newErrorResponse(c, http.StatusBadRequest, "You need to add query param \"collaborator_id\"")
	}

	collabId, err := a.repo.GetUserByName(collab)
	if err != nil {
		log.Println(err)
		newErrorResponse(c, http.StatusBadRequest, "Invalid collaborator nickname")
		return
	}

	project, err := a.repo.GetProjectById(projectId)
	if err != nil {
		log.Println(err)
		newErrorResponse(c, http.StatusBadRequest, "Invalid project ID")
		return
	}

	if project.OwnerId.String() != userId {
		newErrorResponse(c, http.StatusBadRequest, "You cannot modify projects that are not yours")
		return
	}

	if a.repo.IsCollaborator(collabId.Id.String(), projectId) {
		newErrorResponse(c, http.StatusBadRequest, "Collaboration is already exists")
		return
	}

	err = a.repo.AddCollaborator(userId, projectId, collabId.Id.String())
	if err != nil {
		log.Println(err)
		newErrorResponse(c, http.StatusInternalServerError, "Can't add collaborator")
		return
	}

	users, err := a.repo.GetAllCollaborators(projectId)
	if err != nil {
		log.Println(err)
		newErrorResponse(c, http.StatusInternalServerError, "Can't get all collaborators")
		return
	}

	c.JSON(http.StatusOK, users)
}

// DeleteCollaborator godoc
// @Summary      Deletes collaborator
// @Description  Removes a collaborator from the current project and returns all collaborators of this project
// @Tags         collaborations
// @Produce      json
// @Security BearerAuth
// @Param collaborator_name query string true "Collaborator nickname"
// @Param project_id path string true "Project ID"
// @Success 200 {object} []ds.Collaboration
// @Failure 403 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router      /project/{project_id}/collaborator [delete]
func (a *Application) DeleteCollaborator(c *gin.Context) {
	userId, err := a.GetUserIdByJWT(c)
	if err != nil {
		log.Println(err)
		newErrorResponse(c, http.StatusUnauthorized, "No such authoriuzed user")
		return
	}

	projectId := c.Param("project_id")
	collab := c.Query("collaborator_name")

	if collab == "" {
		newErrorResponse(c, http.StatusBadRequest, "You need to add query param \"collaborator_id\"")
	}

	collabId, err := a.repo.GetUserByName(collab)
	if err != nil {
		log.Println(err)
		newErrorResponse(c, http.StatusBadRequest, "Invalid collaborator nickname")
		return
	}

	project, err := a.repo.GetProjectById(projectId)
	if err != nil {
		log.Println(err)
		newErrorResponse(c, http.StatusBadRequest, "Invalid project ID")
		return
	}

	if project.OwnerId.String() != userId {
		newErrorResponse(c, http.StatusBadRequest, "You cannot modify projects that are not yours")
		return
	}

	if !a.repo.IsCollaborator(collabId.Id.String(), projectId) {
		newErrorResponse(c, http.StatusBadRequest, "Collaboration is not exists")
		return
	}

	err = a.repo.DeleteCollaborator(userId, projectId, collabId.Id.String())
	if err != nil {
		log.Println(err)
		newErrorResponse(c, http.StatusInternalServerError, "Can't add collaborator")
		return
	}

	users, err := a.repo.GetAllCollaborators(projectId)
	if err != nil {
		log.Println(err)
		newErrorResponse(c, http.StatusInternalServerError, "Can't get all collaborators")
		return
	}

	c.JSON(http.StatusOK, users)
}
