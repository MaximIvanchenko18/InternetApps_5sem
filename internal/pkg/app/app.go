package app

import (
	"log"
	"time"

	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"

	"InternetApps_5sem/internal/app/config"
	"InternetApps_5sem/internal/app/dsn"
	"InternetApps_5sem/internal/app/repository"
)

type Application struct {
	repo        *repository.Repository
	config      *config.Config
	minioClient *minio.Client
}

func (app *Application) StartServer() {
	log.Println("Server up")

	r := gin.Default()
	r.Use(ErrorHandler())

	r.Static("/image", "./resources")         // images
	r.Static("/styles", "./templates/styles") // css-files

	// ГРУЗЫ (услуги)
	r.GET("/api/cargo", app.GetAllCargo)                          // Отфильтрованный список
	r.GET("/api/cargo/:cargo_id", app.GetCargo)                   // Один груз
	r.DELETE("/api/cargo/:cargo_id", app.DeleteCargo)             // Удаление груза
	r.PUT("/api/cargo/:cargo_id", app.ChangeCargo)                // Изменение груза
	r.POST("/api/cargo", app.AddCargo)                            // Добавление груза
	r.POST("/api/cargo/:cargo_id/add_to_flight", app.AddToFlight) // Добавление в полет

	// ПОЛЕТЫ (заявки)
	r.GET("/api/flights", app.GetAllFlights)                                               // Отфильтрованный список (интервал дат и статус)
	r.GET("/api/flights/:flight_id", app.GetFlight)                                        // Один полет
	r.PUT("/api/flights/:flight_id/update", app.UpdateFlight)                              // Изменение полета (тип ракеты)
	r.DELETE("/api/flights/:flight_id", app.DeleteFlight)                                  // Удаление полета
	r.DELETE("/api/flights/:flight_id/delete_cargo/:cargo_id", app.DeleteFromFlight)       // Удаление груза из заявки
	r.PUT("/api/flights/:flight_id/change_cargo/:cargo_id", app.UpdateFlightCargoQuantity) // Изменение кол-ва груза в полете
	r.PUT("/api/flights/:flight_id/user_confirm", app.UserConfirm)                         // Сформировать создателем
	r.PUT("/api/flights/:flight_id/moderator_confirm", app.ModeratorConfirm)               // Завершить/отклонить модератором

	r.Run("0.0.0.0:7000") // listen and serve on 0.0.0.0:7000 (for windows "localhost:7000")

	log.Println("Server down")
}

func New() (*Application, error) {
	var err error
	loc, _ := time.LoadLocation("UTC")
	time.Local = loc
	app := Application{}
	app.config, err = config.NewConfig()
	if err != nil {
		return nil, err
	}

	app.repo, err = repository.New(dsn.FromEnv())
	if err != nil {
		return nil, err
	}

	app.minioClient, err = minio.New(app.config.MinioEndPoint, &minio.Options{
		Creds:  credentials.NewStaticV4("", "", ""),
		Secure: false,
	})
	if err != nil {
		return nil, err
	}

	return &app, nil
}

func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		for _, err := range c.Errors {
			log.Println(err.Err)
		}

		lastError := c.Errors.Last()
		if lastError != nil {
			switch c.Writer.Status() {
			case http.StatusBadRequest:
				c.JSON(-1, gin.H{"error": "wrong request"})
			case http.StatusNotFound:
				c.JSON(-1, gin.H{"error": lastError.Error()})
			case http.StatusMethodNotAllowed:
				c.JSON(-1, gin.H{"error": lastError.Error()})
			default:
				c.Status(-1)
			}
		}
	}
}
