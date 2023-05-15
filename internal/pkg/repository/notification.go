package repository

import (
	"log"
	"mkn-backend/internal/pkg/ds"

	"github.com/pkg/errors"
)

func (r *Repository) DeleteNotification(notification ds.Notification) error {
	err := r.db.Delete(&notification).Error
	if err != nil {
		return errors.Wrap(err, "Can't delete notification in repo")
	}

	return nil
}

func (r *Repository) GetNotificationById(notificationId string) (ds.Notification, error) {
	notification := ds.Notification{}
	err := r.db.Where("id = ?", notificationId).First(&notification).Error
	if err != nil {
		return ds.Notification{}, errors.Wrap(err, "Can't find such notification in repo")
	}

	return notification, nil
}

func (r *Repository) IsNotificationOwner(userId, notificationId string) bool {
	notification, err := r.GetNotificationById(notificationId)
	if err != nil {
		log.Println(err)
		return false
	}

	section, err := r.GetSectionById(notification.SectionId.String())
	if err != nil {
		log.Println(err)
		return false
	}

	project, err := r.GetProjectById(section.ProjectId.String())
	if err != nil {
		log.Println(err)
		return false
	}

	if userId != project.OwnerId.String() {
		return false
	}

	return true
}
