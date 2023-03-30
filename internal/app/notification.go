package app

import "github.com/gin-gonic/gin"

// UpdateNotification godoc
// @Summary      Updates notifications
// @Description  Updates notifications
// @Tags         change
// @Produce      json
// @Success      200 {integer} 1
// @Failure 500 {object} errorResponse
// @Router      /{project}/{section}/{notification} [put]
func (a *Application) UpdateNotification(c *gin.Context) {

}

// DeleteNotification godoc
// @Summary      Deletes notifications
// @Description  Deletes notifications
// @Tags         delete
// @Produce      json
// @Success      200 {integer} 1
// @Failure 500 {object} errorResponse
// @Router      /{project}/{section}/{notification} [delete]
func (a *Application) DeleteNotification(c *gin.Context) {

}
