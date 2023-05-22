package app

import (
	"encoding/json"
	"log"
	"mkn-backend/internal/pkg/ds"
	"mkn-backend/internal/pkg/grpcApi"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

// UpdateNotification godoc
// @Summary      Update notifications
// @Description  Update information about a specific notification according to the entered parameters
// @Tags         notification
// @Produce      json
// @Security BearerAuth
// @Param notification_id path string true "Notification ID"
// @Param data body ds.UpdateNotificationRequest true "Notification information"
// @Success 200 {object} []ds.Notification
// @Failure 403 {object} errorResponse
// @Failure 404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router      /api/project/section/notification/{notification_id} [put]
func (a *Application) UpdateNotification(c *gin.Context) {
	req := &ds.UpdateNotificationRequest{}

	userId, err := a.GetUserIdByJWT(c)
	if err != nil {
		log.Println(err)
		newErrorResponse(c, http.StatusUnauthorized, "No such authoriuzed user")
		return
	}

	notificationId := c.Param("notification_id")

	notification, err := a.repo.GetNotificationById(notificationId)
	if err != nil {
		log.Println(err)
		newErrorResponse(c, http.StatusBadRequest, "Invalid notification ID")
		return
	}

	if !a.repo.IsNotificationOwner(userId, notificationId) {
		newErrorResponse(c, http.StatusForbidden, "You cannot change a project that does not belong to you")
		return
	}

	err = json.NewDecoder(c.Request.Body).Decode(req)
	if err != nil {
		log.Println(err)
		newErrorResponse(c, http.StatusBadRequest, "Bad request body")
		return
	}

	log.Println(req.DeadLine)

	if req.Title != "" {
		notification.Title = req.Title
	}
	if req.Description != "" {
		notification.Description = req.Description
	}
	if req.DeadLine != "" {
		log.Println(req.DeadLine)
		notification.Deadline, _ = time.Parse(time.RFC3339, req.DeadLine)
	}

	log.Println(notification.Deadline)

	err = a.repo.UpdateNotification(notification)
	if err != nil {
		log.Println(err)
		newErrorResponse(c, http.StatusInternalServerError, "Can't change notification")
		return
	}

	deadline := notification.Deadline.Unix()
	deadlineStr := strconv.Itoa(int(deadline))
	a.grpcClient.ScheduleNotification(*a.ctx, &grpcApi.ScheduleRequest{NotificationId: notification.Id.String(), Deadline: deadlineStr})

	notifications, err := a.repo.GetAllNotifications(notification.SectionId.String())
	if err != nil {
		log.Println(err)
		newErrorResponse(c, http.StatusInternalServerError, "Can't get all notifications")
		return
	}

	c.JSON(http.StatusOK, notifications)
}

// DeleteNotification godoc
// @Summary      Delete notifications
// @Description  Update information about a specific notification
// @Tags         notification
// @Produce      json
// @Security BearerAuth
// @Param notification_id path string true "Notification ID"
// @Success 200 {object} []ds.Notification
// @Failure 403 {object} errorResponse
// @Failure 404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router      /api/project/section/notification/{notification_id} [delete]
func (a *Application) DeleteNotification(c *gin.Context) {
	userId, err := a.GetUserIdByJWT(c)
	if err != nil {
		log.Println(err)
		newErrorResponse(c, http.StatusUnauthorized, "No such authoriuzed user")
		return
	}

	notificationId := c.Param("notification_id")

	notification, err := a.repo.GetNotificationById(notificationId)
	if err != nil {
		log.Println(err)
		newErrorResponse(c, http.StatusBadRequest, "Invalid notification ID")
		return
	}

	if !a.repo.IsNotificationOwner(userId, notificationId) {
		newErrorResponse(c, http.StatusForbidden, "You cannot change a project that does not belong to you")
		return
	}

	err = a.repo.DeleteNotification(notification)
	if err != nil {
		log.Println(err)
		newErrorResponse(c, http.StatusInternalServerError, "Can't delete notification")
		return
	}

	notifications, err := a.repo.GetAllNotifications(notification.SectionId.String())
	if err != nil {
		log.Println(err)
		newErrorResponse(c, http.StatusInternalServerError, "Can't get all notifications")
		return
	}

	a.grpcClient.CancelNotification(*a.ctx, &grpcApi.CancelNotificationRequest{NotificationId: notificationId})
	c.JSON(http.StatusOK, notifications)
}

// ResendNotification godoc
// @Summary      Resend notification
// @Description  Resend notification
// @Tags         notification
// @Produce      json
// @Security BearerAuth
// @Param data body ds.ResendNotificationRequest true "Deadline info"
// @Param notification_id path string true "Notification ID"
// @Success 200 {object} []ds.Notification
// @Failure 403 {object} errorResponse
// @Failure 404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router      /api/project/section/notification/resend/{notification_id} [put]
func (a *Application) ResendNotification(c *gin.Context) {
	req := &ds.ResendNotificationRequest{}

	userId, err := a.GetUserIdByJWT(c)
	if err != nil {
		log.Println(err)
		newErrorResponse(c, http.StatusUnauthorized, "No such authoriuzed user")
		return
	}

	notificationId := c.Param("notification_id")

	notification, err := a.repo.GetNotificationById(notificationId)
	if err != nil {
		log.Println(err)
		newErrorResponse(c, http.StatusBadRequest, "Invalid notification ID")
		return
	}

	if !a.repo.IsNotificationOwner(userId, notificationId) {
		newErrorResponse(c, http.StatusForbidden, "You cannot change a project that does not belong to you")
		return
	}

	err = json.NewDecoder(c.Request.Body).Decode(req)
	if err != nil {
		log.Println(err)
		newErrorResponse(c, http.StatusBadRequest, "Bad request body")
		return
	}

	if req.Deadline != "" {
		notification.Deadline, _ = time.Parse(time.RFC3339, req.Deadline)
	}

	if notification.Status == "undelivered" {
		notification.Status = "scheduled"
	}

	err = a.repo.UpdateNotification(notification)
	if err != nil {
		log.Println(err)
		newErrorResponse(c, http.StatusInternalServerError, "Can't change notification")
		return
	}

	deadline := notification.Deadline.Unix()
	deadlineStr := strconv.Itoa(int(deadline))
	a.grpcClient.ScheduleNotification(*a.ctx, &grpcApi.ScheduleRequest{NotificationId: notification.Id.String(), Deadline: deadlineStr})

	notifications, err := a.repo.GetAllNotifications(notification.SectionId.String())
	if err != nil {
		log.Println(err)
		newErrorResponse(c, http.StatusInternalServerError, "Can't get all notifications")
		return
	}

	c.JSON(http.StatusOK, notifications)
}

// undelivered -> scheduled
