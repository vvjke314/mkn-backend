package repository

import (
	"log"
	"mkn-backend/internal/pkg/ds"

	"github.com/pkg/errors"
)

func (r *Repository) GetSectionById(sectionId string) (ds.Section, error) {
	section := ds.Section{}
	err := r.db.Where("id = ?", sectionId).First(&section).Error
	if err != nil {
		return ds.Section{}, errors.Wrap(err, "Can't find such section in repo")
	}

	return section, nil
}

// Bad error handling :)
func (r *Repository) IsSectionOwner(userId, sectionId string) bool {

	section, err := r.GetSectionById(sectionId)
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

func (r *Repository) UpdateSection(section ds.Section) error {
	err := r.db.Model(&ds.Section{}).Where("id = ?", section.Id.String()).Updates(section).Error
	if err != nil {
		return errors.Wrap(err, "Can't update section in repo")
	}

	return nil
}

func (r *Repository) DeleteSection(section ds.Section) error {
	err := r.db.Delete(&section).Error
	if err != nil {
		return errors.Wrap(err, "Can't delete section in repo")
	}

	return nil
}

func (r *Repository) CreateNotification(notification ds.Notification) error {
	err := r.db.Create(&notification).Error
	if err != nil {
		return errors.Wrap(err, "Can't create notification in repo")
	}

	return nil
}

func (r *Repository) GetAllNotifications(sectionId string) ([]ds.Notification, error) {
	notifications := []ds.Notification{}
	err := r.db.Where("section_id = ?", sectionId).Find(&notifications).Error
	if err != nil {
		return nil, errors.Wrap(err, "Can't find such notification")
	}

	return notifications, nil
}

func (r *Repository) GetNotification(notificationId string) ([]ds.Notification, error) {
	notifications := []ds.Notification{}
	err := r.db.Where("id = ?", notificationId).Find(&notifications).Error
	if err != nil {
		return nil, errors.Wrap(err, "Can't find such notification")
	}

	return notifications, nil
}
