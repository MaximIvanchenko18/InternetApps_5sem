package repository

import (
	"InternetApps_5sem/internal/app/ds"
	"errors"
	"strings"
	"time"

	"gorm.io/gorm"
)

func (r *Repository) GetAllFlights(formDateStart *time.Time, formDateEnd *time.Time, status string) ([]ds.Flight, error) {
	var flights []ds.Flight
	query := r.db.Preload("Client").Preload("Moderator").
		Where("LOWER(status) LIKE ?", "%"+strings.ToLower(status)+"%").
		Where("status != ?", ds.DELETED)

	if formDateStart != nil && formDateEnd != nil {
		query = query.Where("formation_date BETWEEN ? AND ?", *formDateStart, *formDateEnd)
	} else if formDateStart != nil {
		query = query.Where("formation_date >= ?", *formDateStart)
	} else if formDateEnd != nil {
		query = query.Where("formation_date <= ?", *formDateEnd)
	}

	err := query.Find(&flights).Error
	if err != nil {
		return nil, err
	}

	return flights, nil
}

func (r *Repository) GetDraftFlight(clientId string) (*ds.Flight, error) {
	flight := &ds.Flight{}
	err := r.db.First(flight, ds.Flight{Status: ds.DRAFT, ClientId: clientId}).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}

	return flight, nil
}

func (r *Repository) CreateDraftFlight(clientId string) (*ds.Flight, error) {
	flight := &ds.Flight{Status: ds.DRAFT, CreationDate: time.Now(), ClientId: clientId}

	err := r.db.Create(flight).Error
	if err != nil {
		return nil, err
	}

	return flight, nil
}

func (r *Repository) GetFlightById(flightId string, clientId string) (*ds.Flight, error) {
	flight := &ds.Flight{}
	err := r.db.Preload("Client").Preload("Moderator").
		Where("status != ?", ds.DELETED).
		First(flight, ds.Flight{UUID: flightId, ClientId: clientId}).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}

	return flight, nil
}

func (r *Repository) GetFlightCargos(flightId string) ([]ds.Cargo, error) {
	var cargos []ds.Cargo

	err := r.db.Table("flight_cargos").
		Select("cargos.*").
		Joins("JOIN cargos ON flight_cargos.cargo_id = cargos.uuid").
		Where(ds.FlightCargo{FlightId: flightId}).
		Scan(&cargos).Error

	if err != nil {
		return nil, err
	}

	return cargos, nil
}

func (r *Repository) SaveFlight(flight *ds.Flight) error {
	err := r.db.Save(flight).Error

	if err != nil {
		return err
	}

	return nil
}

func (r *Repository) DeleteFromFlight(flightId string, cargoId string) error {
	err := r.db.Delete(&ds.FlightCargo{FlightId: flightId, CargoId: cargoId}).Error

	if err != nil {
		return err
	}

	return nil
}
