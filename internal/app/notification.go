package app

import "github.com/gin-gonic/gin"

// UpdateNotification godoc
// @Summary      Update notifications
// @Description  Update information about a specific notification according to the entered parameters
// @Tags         change
// @Produce      json
// @Success      200 {integer} 1
// @Param project_id path string true "Project ID"
// @Param section_id path string true "Section ID"
// @Param notification_id path string true "Notification ID"
// @Param title body string false "Notification title"
// @Param description body string false "Notification description"
// @Param deadline body time false "Notification deadline"
// @Param status body string false "Notification status"
// @Param error_status body string falase "Notification error status"
// @Failure 403 {object} errorResponse
// @Failure 404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router      /{project}/{section}/{notification} [put]
func (a *Application) UpdateNotification(c *gin.Context) {

}

// DeleteNotification godoc
// @Summary      Delete notifications
// @Description  Update information about a specific notification
// @Tags         delete
// @Produce      json
// @Param project_id path string true "Project ID"
// @Param section_id path string true "Section ID"
// @Param notification_id path string true "Notification ID"
// @Success      200 {integer} 1
// @Failure 403 {object} errorResponse
// @Failure 404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router      /{project}/{section}/{notification} [delete]
func (a *Application) DeleteNotification(c *gin.Context) {

}
