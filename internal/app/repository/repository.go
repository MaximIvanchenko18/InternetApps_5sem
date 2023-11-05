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

	err := r.db.Where("is_deleted = ?", false).Find(&cargo).Error
	if err != nil {
		return nil, err
	}

	return cargo, nil
}

func (r *Repository) GetCargoByID(id int) (*ds.Cargo, error) {
	cargo := &ds.Cargo{}
	err := r.db.Where("cargo_id = ? AND is_deleted = ?", id, false).First(&cargo).Error

	if err != nil {
		return nil, err
	}

	return cargo, nil
}

func (r *Repository) GetCargoByEnName(en_name string) (*ds.Cargo, error) {
	cargo := &ds.Cargo{}
	err := r.db.Where("english_name = ? AND is_deleted = ?", en_name, false).First(&cargo).Error

	if err != nil {
		return nil, err
	}

	return cargo, nil
}

func (r *Repository) GetFilteredCargo(name string, lowprice uint, highprice uint) ([]ds.Cargo, error) {
	var cargo []ds.Cargo

	err := r.db.Where("LOWER(cargos.name) LIKE ? AND price >= ? AND price <= ? AND is_deleted = ?",
		"%"+strings.ToLower(name)+"%", lowprice, highprice, false).Find(&cargo).Error

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

func (r *Repository) GetLowestPrice() (uint, error) {
	cargos, err := r.GetAllCargo()

	if err != nil {
		return 0, err
	}

	min_price := cargos[0].Price
	for _, cargo := range cargos {
		if cargo.Price < min_price {
			min_price = cargo.Price
		}
	}

	return min_price, nil
}

func (r *Repository) GetHighestPrice() (uint, error) {
	cargos, err := r.GetAllCargo()

	if err != nil {
		return 0, err
	}

	max_price := cargos[0].Price
	for _, cargo := range cargos {
		if cargo.Price > max_price {
			max_price = cargo.Price
		}
	}

	return max_price, nil
}
