package repository

import (
	"mkn-backend/internal/pkg/ds"

	"github.com/google/uuid"
	"github.com/pkg/errors"
)

func (r *Repository) CreateProject(project *ds.Project) (*ds.Project, error) {
	_ = r.db.Create(project)

	res, err := r.GetProject(project.Id.String())
	if err != nil {
		return nil, errors.Wrap(err, "Can't find such project")
	}

	return res, nil
}

func (r *Repository) GetProject(id string) (*ds.Project, error) {
	project := &ds.Project{}
	err := r.db.Where("id = ?", id).First(&project).Error
	if err != nil {
		return nil, errors.Wrap(err, "No such project at repository")
	}
	return project, nil
}

func (r *Repository) GetAllProjects() ([]ds.Project, error) {
	projects := []ds.Project{}

	_ = r.db.Find(&projects)
	return projects, nil
}

func (r *Repository) GetAllOwnedProjects(user_id string) ([]ds.Project, error) {
	projects := []ds.Project{}

	_ = r.db.Where("owner_id = ?", user_id).Find(&projects)
	return projects, nil
}

func (r *Repository) AddFavorite(userIdStr, projectIdStr string) error {
	userId, err := uuid.Parse(userIdStr)
	if err != nil {
		return errors.New("Wrong user id")
	}

	projectId, err := uuid.Parse(projectIdStr)
	if err != nil {
		return errors.New("Wrong project id")
	}

	favoriteProject := ds.FavoriteProject{
		Id:        uuid.New(),
		UserId:    userId,
		ProjectId: projectId,
	}

	_ = r.db.Create(favoriteProject)
	return nil
}

func (r *Repository) GetFavoriteProjects(userId string) ([]ds.FavoriteProject, error) {
	favoriteProjects := []ds.FavoriteProject{}
	err := r.db.Where("user_id = ?", userId).Find(&favoriteProjects).Error
	if err != nil {
		return []ds.FavoriteProject{}, errors.New("Can't get all favorites projects")
	}

	return favoriteProjects, nil
}

func (r *Repository) GetFavoriteProject(userId, projectId string) (ds.FavoriteProject, error) {
	favoriteProject := ds.FavoriteProject{}
	err := r.db.Where("user_id = ? AND project_id = ?", userId, projectId).First(&favoriteProject).Error
	if err != nil {
		return ds.FavoriteProject{}, errors.Wrap(err, "No such user-project pair")
	}

	return favoriteProject, nil
}

func (r *Repository) DeleteFavorite(userId, projectId string) error {
	var favoriteProjects []ds.FavoriteProject

	if _, err := r.GetFavoriteProject(userId, projectId); err != nil {
		return err
	}

	err := r.db.Where("user_id = ? AND project_id = ?", userId, projectId).Delete(&favoriteProjects).Error
	if err != nil {
		return errors.Wrap(err, "Can't delete favorite project in repo")
	}

	return nil
}

func (r *Repository) GetProjectById(projectId string) (ds.Project, error) {
	project := ds.Project{}

	err := r.db.Where("id = ?", projectId).First(&project).Error
	if err != nil {
		return ds.Project{}, errors.Wrap(err, "Can't get project in repo")
	}

	return project, nil
}

func (r *Repository) ChangeEmail(userId, newEmail string) error {
	err := r.db.Model(&ds.User{}).Where("id = ?", userId).Update("email", newEmail).Error
	if err != nil {
		return errors.Wrap(err, "Can't change email in repo")
	}

	return nil
}

func (r *Repository) LastThreeProjects(userId string) ([]ds.Project, error) {
	projects := []ds.Project{}
	err := r.db.Order("last_edited desc").Where("owner_id = ?", userId).Limit(3).Find(&projects).Error
	if err != nil {
		return []ds.Project{}, errors.Wrap(err, "Can't get projects from repo")
	}

	return projects, nil
}
