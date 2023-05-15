package app

import "github.com/gin-gonic/gin"

// UpdateNotification godoc
// @Summary      Update notifications
// @Description  Update information about a specific notification according to the entered parameters
// @Tags         notification
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
// @Tags         notification
// @Produce      json
// @Param notification_id path string true "Notification ID"
// @Success 200 {object} []ds.Notification
// @Failure 403 {object} errorResponse
// @Failure 404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router      /project/section/notification/{notification_id} [delete]
func (a *Application) DeleteNotification(c *gin.Context) {

}

// ResendNotification godoc
// @Summary      Resend notification
// @Description  Resend notification
// @Tags         notification
// @Produce      json
// @Security BearerAuth
// @Param data body ds.ResendNotificationRequest true "Section information"
// @Param notification_id path string true "Notification ID"
// @Success 200 {object} []ds.Notification
// @Failure 403 {object} errorResponse
// @Failure 404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router      /project/section/notification/resend/{notification_id} [put]
func (a *Application) ResendNotification(c *gin.Context) {

}
