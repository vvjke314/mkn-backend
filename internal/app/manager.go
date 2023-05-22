package app

import "github.com/gin-gonic/gin"

// GetUndeliviredNotifications godoc
// @Summary      Gets undelivired notifications
// @Description  Allows the manager to view undelivered notifications
// @Tags         notification
// @Produce      json
// @Success      200 {object} []ds.Notification
// @Failure 403 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router      /api/undelivered_notifications [get]
func (a *Application) GetUndeliviredNotifications(c *gin.Context) {

}
