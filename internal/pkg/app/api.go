package app

import (
	"fmt"
	"net/http"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"InternetApps_5sem/internal/app/ds"
	"InternetApps_5sem/internal/app/schemes"

	"mime/multipart"

	"github.com/gin-gonic/gin"
	"github.com/minio/minio-go/v7"
)

func (app *Application) uploadImage(c *gin.Context, image *multipart.FileHeader, UUID string) (*string, error) {
	src, err := image.Open()
	if err != nil {
		return nil, err
	}
	defer src.Close()

	extension := filepath.Ext(image.Filename)
	if extension != ".jpg" && extension != ".jpeg" {
		return nil, fmt.Errorf("разрешены только изображения формата jpeg")
	}
	imageName := UUID + extension

	_, err = app.minioClient.PutObject(c, app.config.BucketName, imageName, src, image.Size, minio.PutObjectOptions{
		ContentType: "image/jpeg",
	})
	if err != nil {
		return nil, err
	}
	imageURL := fmt.Sprintf("%s/%s/%s", app.config.MinioEndPoint, app.config.BucketName, imageName)
	return &imageURL, nil
}

func (app *Application) getClient() string {
	return "2d217868-ab6d-41fe-9b34-7809083a2e8a"
}

func (app *Application) getModerator() *string {
	moderaorId := "87d54d58-1e24-4cca-9c83-bd2523902729"
	return &moderaorId
}

func (app *Application) GetAllCargo(c *gin.Context) {
	var request schemes.GetAllCargosRequest
	err := c.ShouldBindQuery(&request)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	// Фильтр по цене
	var lowprice uint
	var highprice uint
	lowprice64, err1 := strconv.ParseUint(request.LowPrice, 10, 64)
	highprice64, err2 := strconv.ParseUint(request.HighPrice, 10, 64)
	if err1 != nil {
		lowprice, _ = app.repo.GetLowestPrice()
	} else {
		lowprice = uint(lowprice64)
	}
	if err2 != nil {
		highprice, _ = app.repo.GetHighestPrice()
	} else {
		highprice = uint(highprice64)
	}

	cargos, err := app.repo.GetFilteredCargo(request.Name, lowprice, highprice)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	draftFlight, err := app.repo.GetDraftFlight(app.getClient())
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	response := schemes.GetAllCargosResponse{DraftFlight: nil, Cargos: cargos}
	if draftFlight != nil {
		response.DraftFlight = &schemes.FlightShort{UUID: draftFlight.UUID}
		flightCargos, err := app.repo.GetFlightCargos(draftFlight.UUID)
		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}
		response.DraftFlight.CargoCount = len(flightCargos)
	}

	c.JSON(http.StatusOK, response)
}

func (app *Application) GetCargo(c *gin.Context) {
	var request schemes.CargoRequest
	err := c.ShouldBindUri(&request)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	cargo, err := app.repo.GetCargoByID(request.CargoId)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	if cargo == nil {
		c.AbortWithError(http.StatusNotFound, fmt.Errorf("груз не найден"))
		return
	}

	c.JSON(http.StatusOK, cargo)
}

func (app *Application) DeleteCargo(c *gin.Context) {
	var request schemes.CargoRequest
	err := c.ShouldBindUri(&request)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	cargo, err := app.repo.GetCargoByID(request.CargoId)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	if cargo == nil {
		c.AbortWithError(http.StatusNotFound, fmt.Errorf("груз не найден"))
		return
	}

	cargo.IsDeleted = true
	err = app.repo.SaveCargo(cargo)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.Status(http.StatusOK)
}

func (app *Application) AddCargo(c *gin.Context) {
	var request schemes.AddCargoRequest
	err := c.ShouldBind(&request)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	cargo := ds.Cargo(request.Cargo)
	err = app.repo.AddCargo(&cargo)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	if request.Image != nil {
		imageURL, err := app.uploadImage(c, request.Image, cargo.UUID)
		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}
		cargo.Photo = imageURL
	}

	err = app.repo.SaveCargo(&cargo)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.Status(http.StatusOK)
}

func (app *Application) ChangeCargo(c *gin.Context) {
	var request schemes.ChangeCargoRequest
	err := c.ShouldBindUri(&request)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	err = c.ShouldBind(&request)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	cargo, err := app.repo.GetCargoByID(request.CargoId)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	if cargo == nil {
		c.AbortWithError(http.StatusNotFound, fmt.Errorf("груз не найден"))
		return
	}

	if request.Name != nil {
		cargo.Name = *request.Name
	}
	if request.EnglishName != nil {
		cargo.EnglishName = *request.EnglishName
	}
	if request.Image != nil {
		imageURL, err := app.uploadImage(c, request.Image, cargo.UUID)
		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}
		cargo.Photo = imageURL
	}
	if request.Category != nil {
		cargo.Category = *request.Category
	}
	if request.Price != nil {
		cargo.Price = *request.Price
	}
	if request.Weight != nil {
		cargo.Weight = *request.Weight
	}
	if request.Capacity != nil {
		cargo.Capacity = *request.Capacity
	}
	if request.Description != nil {
		cargo.Description = *request.Description
	}

	err = app.repo.SaveCargo(cargo)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, cargo)
}

func (app *Application) AddToFlight(c *gin.Context) {
	var request schemes.AddToFlightRequest
	err := c.ShouldBindUri(&request)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	cargo, err := app.repo.GetCargoByID(request.CargoId)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	if cargo == nil {
		c.AbortWithError(http.StatusNotFound, fmt.Errorf("груз не найден"))
		return
	}

	flight, err := app.repo.GetDraftFlight(app.getClient())
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	if flight == nil {
		flight, err = app.repo.CreateDraftFlight(app.getClient())
		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}
	}

	err = app.repo.AddToFlight(flight.UUID, request.CargoId, request.Quantity)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	cargos, err := app.repo.GetFlightCargos(flight.UUID)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, schemes.AllCargosResponse{Cargos: cargos})
}

func (app *Application) GetAllFlights(c *gin.Context) {
	var request schemes.GetAllFlightsRequest
	err := c.ShouldBindQuery(&request)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	flights, err := app.repo.GetAllFlights(request.FormDateStart, request.FormDateEnd, request.Status)
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

func (app *Application) GetFlight(c *gin.Context) {
	var request schemes.FlightRequest
	err := c.ShouldBindUri(&request)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	flight, err := app.repo.GetFlightById(request.FlightId, app.getClient())
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	if flight == nil {
		c.AbortWithError(http.StatusNotFound, fmt.Errorf("полет не найден"))
		return
	}

	cargos, err := app.repo.GetFlightCargos(request.FlightId)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, schemes.FlightResponse{Flight: schemes.ConvertFlight(flight), Cargos: cargos})
}

func (app *Application) UpdateFlight(c *gin.Context) {
	var request schemes.UpdateFlightRequest
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

	flight, err := app.repo.GetFlightById(request.URI.FlightId, app.getClient())
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	if flight == nil {
		c.AbortWithError(http.StatusNotFound, fmt.Errorf("полет не найден"))
		return
	}

	flight.RocketType = request.RocketType
	err = app.repo.SaveFlight(flight)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, schemes.UpdateFlightResponse{Flight: schemes.ConvertFlight(flight)})
}

func (app *Application) DeleteFlight(c *gin.Context) {
	var request schemes.FlightRequest
	err := c.ShouldBindUri(&request)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	flight, err := app.repo.GetFlightById(request.FlightId, app.getClient())
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	if flight == nil {
		c.AbortWithError(http.StatusNotFound, fmt.Errorf("полет не найден"))
		return
	}
	flight.Status = ds.DELETED

	err = app.repo.SaveFlight(flight)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.Status(http.StatusOK)
}

func (app *Application) DeleteFromFlight(c *gin.Context) {
	var request schemes.DeleteFromFlightRequest
	err := c.ShouldBindUri(&request)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	flight, err := app.repo.GetFlightById(request.FlightId, app.getClient())
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	if flight == nil {
		c.AbortWithError(http.StatusNotFound, fmt.Errorf("полет не найден"))
		return
	}

	if !strings.EqualFold(flight.Status, ds.DRAFT) {
		c.AbortWithError(http.StatusMethodNotAllowed, fmt.Errorf("нельзя редактировать полет статуса %s", flight.Status))
		return
	}

	err = app.repo.DeleteFromFlight(request.FlightId, request.CargoId)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	cargos, err := app.repo.GetFlightCargos(request.FlightId)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, schemes.AllCargosResponse{Cargos: cargos})
}

func (app *Application) UserConfirm(c *gin.Context) {
	var request schemes.UserConfirmRequest
	err := c.ShouldBindUri(&request)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	flight, err := app.repo.GetFlightById(request.FlightId, app.getClient())
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	if flight == nil {
		c.AbortWithError(http.StatusNotFound, fmt.Errorf("полет не найден"))
		return
	}

	if flight.Status != ds.DRAFT {
		c.AbortWithError(http.StatusMethodNotAllowed, fmt.Errorf("нельзя изменить полет статуса %s", flight.Status))
		return
	}

	flight.Status = ds.FORMED
	now := time.Now()
	flight.FormationDate = &now

	err = app.repo.SaveFlight(flight)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.Status(http.StatusOK)
}

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

	if request.Status != ds.REJECTED && request.Status != ds.COMPLETED {
		c.AbortWithError(http.StatusBadRequest, fmt.Errorf("нельзя изменить статус на %s", request.Status))
		return
	}

	flight, err := app.repo.GetFlightById(request.URI.FlightId, app.getClient())
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	if flight == nil {
		c.AbortWithError(http.StatusNotFound, fmt.Errorf("полет не найден"))
		return
	}

	if flight.Status != ds.FORMED {
		c.AbortWithError(http.StatusMethodNotAllowed, fmt.Errorf("нельзя изменить статус \"%s\" на \"%s\"", flight.Status, request.Status))
		return
	}

	flight.Status = request.Status
	flight.ModeratorId = app.getModerator()
	if request.Status == ds.COMPLETED {
		now := time.Now()
		flight.CompletionDate = &now
	}

	err = app.repo.SaveFlight(flight)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.Status(http.StatusOK)
}
