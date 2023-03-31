package app

import "github.com/gin-gonic/gin"

// UpdateProject godoc
// @Summary      Update project
// @Description  Updates a specific project according to the entered parameters
// @Tags         change
// @Produce      json
// @Success      200 {integer} 1
// @Param id path string true "Project ID"
// @Param owner_id body string false "Owner ID"
// @Param title body string false "Project title"
// @Param description body string false "Project description"
// @Failure 403 {object} errorResponse
// @Failure 404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router      /{project} [put]
func (a *Application) UpdateProject(c *gin.Context) {

}

// DeleteProject godoc
// @Summary      Deletes project
// @Description  Deletes a specific project
// @Tags         delete
// @Produce      json
// @Param id path string true "Project ID"
// @Success      200 {integer} 1
// @Failure 403 {object} errorResponse
// @Failure 404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router      /{project} [delete]
func (a *Application) DeleteProject(c *gin.Context) {

}

// CreateSection godoc
// @Summary      Creates section
// @Description  Creates a section in the project
// @Tags         add
// @Produce      json
// @Param id path string true "Project ID"
// @Param title body string false "Section title"
// @Param color body string false "Section color"
// @Success      200 {integer} 1
// @Failure 403 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router      /{project}/section [post]
func (a *Application) CreateSection(c *gin.Context) {

}

// GetCollaborators godoc
// @Summary      Returns collaborators
// @Description  Returns all collaborators of the project
// @Tags         info
// @Produce      json
// @Param id path string true "Project ID"
// @Success      200 {object} ds.Collaborations
// @Failure 403 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router      /{project}/collaborators [get]
func (a *Application) GetCollaborators(c *gin.Context) {

}

// GetAllSections godoc
// @Summary      Returns all sections
// @Description  Returns all sections of the current project
// @Tags         info
// @Produce      json
// @Param id path string true "Project ID"
// @Success      200 {object} []ds.Section
// @Failure 403 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router      /{project}/sections [get]
func (a *Application) GetAllSections(c *gin.Context) {

}

// AddCollaborator godoc
// @Summary      Adds collaborators
// @Description  Adds a collaborator to the current project
// @Tags         add
// @Produce      json
// @Param id path string true "Project ID"
// @Success      200 {integer} 1
// @Failure 403 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router      /{project}/collaborator [post]
func (a *Application) AddCollaborator(c *gin.Context) {

}

// DeleteCollaborator godoc
// @Summary      Deletes collaborator
// @Description  Removes a collaborator from the current project
// @Tags         delete
// @Produce      json
// @Param id path string true "Project ID"
// @Success      200 {integer} 1
// @Failure 403 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router      /{project}/collaborator [delete]
func (a *Application) DeleteCollaborator(c *gin.Context) {

}
