package repository

import (
	"mkn-backend/internal/pkg/ds"

	"github.com/pkg/errors"
)

func (r *Repository) UpdateProject(userId, projectId string, project ds.Project) error {
	err := r.db.Model(&ds.Project{}).Where("owner_id = ? AND id = ?", userId, projectId).Updates(ds.Project{LastEdited: project.LastEdited, Description: project.Description, Title: project.Title, Color: project.Color}).Error
	if err != nil {
		return errors.Wrap(err, "Can't update project")
	}

	return nil
}

func (r *Repository) DeleteProject(projectId string) error {
	project := &ds.Project{}
	favProjects := &ds.FavoriteProject{}
	err := r.db.Where("project_id = ?", projectId).Delete(&favProjects).Error
	if err != nil {
		return errors.Wrap(err, "Can't delete this project")
	}

	err = r.db.Where("id = ?", projectId).Delete(&project).Error
	if err != nil {
		return errors.Wrap(err, "Can't delete this project")
	}

	return nil
}
