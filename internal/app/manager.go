package app

import "github.com/gin-gonic/gin"

// GetUndeliviredNotifications godoc
// @Summary      Gets undelivired notifications
// @Description  Gets undelivired notifications
// @Tags         info
// @Produce      json
// @Success      200 {integer} 1
// @Failure 403 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router      /undelivered_notifications [get]
func (a *Application) GetUndeliviredNotifications(c *gin.Context) {

}
