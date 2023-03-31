package app

import "github.com/gin-gonic/gin"

// UpdateProject godoc
// @Summary      Update project
// @Description  Updates a specific project according to the entered parameters
// @Tags         change
// @Produce      json
// @Success 200 {object} []ds.Project
// @Param project_id path string true "Project ID"
// @Failure 403 {object} errorResponse
// @Failure 404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router      /project/{project_id} [put]
func (a *Application) UpdateProject(c *gin.Context) {

}

// DeleteProject godoc
// @Summary      Deletes project
// @Description  Deletes a specific project
// @Tags         delete
// @Produce      json
// @Param project_id path string true "Project ID"
// @Success 200 {object} []ds.Project
// @Failure 403 {object} errorResponse
// @Failure 404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router      /project/{project_id} [delete]
func (a *Application) DeleteProject(c *gin.Context) {

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
