package repository

import (
	"strings"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"InternetApps_5sem/internal/app/ds"
)

type Repository struct {
	db *gorm.DB
}

func New(dsn string) (*Repository, error) {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return &Repository{
		db: db,
	}, nil
}

func (r *Repository) GetAllCargo() ([]ds.Cargo, error) {
	var cargo []ds.Cargo

	err := r.db.Find(&cargo).Error
	if err != nil {
		return nil, err
	}

	return cargo, nil
}

func (r *Repository) GetCargoByID(id int) (*ds.Cargo, error) {
	cargo := &ds.Cargo{}
	err := r.db.Where("cargo_id = ?", id).First(&cargo).Error

	if err != nil {
		return nil, err
	}

	return cargo, nil
}

func (r *Repository) GetCargoByEnName(en_name string) (*ds.Cargo, error) {
	cargo := &ds.Cargo{}
	err := r.db.Where("english_name = ", en_name).First(&cargo).Error

	if err != nil {
		return nil, err
	}

	return cargo, nil
}

func (r *Repository) GetFilteredCargo(name string, lowprice int, highprice int) ([]ds.Cargo, error) {
	var cargo []ds.Cargo
	var err error = nil

	if lowprice < 0 {
		lowprice = 0
	}
	if highprice < 0 {
		highprice = 0
	}

	if highprice > 0 {
		err = r.db.Where("LOWER(cargos.name) LIKE ? AND price >= ? AND price <= ? AND is_deleted = ?",
			"%"+strings.ToLower(name)+"%", lowprice, highprice, false).Find(&cargo).Error
	} else {
		err = r.db.Where("LOWER(cargos.name) LIKE ? AND price >= ? AND is_deleted = ?",
			"%"+strings.ToLower(name)+"%", lowprice, false).Find(&cargo).Error
	}

	if err != nil {
		return nil, err
	}

	return cargo, nil
}

func (r *Repository) DeleteCargoById(id int) error {
	err := r.db.Exec("UPDATE cargos SET is_deleted = ? WHERE cargo_id = ?", true, id).Error
	if err != nil {
		return err
	}

	return nil
}
