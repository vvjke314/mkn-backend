package app

import "github.com/gin-gonic/gin"

// UpdateProject godoc
// @Summary      Updates project
// @Description  Updates project
// @Tags         change
// @Produce      json
// @Success      200 {integer} 1
// @Failure 403 {object} errorResponse
// @Failure 404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router      /{project} [put]
func (a *Application) UpdateProject(c *gin.Context) {

}

// DeleteProject godoc
// @Summary      Deletes project
// @Description  Deletes project
// @Tags         delete
// @Produce      json
// @Success      200 {integer} 1
// @Failure 403 {object} errorResponse
// @Failure 404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router      /{project} [delete]
func (a *Application) DeleteProject(c *gin.Context) {

}

// CreateSection godoc
// @Summary      Creates section
// @Description  Creates section
// @Tags         add
// @Produce      json
// @Success      200 {integer} 1
// @Failure 403 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router      /{project}/section [post]
func (a *Application) CreateSection(c *gin.Context) {

}

// GetCollaborators godoc
// @Summary      Gets collaborators
// @Description  Gets collaborators
// @Tags         info
// @Produce      json
// @Success      200 {integer} 1
// @Failure 403 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router      /{project}/collaborators [get]
func (a *Application) GetCollaborators(c *gin.Context) {

}

// GetAllSections godoc
// @Summary      Gets all sections
// @Description  Gets all sections
// @Tags         info
// @Produce      json
// @Success      200 {integer} 1
// @Failure 403 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router      /{project}/sections [get]
func (a *Application) GetAllSections(c *gin.Context) {

}

// AddCollaborator godoc
// @Summary      Adds collaborators
// @Description  Adds collaborators
// @Tags         add
// @Produce      json
// @Success      200 {integer} 1
// @Failure 403 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router      /{project}/collaborator [post]
func (a *Application) AddCollaborator(c *gin.Context) {

}

// DeleteCollaborator godoc
// @Summary      Deletes collaborator
// @Description  Deletes collaborator
// @Tags         delete
// @Produce      json
// @Success      200 {integer} 1
// @Failure 403 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router      /{project}/collaborator [delete]
func (a *Application) DeleteCollaborator(c *gin.Context) {

}
