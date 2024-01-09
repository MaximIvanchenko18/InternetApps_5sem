package app

import (
	"fmt"
	"net/http"
	"time"

	"InternetApps_5sem/internal/app/ds"
	"InternetApps_5sem/internal/app/role"
	"InternetApps_5sem/internal/app/schemes"

	"github.com/gin-gonic/gin"
)

// @Summary		Получить все полеты
// @Tags		Полеты
// @Description	Возвращает все полеты с фильтрацией по статусу и дате формирования
// @Produce		json
// @Param		status query string false "статус полета"
// @Param		form_date_start query string false "начальная дата формирования"
// @Param		form_date_end query string false "конечная дата формирвания"
// @Success		200 {object} schemes.AllFlightsResponse
// @Router		/api/flights [get]
func (app *Application) GetAllFlights(c *gin.Context) {
	var request schemes.GetAllFlightsRequest
	err := c.ShouldBindQuery(&request)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	userId := getUserId(c)
	userRole := getUserRole(c)

	var flights []ds.Flight
	if userRole == role.Customer {
		flights, err = app.repo.GetAllFlights(&userId, request.FormDateStart, request.FormDateEnd, request.Status)
	} else {
		flights, err = app.repo.GetAllFlights(nil, request.FormDateStart, request.FormDateEnd, request.Status)
	}
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	outputFlights := make([]schemes.FlightOutput, len(flights))
	for i, flight := range flights {
		outputFlights[i] = schemes.ConvertFlight(&flight)
	}

	c.JSON(http.StatusOK, schemes.AllFlightsResponse{Flights: outputFlights})
}

// @Summary		Получить один полет
// @Tags		Полеты
// @Description	Возвращает подробную информацию о полете
// @Produce		json
// @Param		flight_id path string true "id полета"
// @Success		200 {object} schemes.FlightResponse
// @Router		/api/flights/{flight_id} [get]
func (app *Application) GetFlight(c *gin.Context) {
	var request schemes.FlightRequest
	err := c.ShouldBindUri(&request)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	userId := getUserId(c)
	userRole := getUserRole(c)

	var flight *ds.Flight
	if userRole == role.Moderator {
		flight, err = app.repo.GetFlightById(request.FlightId, nil)
	} else {
		flight, err = app.repo.GetFlightById(request.FlightId, &userId)
	}
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	if flight == nil {
		c.AbortWithError(http.StatusNotFound, fmt.Errorf("полет не найден"))
		return
	}

	cargos, quantities, err := app.repo.GetFlightCargos(request.FlightId)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	cargosQuantity := make([]schemes.CargoQuantity, len(cargos))
	for i := 0; i < len(cargos); i++ {
		cargosQuantity[i].Cargo = cargos[i]
		cargosQuantity[i].Quantity = quantities[i]
	}

	c.JSON(http.StatusOK, schemes.FlightResponse{Flight: schemes.ConvertFlight(flight), Cargos: cargosQuantity})
}

type SwaggerUpdateFlightRequest struct {
	RocketType string `json:"rocket_type"`
}

// @Summary		Указать тип ракеты
// @Tags		Полеты
// @Description	Позволяет изменить тип ракеты
// @Access		json
// @Produce		json
// @Param		rocket_type body SwaggerUpdateFlightRequest true "Тип ракеты"
// @Success		200
// @Router		/api/flights [put]
func (app *Application) UpdateFlight(c *gin.Context) {
	var request schemes.UpdateFlightRequest

	err := c.ShouldBind(&request)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	userId := getUserId(c)
	flight, err := app.repo.GetDraftFlight(userId)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	if flight == nil {
		c.AbortWithError(http.StatusNotFound, fmt.Errorf("полет не найден"))
		return
	}

	flight.RocketType = &request.RocketType
	err = app.repo.SaveFlight(flight)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.Status(http.StatusOK)
}

// @Summary		Удалить полет-черновик
// @Tags		Полеты
// @Description	Удаляет полет-черновик пользователя
// @Success		200
// @Router		/api/flights [delete]
func (app *Application) DeleteFlight(c *gin.Context) {
	userId := getUserId(c)
	flight, err := app.repo.GetDraftFlight(userId)

	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	if flight == nil {
		c.AbortWithError(http.StatusNotFound, fmt.Errorf("полет не найден"))
		return
	}

	flight.Status = ds.StatusDeleted

	err = app.repo.SaveFlight(flight)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.Status(http.StatusOK)
}

// @Summary		Удалить груз из полета-черновика
// @Tags		Полеты
// @Description	Удалить груз из полета-черновика пользователя
// @Produce		json
// @Param		cargo_id path string true "id груза"
// @Success		200
// @Router		/api/flights/delete_cargo/{cargo_id} [delete]
func (app *Application) DeleteFromFlight(c *gin.Context) {
	var request schemes.DeleteFromFlightRequest
	err := c.ShouldBindUri(&request)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	userId := getUserId(c)
	flight, err := app.repo.GetDraftFlight(userId)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	if flight == nil {
		c.AbortWithError(http.StatusNotFound, fmt.Errorf("полет не найден"))
		return
	}

	err = app.repo.DeleteFromFlight(flight.UUID, request.CargoId)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.Status(http.StatusOK)
}

// @Summary		Изменить количество груза в полете
// @Tags		Полеты
// @Description	Изменить количество груза в полете
// @Produce		json
// @Param		cargo_id path string true "id груза"
// @Param		quantity query uint true "количество груза"
// @Success		200
// @Router		/api/flights/change_cargo/{cargo_id} [put]
func (app *Application) UpdateFlightCargoQuantity(c *gin.Context) {
	var request schemes.UpdateFlightCargoQuantityRequest
	err := c.ShouldBindUri(&request.URI)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	err = c.ShouldBind(&request)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	if request.Quantity <= 0 {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	userId := getUserId(c)
	flight, err := app.repo.GetDraftFlight(userId)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	if flight == nil {
		c.AbortWithError(http.StatusNotFound, fmt.Errorf("полет не найден"))
		return
	}
	if flight.Status != ds.StatusDraft {
		c.AbortWithError(http.StatusMethodNotAllowed, fmt.Errorf("нельзя редактировать полет статуса: %s", flight.Status))
		return
	}

	cargos, _, err := app.repo.GetFlightCargos(flight.UUID)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	isCargoInFlight := false
	for _, cargo := range cargos {
		if cargo.UUID == request.URI.CargoId {
			isCargoInFlight = true
			break
		}
	}
	if !isCargoInFlight {
		c.AbortWithError(http.StatusNotFound, fmt.Errorf("груз в полете не найден"))
		return
	}

	flightcargo := &ds.FlightCargo{FlightId: flight.UUID, CargoId: request.URI.CargoId, Quantity: request.Quantity}
	err = app.repo.SaveFlightCargo(flightcargo)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.Status(http.StatusOK)
}

// @Summary		Сформировать полет
// @Tags		Полеты
// @Description	Сформировать полет пользователем
// @Success		200
// @Router		/api/flights/user_confirm [put]
func (app *Application) UserConfirm(c *gin.Context) {
	userId := getUserId(c)

	flight, err := app.repo.GetDraftFlight(userId)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	if flight == nil {
		c.AbortWithError(http.StatusNotFound, fmt.Errorf("полет не найден"))
		return
	}

	err = shipmentRequest(flight.UUID)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, fmt.Errorf(`shipment service is unavailable: {%s}`, err))
		return
	}

	customer, err := app.repo.GetUserById(userId)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	flight.Customer = *customer

	shipmentStatus := ds.ShipmentStarted
	flight.ShipmentStatus = &shipmentStatus
	flight.Status = ds.StatusFormed
	now := time.Now()
	flight.FormationDate = &now

	err = app.repo.SaveFlight(flight)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.Status(http.StatusOK)
}

// @Summary		Завершить полет
// @Tags		Полеты
// @Description	Подтвердить или отклонить полет модератором
// @Param		flight_id path string true "id полета"
// @Param		confirm body boolean true "подтвердить"
// @Success		200
// @Router		/api/flights/{flight_id}/moderator_confirm [put]
func (app *Application) ModeratorConfirm(c *gin.Context) {
	var request schemes.ModeratorConfirmRequest
	err := c.ShouldBindUri(&request.URI)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	err = c.ShouldBind(&request)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	userId := getUserId(c)
	flight, err := app.repo.GetFlightById(request.URI.FlightId, nil)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	if flight == nil {
		c.AbortWithError(http.StatusNotFound, fmt.Errorf("полет не найден"))
		return
	}
	if flight.Status != ds.StatusFormed {
		c.AbortWithError(http.StatusMethodNotAllowed, fmt.Errorf("нельзя завершить статус %s", flight.Status))
		return
	}

	if *request.Confirm {
		flight.Status = ds.StatusCompleted
	} else {
		flight.Status = ds.StatusRejected
	}
	now := time.Now()
	flight.CompletionDate = &now

	moderator, err := app.repo.GetUserById(userId)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	flight.ModeratorId = &userId
	flight.Moderator = moderator

	err = app.repo.SaveFlight(flight)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.Status(http.StatusOK)
}

func (app *Application) Shipment(c *gin.Context) {
	var request schemes.ShipmentReq
	err := c.ShouldBindUri(&request.URI)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	err = c.ShouldBind(&request)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	if request.Token != app.config.Token {
		c.AbortWithStatus(http.StatusForbidden)
		return
	}

	flight, err := app.repo.GetFlightById(request.URI.FlightId, nil)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	if flight == nil {
		c.AbortWithError(http.StatusNotFound, fmt.Errorf("груз не найден"))
		return
	}

	var shipmentStatus string
	if *request.ShipmentStatus {
		shipmentStatus = ds.ShipmentCompleted
	} else {
		shipmentStatus = ds.ShipmentFailed
	}
	flight.ShipmentStatus = &shipmentStatus

	err = app.repo.SaveFlight(flight)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.Status(http.StatusOK)
}
