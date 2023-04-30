package repository

import "mkn-backend/internal/pkg/ds"

func (repo *Repository) SignUp(user *ds.User) (*ds.User, error) {
	_ = repo.db.Create(user)
	result := &ds.User{}
	err := repo.db.First(&result, "username = ?", user.Username).Error
	if err != nil {
		return nil, err
	}

	return result, nil
}
