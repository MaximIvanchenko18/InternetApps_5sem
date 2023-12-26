package schemes

import (
	"InternetApps_5sem/internal/app/ds"
	"InternetApps_5sem/internal/app/role"
	"time"
)

type AllCargosResponse struct {
	Cargos []ds.Cargo `json:"cargos"`
}

type FlightShort struct {
	UUID       string `json:"uuid"`
	CargoCount int64  `json:"cargo_count"`
}

type GetAllCargosResponse struct {
	DraftFlight *FlightShort `json:"draft_flight"`
	Cargos      []ds.Cargo   `json:"cargos"`
}

type AllFlightsResponse struct {
	Flights []FlightOutput `json:"flights"`
}

type FlightResponse struct {
	Flight FlightOutput `json:"flight"`
	Cargos []ds.Cargo   `json:"cargos"`
}

type UpdateFlightResponse struct {
	Flight FlightOutput `json:"flight"`
}

type FlightOutput struct {
	UUID           string  `json:"uuid"`
	Status         string  `json:"status"`
	CreationDate   string  `json:"creation_date"`
	FormationDate  *string `json:"formation_date"`
	CompletionDate *string `json:"completion_date"`
	Customer       string  `json:"customer"`
	Moderator      *string `json:"moderator"`
	RocketType     *string `json:"rocket_type"`
	ShipmentStatus *string `json:"shipment_status"`
}

func ConvertFlight(flight *ds.Flight) FlightOutput {
	output := FlightOutput{
		UUID:           flight.UUID,
		Status:         flight.Status,
		CreationDate:   flight.CreationDate.Format("2006-01-02T15:04:05Z07:00"),
		RocketType:     flight.RocketType,
		Customer:       flight.Customer.Login,
		ShipmentStatus: flight.ShipmentStatus,
	}

	if flight.FormationDate != nil {
		formationDate := flight.FormationDate.Format("2006-01-02T15:04:05Z07:00")
		output.FormationDate = &formationDate
	}

	if flight.CompletionDate != nil {
		completionDate := flight.CompletionDate.Format("2006-01-02T15:04:05Z07:00")
		output.CompletionDate = &completionDate
	}

	if flight.Moderator != nil {
		output.Moderator = &flight.Moderator.Login
	}

	return output
}

type AddToFlightResp struct {
	CargoCount int64 `json:"cargo_count"`
}

type AuthResp struct {
	ExpiresIn   time.Duration `json:"expires_in"`
	AccessToken string        `json:"access_token"`
	Role        role.Role     `json:"role"`
	Login       string        `json:"login"`
	TokenType   string        `json:"token_type"`
}

type SwaggerLoginResp struct {
	ExpiresIn   int64  `json:"expires_in"`
	AccessToken string `json:"access_token"`
	Role        int    `json:"role"`
	Login       string `json:"login"`
	TokenType   string `json:"token_type"`
}
