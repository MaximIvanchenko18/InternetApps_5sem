package repository

import (
	"InternetApps_5sem/internal/app/ds"
)

func (r *Repository) AddUser(user *ds.User) error {
	return r.db.Create(user).Error
}

func (r *Repository) GetUserByLogin(login string) (*ds.User, error) {
	user := &ds.User{}
	err := r.db.Where("login = ?", login).First(user).Error

	if err != nil {
		return nil, err
	}

	return user, nil
}
