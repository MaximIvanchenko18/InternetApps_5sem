package repository

import (
	"InternetApps_5sem/internal/app/ds"
	"errors"
	"strings"
	"time"

	"gorm.io/gorm"
)

func (r *Repository) GetAllFlights(customerId *string, formDateStart *time.Time, formDateEnd *time.Time, status string) ([]ds.Flight, error) {
	var flights []ds.Flight
	query := r.db.Preload("Customer").Preload("Moderator").
		Where("LOWER(status) LIKE ?", "%"+strings.ToLower(status)+"%").
		Where("status != ? AND status != ?", ds.StatusDeleted, ds.StatusDraft)

	if customerId != nil {
		query = query.Where("customer_id = ?", *customerId)
	}

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

func (r *Repository) GetDraftFlight(customerId string) (*ds.Flight, error) {
	flight := &ds.Flight{}
	err := r.db.First(flight, ds.Flight{Status: ds.StatusDraft, CustomerId: customerId}).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}

	return flight, nil
}

func (r *Repository) CreateDraftFlight(customerId string) (*ds.Flight, error) {
	flight := &ds.Flight{Status: ds.StatusDraft, CreationDate: time.Now(), CustomerId: customerId}

	err := r.db.Create(flight).Error
	if err != nil {
		return nil, err
	}

	return flight, nil
}

func (r *Repository) GetFlightById(flightId string, customerId *string) (*ds.Flight, error) {
	flight := &ds.Flight{}
	query := r.db.Preload("Customer").Preload("Moderator").
		Where("status != ?", ds.StatusDeleted)

	if customerId != nil {
		query = query.Where("customer_id = ?", customerId)
	}

	err := query.First(flight, ds.Flight{UUID: flightId}).Error

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

func (r *Repository) CountCargo(flightId string) (int64, error) {
	var count int64
	err := r.db.Model(&ds.FlightCargo{}).
		Where("flight_id = ?", flightId).
		Count(&count).Error

	if err != nil {
		return 0, err
	}

	return count, nil
}
