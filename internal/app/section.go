package app

import "github.com/gin-gonic/gin"

// UpdateSection godoc
// @Summary      Updates section
// @Description  Updates section
// @Tags         change
// @Produce      json
// @Success      200 {integer} 1
// @Failure 403 {object} errorResponse
// @Failure 404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router      /{project}/{section} [put]
func (a *Application) UpdateSection(c *gin.Context) {

}

// DeleteSection godoc
// @Summary      Deletes section
// @Description  Deletes section
// @Tags         delete
// @Produce      json
// @Success      200 {integer} 1
// @Failure 403 {object} errorResponse
// @Failure 404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router      /{project}/{section} [delete]
func (a *Application) DeleteSection(c *gin.Context) {

}

// CreateNotification godoc
// @Summary      Creates notification
// @Description  Creates notification
// @Tags         add
// @Produce      json
// @Success      200 {integer} 1
// @Failure 403 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router      /{project}/{section}/notification [post]
func (a *Application) CreateNotification(c *gin.Context) {

}

// GetAllNotifications godoc
// @Summary      Gets All Notifications
// @Description  Gets All Notifications
// @Tags         info
// @Produce      json
// @Success      200 {object} []ds.Notification
// @Failure 403 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router      /{project}/{section}/notifications [get]
func (a *Application) GetAllNotifications(c *gin.Context) {

}

// GetNotification godoc
// @Summary      Gets Notification
// @Description  Gets Notification by ID
// @Tags         info
// @Produce      json
// @Success      200 {object} ds.Notification
// @Failure 403 {object} errorResponse
// @Failure 404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router      /{project}/{section}/{notification} [get]
func (a *Application) GetNotification(c *gin.Context) {

}
