package app

import "github.com/gin-gonic/gin"

// UpdateSection godoc
// @Summary      Updates section
// @Description  Updates a section in the current project
// @Tags         change
// @Produce      json
// @Success      200 {integer} 1
// @Param project_id path string true "Project ID"
// @Param section_id path string true "Section ID"
// @Param title body string false "Section title"
// @Param color body string false "Section color"
// @Failure 403 {object} errorResponse
// @Failure 404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router      /{project}/{section} [put]
func (a *Application) UpdateSection(c *gin.Context) {

}

// DeleteSection godoc
// @Summary      Deletes section
// @Description  Deletes section from current project
// @Tags         delete
// @Produce      json
// @Param project_id path string true "Project ID"
// @Param section_id path string true "Section ID"
// @Success      200 {integer} 1
// @Failure 403 {object} errorResponse
// @Failure 404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router      /{project}/{section} [delete]
func (a *Application) DeleteSection(c *gin.Context) {

}

// CreateNotification godoc
// @Summary      Creates notification
// @Description  Creates notification in accordance with the entered parameters
// @Tags         add
// @Produce      json
// @Param project_id path string true "Project ID"
// @Param section_id path string true "Section ID"
// @Param title body string false "Notification title"
// @Param description body string false "Notification description"
// @Param deadline body time false "Notification deadline"
// @Param status body string false "Notification status"
// @Param error_status body string falase "Notification error status"
// @Success      200 {integer} 1
// @Failure 403 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router      /{project}/{section}/notification [post]
func (a *Application) CreateNotification(c *gin.Context) {

}

// GetAllNotifications godoc
// @Summary      Gets All Notifications
// @Description  Returns all notifications in the current section
// @Tags         info
// @Produce      json
// @Param project_id path string true "Project ID"
// @Param section_id path string true "Section ID"
// @Success      200 {object} []ds.Notification
// @Failure 403 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router      /{project}/{section}/notifications [get]
func (a *Application) GetAllNotifications(c *gin.Context) {

}

// GetNotification godoc
// @Summary      Gets Notification
// @Description  Returns Notification by ID
// @Tags         info
// @Produce      json
// @Param project_id path string true "Project ID"
// @Param section_id path string true "Section ID"
// @Param notification_id path string true "Notification ID"
// @Success      200 {object} ds.Notification
// @Failure 403 {object} errorResponse
// @Failure 404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router      /{project}/{section}/{notification} [get]
func (a *Application) GetNotification(c *gin.Context) {

}
