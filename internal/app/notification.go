package app

import "github.com/gin-gonic/gin"

// UpdateNotification godoc
// @Summary      Update notifications
// @Description  Update information about a specific notification according to the entered parameters
// @Tags         change
// @Produce      json
// @Param notification_id path string true "Notification ID"
// @Success 200 {object} []ds.Notification
// @Failure 403 {object} errorResponse
// @Failure 404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router      /project/section/notification/{notification_id} [put]
func (a *Application) UpdateNotification(c *gin.Context) {

}

// DeleteNotification godoc
// @Summary      Delete notifications
// @Description  Update information about a specific notification
// @Tags         delete
// @Produce      json
// @Param notification_id path string true "Notification ID"
// @Success 200 {object} []ds.Notification
// @Failure 403 {object} errorResponse
// @Failure 404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router      /project/section/notification/{notification_id} [delete]
func (a *Application) DeleteNotification(c *gin.Context) {

}
