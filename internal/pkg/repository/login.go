package repository

import (
	"mkn-backend/internal/pkg/ds"

	"github.com/pkg/errors"
)

func (repo *Repository) GetUserById(id string) (*ds.User, error) {
	user := &ds.User{}
	err := repo.db.First(&user, "id = ?", id).Error
	if err != nil {
		return nil, errors.Wrap(err, "No such user at repository")
	}
	return user, nil
}

func (repo *Repository) GetUser(username, password string) (*ds.User, error) {
	user := &ds.User{}
	err := repo.db.First(&user, "username = ? AND password = ?", username, password).Error
	if err != nil {
		return nil, errors.Wrap(err, "No such user at repository")
	}
	return user, nil
}

func (repo *Repository) GetUserByName(username string) (*ds.User, error) {
	user := &ds.User{}
	err := repo.db.First(&user, "username = ?", username).Error
	if err != nil {
		return nil, errors.Wrap(err, "No such user at repository")
	}
	return user, nil
}

func (repo *Repository) SignUp(user *ds.User) (*ds.User, error) {
	_, err := repo.GetUserByName(user.Username)
	if err == nil {
		return nil, errors.Wrap(err, "This nickname is already taken")
	}
	_ = repo.db.Create(user)

	res, err := repo.GetUser(user.Username, user.Password)
	if err != nil {
		return nil, errors.Wrap(err, "Can't return user in response")
	}

	return res, nil
}
