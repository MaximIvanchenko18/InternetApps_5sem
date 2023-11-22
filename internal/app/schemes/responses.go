package schemes

import (
	"InternetApps_5sem/internal/app/ds"
)

type AllCargosResponse struct {
	Cargos []ds.Cargo `json:"cargos"`
}

type FlightShort struct {
	UUID       string `json:"uuid"`
	CargoCount int    `json:"cargo_count"`
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
	Client         string  `json:"customer"`
	Moderator      *string `json:"moderator"`
	RocketType     string  `json:"rocket_type"`
}

func ConvertFlight(flight *ds.Flight) FlightOutput {
	output := FlightOutput{
		UUID:         flight.UUID,
		Status:       flight.Status,
		CreationDate: flight.CreationDate.Format("2006-01-02 15:04:05"),
		RocketType:   flight.RocketType,
		Client:       flight.Client.Name,
	}

	if flight.FormationDate != nil {
		formationDate := flight.FormationDate.Format("2006-01-02 15:04:05")
		output.FormationDate = &formationDate
	}

	if flight.CompletionDate != nil {
		completionDate := flight.CompletionDate.Format("2006-01-02 15:04:05")
		output.CompletionDate = &completionDate
	}

	if flight.Moderator != nil {
		output.Moderator = &flight.Moderator.Name
	}

	return output
}
