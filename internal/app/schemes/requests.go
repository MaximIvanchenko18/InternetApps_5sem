package schemes

import (
	"InternetApps_5sem/internal/app/ds"

	"mime/multipart"
	"time"
)

type CargoRequest struct {
	CargoId string `uri:"cargo_id" binding:"required,uuid"`
}

type GetAllCargosRequest struct {
	Name      string `form:"name"`
	LowPrice  string `form:"low_price"`
	HighPrice string `form:"high_price"`
}

type AddCargoRequest struct {
	ds.Cargo
	Image *multipart.FileHeader `form:"image" json:"image" binding:"required"`
}

type ChangeCargoRequest struct {
	CargoId     string                `uri:"cargo_id" binding:"required,uuid"`
	Name        *string               `form:"name" json:"name" binding:"omitempty,max=30"`
	EnglishName *string               `form:"en_name" json:"en_name" binding:"omitempty,max=30"`
	Image       *multipart.FileHeader `form:"image" json:"image"`
	Category    *string               `form:"category" json:"category" binding:"omitempty,max=50"`
	Price       *uint                 `form:"price" json:"price"`
	Weight      *float32              `form:"weight" json:"weight"`
	Capacity    *float32              `form:"capacity" json:"capacity"`
	Description *string               `form:"description" json:"description" binding:"omitempty,max=500"`
}

type AddToFlightRequest struct {
	CargoId  string `uri:"cargo_id" binding:"required,uuid"`
	Quantity uint   `form:"quantity"`
}

type GetAllFlightsRequest struct {
	FormDateStart *time.Time `form:"form_date_start" json:"form_date_start" time_format:"2003-10-15 18:10:18"`
	FormDateEnd   *time.Time `form:"form_date_end" json:"form_date_end" time_format:"2003-10-15 18:10:18"`
	Status        string     `form:"status"`
}

type FlightRequest struct {
	FlightId string `uri:"flight_id" binding:"required,uuid"`
}

type UpdateFlightRequest struct {
	URI struct {
		FlightId string `uri:"flight_id" binding:"required,uuid"`
	}
	RocketType string `form:"rocket_type" json:"rocket_type" binding:"required,max=50"`
}

type DeleteFromFlightRequest struct {
	FlightId string `uri:"flight_id" binding:"required,uuid"`
	CargoId  string `uri:"cargo_id" binding:"required,uuid"`
}

type UserConfirmRequest struct {
	FlightId string `uri:"flight_id" binding:"required,uuid"`
}

type ModeratorConfirmRequest struct {
	URI struct {
		FlightId string `uri:"flight_id" binding:"required,uuid"`
	}
	Status string `form:"status" json:"status" binding:"required"`
}
