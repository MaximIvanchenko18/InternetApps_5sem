package repository

import (
	"InternetApps_5sem/internal/app/ds"
	"errors"
	"strings"

	"gorm.io/gorm"
)

func (r *Repository) GetAllCargo() ([]ds.Cargo, error) {
	var cargo []ds.Cargo

	err := r.db.Where("is_deleted = ?", false).Find(&cargo).Error
	if err != nil {
		return nil, err
	}

	return cargo, nil
}

func (r *Repository) GetCargoByID(id string) (*ds.Cargo, error) {
	cargo := &ds.Cargo{UUID: id}
	err := r.db.First(cargo, "is_deleted = ?", false).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}

	return cargo, nil
}

func (r *Repository) GetCargoByEnName(en_name string) (*ds.Cargo, error) {
	cargo := &ds.Cargo{}
	err := r.db.Where("english_name = ? AND is_deleted = ?", en_name, false).First(cargo).Error

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

func (r *Repository) AddCargo(cargo *ds.Cargo) error {
	err := r.db.Create(&cargo).Error

	if err != nil {
		return err
	}

	return nil
}

func (r *Repository) SaveCargo(cargo *ds.Cargo) error {
	err := r.db.Save(cargo).Error

	if err != nil {
		return err
	}

	return nil
}

func (r *Repository) AddToFlight(flightId string, cargoId string, quantity uint) error {
	CargoToFlight := ds.FlightCargo{FlightId: flightId, CargoId: cargoId, Quantity: quantity}
	err := r.db.Create(&CargoToFlight).Error

	if err != nil {
		return err
	}

	return nil
}

func (r *Repository) SaveFlightCargo(flightcargo *ds.FlightCargo) error {
	err := r.db.Save(flightcargo).Error

	if err != nil {
		return err
	}

	return nil
}
