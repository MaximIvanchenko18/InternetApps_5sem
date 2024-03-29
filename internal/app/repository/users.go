package repository

import (
	"InternetApps_5sem/internal/app/ds"
	"errors"

	"gorm.io/gorm"
)

func (r *Repository) AddUser(user *ds.User) error {
	return r.db.Create(user).Error
}

func (r *Repository) GetUserByLogin(login string) (*ds.User, error) {
	user := &ds.User{}

	err := r.db.Where("login = ?", login).First(user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}

	return user, nil
}

func (r *Repository) GetUserById(uuid string) (*ds.User, error) {
	user := &ds.User{}

	err := r.db.Where("uuid = ?", uuid).First(user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}

	return user, nil
}
