package app

import (
	"encoding/json"
	"mkn-backend/internal/pkg/ds"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
)

// UpdateSection godoc
// @Summary      Updates section
// @Description  Updates a section in the current project
// @Tags         section
// @Produce      json
// @Security BearerAuth
// @Param data body ds.UpdateSectionRequest true "Section information"
// @Param section_id path string true "Section ID"
// @Success 200 {object} []ds.Section
// @Failure 403 {object} errorResponse
// @Failure 404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router      /project/section/{section_id} [put]
func (a *Application) UpdateSection(c *gin.Context) {
	userId, err := a.GetUserIdByJWT(c)
	if err != nil {
		log.Println(err)
		newErrorResponse(c, http.StatusUnauthorized, "No such authoriuzed user")
		return
	}

	log.Println(userId)
	req := &ds.UpdateSectionRequest{}

	err = json.NewDecoder(c.Request.Body).Decode(req)
	if err != nil {
		log.Println(err)
		newErrorResponse(c, http.StatusBadRequest, "Request body passed incorrectly")
		return
	}

	log.Println(req)

	sectionId := c.Param("section_id")

	log.Println(sectionId)

	if !a.repo.IsSectionOwner(userId, sectionId) {
		newErrorResponse(c, http.StatusForbidden, "You cannot change a project that does not belong to you")
		return
	}

	section, err := a.repo.GetSectionById(sectionId)
	if err != nil {
		log.Println(err)
		newErrorResponse(c, http.StatusInternalServerError, "Can't get section")
		return
	}

	if req.Color != "" {
		section.Color = req.Color
	}
	if req.Title != "" {
		section.Title = req.Title
	}

	err = a.repo.UpdateSection(section)
	if err != nil {
		log.Println(err)
		newErrorResponse(c, http.StatusInternalServerError, "Can't update project")
		return
	}

	sections, err := a.repo.GetAllSections(section.ProjectId.String())
	if err != nil {
		log.Println(err)
		newErrorResponse(c, http.StatusInternalServerError, "Can't get all projects")
		return
	}

	c.JSON(http.StatusOK, sections)
}

// DeleteSection godoc
// @Summary      Deletes section
// @Description  Deletes section from current project
// @Tags         section
// @Produce      json
// @Security BearerAuth
// @Param section_id path string true "Section ID"
// @Success 200 {object} []ds.Section
// @Failure 403 {object} errorResponse
// @Failure 404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router      /project/section/{section_id} [delete]
func (a *Application) DeleteSection(c *gin.Context) {
	userId, err := a.GetUserIdByJWT(c)
	if err != nil {
		log.Println(err)
		newErrorResponse(c, http.StatusUnauthorized, "No such authoriuzed user")
		return
	}

	sectionId := c.Param("section_id")

	if !a.repo.IsSectionOwner(userId, sectionId) {
		newErrorResponse(c, http.StatusForbidden, "You cannot change a project that does not belong to you")
		return
	}

	section, err := a.repo.GetSectionById(sectionId)
	if err != nil {
		log.Println(err)
		newErrorResponse(c, http.StatusInternalServerError, "Can't get section")
		return
	}

	err = a.repo.DeleteSection(section)
	if err != nil {
		log.Println(err)
		newErrorResponse(c, http.StatusInternalServerError, "Can't delete section")
		return
	}

	sections, err := a.repo.GetAllSections(section.ProjectId.String())
	if err != nil {
		log.Println(err)
		newErrorResponse(c, http.StatusInternalServerError, "Can't get all projects")
		return
	}

	c.JSON(http.StatusOK, sections)
}

// CreateNotification godoc
// @Summary      Creates notification
// @Description  Creates notification in accordance with the entered parameters
// @Tags         notification
// @Produce      json
// @Security BearerAuth
// @Param section_id path string true "Section ID"
// @Param data body ds.CreateNotificationRequest true "Notification information"
// @Success 200 {object} []ds.Notification
// @Failure 403 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router      /project/section/{section_id}/notification [post]
func (a *Application) CreateNotification(c *gin.Context) {
	req := &ds.CreateNotificationRequest{}

	userId, err := a.GetUserIdByJWT(c)
	if err != nil {
		log.Println(err)
		newErrorResponse(c, http.StatusUnauthorized, "No such authoriuzed user")
		return
	}

	sectionId := c.Param("section_id")

	section, err := a.repo.GetSectionById(sectionId)
	if err != nil {
		log.Println(err)
		newErrorResponse(c, http.StatusBadRequest, "Invalid section ID")
		return
	}

	if !a.repo.IsSectionOwner(userId, sectionId) {
		newErrorResponse(c, http.StatusForbidden, "You cannot change a project that does not belong to you")
		return
	}

	err = json.NewDecoder(c.Request.Body).Decode(req)
	if err != nil {
		log.Println(err)
		newErrorResponse(c, http.StatusBadRequest, "Bad request body")
		return
	}

	notification := &ds.Notification{
		Id:          uuid.New(),
		SectionId:   section.Id,
		Title:       req.Title,
		Description: req.Description,
		Deadline:    req.Deadline,
		Status:      "scheduled",
		ErrorStatus: 0,
	}

	err = a.repo.CreateNotification(*notification)
	if err != nil {
		log.Println(err)
		newErrorResponse(c, http.StatusInternalServerError, "Can't create notification")
		return
	}

	notifications, err := a.repo.GetAllNotifications(sectionId)
	if err != nil {
		log.Println(err)
		newErrorResponse(c, http.StatusInternalServerError, "Can't get all notifications")
		return
	}

	c.JSON(http.StatusOK, notifications)
}

// GetAllNotifications godoc
// @Summary      Gets All Notifications
// @Description  Returns all notifications in the current section
// @Tags         notification
// @Security BearerAuth
// @Produce      json
// @Param section_id path string true "Section ID"
// @Success      200 {object} []ds.Notification
// @Failure 403 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router      /project/section/{section_id}/notifications [get]
func (a *Application) GetAllNotifications(c *gin.Context) {
	userId, err := a.GetUserIdByJWT(c)
	if err != nil {
		log.Println(err)
		newErrorResponse(c, http.StatusUnauthorized, "No such authoriuzed user")
		return
	}

	sectionId := c.Param("section_id")

	if !a.repo.IsSectionOwner(userId, sectionId) {
		newErrorResponse(c, http.StatusForbidden, "You cannot change a project that does not belong to you")
		return
	}

	notifications, err := a.repo.GetAllNotifications(sectionId)
	if err != nil {
		log.Println(err)
		newErrorResponse(c, http.StatusInternalServerError, "Can't get all notifications")
		return
	}

	c.JSON(http.StatusOK, notifications)

}

// GetNotification godoc
// @Summary      Gets Notification
// @Description  Returns Notification by ID
// @Tags         notification
// @Security BearerAuth
// @Produce      json
// @Param notification_id path string true "Notification ID"
// @Success      200 {object} ds.Notification
// @Failure 403 {object} errorResponse
// @Failure 404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router      /project/section/notification/{notification_id} [get]
func (a *Application) GetNotification(c *gin.Context) {
	userId, err := a.GetUserIdByJWT(c)
	if err != nil {
		log.Println(err)
		newErrorResponse(c, http.StatusUnauthorized, "No such authoriuzed user")
		return
	}

	notificationId := c.Param("notification_id")

	if !a.repo.IsNotificationOwner(userId, notificationId) {
		newErrorResponse(c, http.StatusForbidden, "You cannot change a project that does not belong to you")
		return
	}

	notifications, err := a.repo.GetNotification(notificationId)
	if err != nil {
		log.Println(err)
		newErrorResponse(c, http.StatusInternalServerError, "Can't get all notifications")
		return
	}

	c.JSON(http.StatusOK, notifications)
}
