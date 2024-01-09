package app

import (
	"fmt"
	"math"
	"net/http"
	"strconv"

	_ "InternetApps_5sem/docs"
	"InternetApps_5sem/internal/app/ds"
	"InternetApps_5sem/internal/app/schemes"

	"github.com/gin-gonic/gin"
)

// @Summary		Получить все грузы
// @Tags		Грузы
// @Description	Возвращает все доступные грузы с опциональной фильтрацией по названию и диапазону цены
// @Produce		json
// @Param		name query string false "Название груза для фильтрации"
// @Param		low_price query string false "Нижний порог цены для фильтрации"
// @Param		high_price query string false "Верхний порог цены для фильтрации"
// @Success		200 {object} schemes.GetAllCargosResponse
// @Router		/api/cargo [get]
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

	response := schemes.GetAllCargosResponse{DraftFlight: nil, Cargos: cargos}
	userId, exists := c.Get("userId")
	if exists {
		draftFlight, err := app.repo.GetDraftFlight(userId.(string))
		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}
		if draftFlight != nil {
			response.DraftFlight = &draftFlight.UUID
		}
	}

	c.JSON(http.StatusOK, response)
}

// @Summary		Получить один груз
// @Tags		Грузы
// @Description	Возвращает подробную информацию об одном грузе
// @Produce		json
// @Param		cargo_id path string true "id груза"
// @Success		200 {object} ds.Cargo
// @Router		/api/cargo/{cargo_id} [get]
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

// @Summary		Удалить груз
// @Tags		Грузы
// @Description	Удаляет груз по id
// @Param		cargo_id path string true "id груза"
// @Success		200
// @Router		/api/cargo/{cargo_id} [delete]
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
	if cargo.Photo != nil {
		err = app.deleteImage(c, cargo.UUID)
		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}
	}

	cargo.Photo = nil
	cargo.IsDeleted = true
	err = app.repo.SaveCargo(cargo)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.Status(http.StatusOK)
}

// @Summary		Добавить груз
// @Tags		Грузы
// @Description	Добавить новый груз
// @Accept		mpfd
// @Param     	image formData file false "Изображение груза"
// @Param     	name formData string true "Название" format:"string" maxLength:100
// @Param     	en_name formData string true "Английское название" format:"string" maxLength:100
// @Param     	category formData string true "Категория" format:"string" maxLength:50
// @Param     	price formData uint true "Цена" format:"uint"
// @Param     	weight formData float32 true "Вес" format:"float32"
// @Param     	capacity formData float32 true "Объем" format:"float32"
// @Param     	description formData string true "Описание" format:"string" maxLength:500
// @Success		200 {object} string
// @Router		/api/cargo/ [post]
func (app *Application) AddCargo(c *gin.Context) {
	var request schemes.AddCargoRequest
	err := c.ShouldBind(&request)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	cargo := ds.Cargo(request.Cargo)
	cargo.Weight = math.Round(cargo.Weight*1000) / 1000
	cargo.Capacity = math.Round(cargo.Capacity*100000) / 100000
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

	c.JSON(http.StatusCreated, cargo.UUID)
}

// @Summary		Изменить груз
// @Tags		Грузы
// @Description	Изменить данные о грузе
// @Accept		mpfd
// @Produce		json
// @Param		cargo_id path string true "Идентификатор груза" format:"uuid"
// @Param     	image formData file false "Изображение груза"
// @Param     	name formData string false "Название" format:"string" maxLength:100
// @Param     	en_name formData string false "Английское название" format:"string" maxLength:100
// @Param     	category formData string false "Категория" format:"string" maxLength:50
// @Param     	price formData uint false "Цена" format:"uint"
// @Param     	weight formData float32 false "Вес" format:"float32"
// @Param     	capacity formData float32 false "Объем" format:"float32"
// @Param     	description formData string false "Описание" format:"string" maxLength:500
// @Success     200
// @Router		/api/cargo/{cargo_id} [put]
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
		if cargo.Photo != nil {
			err := app.deleteImage(c, cargo.UUID)
			if err != nil {
				c.AbortWithError(http.StatusInternalServerError, err)
				return
			}
		}
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

	c.Status(http.StatusOK)
}

// @Summary		Добавить груз в полет
// @Tags		Грузы
// @Description	Добавить выбранный груз в черновик полета
// @Param		cargo_id path string true "id груза"
// @Success		200
// @Router		/api/cargo/{cargo_id}/add_to_flight [post]
func (app *Application) AddToFlight(c *gin.Context) {
	var request schemes.AddToFlightRequest
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

	cargo, err := app.repo.GetCargoByID(request.URI.CargoId)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	if cargo == nil {
		c.AbortWithError(http.StatusNotFound, fmt.Errorf("груз не найден"))
		return
	}

	userId := getUserId(c)
	flight, err := app.repo.GetDraftFlight(userId)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	if flight == nil {
		flight, err = app.repo.CreateDraftFlight(userId)
		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}
	}

	err = app.repo.AddToFlight(flight.UUID, request.URI.CargoId, request.Quantity)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.Status(http.StatusOK)
}
