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
