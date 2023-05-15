package repository

import (
	"mkn-backend/internal/pkg/ds"

	"github.com/google/uuid"
	"github.com/pkg/errors"
)

func (r *Repository) UpdateProject(userId, projectId string, project ds.Project) error {
	err := r.db.Model(&ds.Project{}).Where("owner_id = ? AND id = ?", userId, projectId).Updates(ds.Project{LastEdited: project.LastEdited, Description: project.Description, Title: project.Title, Color: project.Color}).Error
	if err != nil {
		return errors.Wrap(err, "Can't update project")
	}

	return nil
}

// !!!!!!!!!!!!!!!!!!!!!!
func (r *Repository) DeleteProject(projectId string) error {
	project := &ds.Project{}
	favProjects := &ds.FavoriteProject{}
	collabs := &ds.Collaboration{}
	sections := &ds.Section{}

	err := r.db.Where("project_id = ?", projectId).Delete(&collabs).Error
	if err != nil {
		return errors.Wrap(err, "Can't delete this project")
	}

	err = r.db.Where("project_id = ?", projectId).Delete(&sections).Error
	if err != nil {
		return errors.Wrap(err, "Can't delete this project")
	}

	err = r.db.Where("project_id = ?", projectId).Delete(&favProjects).Error
	if err != nil {
		return errors.Wrap(err, "Can't delete this project")
	}

	err = r.db.Where("id = ?", projectId).Delete(&project).Error
	if err != nil {
		return errors.Wrap(err, "Can't delete this project")
	}

	return nil
}

func (r *Repository) AddCollaborator(ownerId, projectIdStr, collabIdStr string) error {
	projectId, _ := uuid.Parse(projectIdStr)
	collaboratorId, _ := uuid.Parse(collabIdStr)

	collab := ds.Collaboration{
		Id:        uuid.New(),
		ProjectId: projectId,
		UserId:    collaboratorId,
	}

	err := r.db.Create(&collab).Error
	if err != nil {
		return errors.Wrap(err, "Can't create such data row")
	}

	return nil
}

func (r *Repository) GetCollaborator(userId, projectId string) (*ds.User, error) {
	collaborator := ds.Collaboration{}
	err := r.db.Where("project_id = ? AND user_id = ?", projectId, userId).First(&collaborator).Error
	if err != nil {
		return nil, errors.Wrap(err, "Can't find such collaborator")
	}

	user, err := r.GetUserById(collaborator.UserId.String())
	if err != nil {
		return nil, errors.Wrap(err, "Can't get such user")
	}

	return user, nil
}

func (r *Repository) GetAllCollaborators(projectId string) ([]ds.User, error) {
	collaborators := []ds.Collaboration{}
	err := r.db.Where("project_id = ?", projectId).Find(&collaborators).Error
	if err != nil {
		return nil, errors.Wrap(err, "Can't find such collaborator")
	}

	users := make([]ds.User, 0, 100)

	for i := range collaborators {
		user, err := r.GetUserById(collaborators[i].UserId.String())
		if err != nil {
			return nil, errors.Wrap(err, "Can't get such user")
		}
		users = append(users, *user)
	}

	return users, nil
}

func (r *Repository) IsCollaborator(userId, projectId string) bool {
	if 0 == r.db.Where("user_id = ? AND project_id = ?", userId, projectId).First(&ds.Collaboration{}).RowsAffected {
		return false
	}

	return true
}

func (r *Repository) DeleteCollaborator(ownerId, projectIdStr, collabIdStr string) error {
	err := r.db.Where("user_id = ? AND project_id = ?", collabIdStr, projectIdStr).Delete(&ds.Collaboration{}).Error
	if err != nil {
		return errors.Wrap(err, "Can't create such data row")
	}

	return nil
}

func (r *Repository) CreateSection(section ds.Section) error {
	err := r.db.Create(&section).Error
	if err != nil {
		return errors.Wrap(err, "Can't create section in repository")
	}

	return nil
}

func (r *Repository) GetAllSections(projectId string) ([]ds.Section, error) {
	sections := []ds.Section{}
	err := r.db.Where("project_id = ?", projectId).Find(&sections).Error
	if err != nil {
		return nil, errors.Wrap(err, "Can't find such section")
	}

	return sections, nil
}
