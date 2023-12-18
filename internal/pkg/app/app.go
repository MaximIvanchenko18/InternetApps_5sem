package app

import (
	"fmt"
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"

	"InternetApps_5sem/internal/app/config"
	"InternetApps_5sem/internal/app/dsn"
	"InternetApps_5sem/internal/app/redis"
	"InternetApps_5sem/internal/app/repository"
	"InternetApps_5sem/internal/app/role"

	_ "InternetApps_5sem/docs"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Application struct {
	repo        *repository.Repository
	config      *config.Config
	minioClient *minio.Client
	redisClient *redis.Client
}

func (app *Application) StartServer() {
	log.Println("Server up")

	r := gin.Default()
	r.Use(ErrorHandler())

	api := r.Group("/api")
	{
		// ГРУЗЫ (услуги)
		c := api.Group("/cargo")
		{
			c.GET("", app.WithAuthCheck(role.NotAuthorized, role.Customer, role.Moderator), app.GetAllCargo)        // Отфильтрованный список
			c.GET("/:cargo_id", app.WithAuthCheck(role.NotAuthorized, role.Customer, role.Moderator), app.GetCargo) // Один груз
			c.DELETE("/:cargo_id", app.WithAuthCheck(role.Moderator), app.DeleteCargo)                              // Удаление груза
			c.PUT("/:cargo_id", app.WithAuthCheck(role.Moderator), app.ChangeCargo)                                 // Изменение груза
			c.POST("", app.WithAuthCheck(role.Moderator), app.AddCargo)                                             // Добавление груза
			c.POST("/:cargo_id/add_to_flight", app.WithAuthCheck(role.Customer, role.Moderator), app.AddToFlight)   // Добавление в полет
		}
		// ПОЛЕТЫ (заявки)
		f := api.Group("/flights")
		{
			f.GET("", app.WithAuthCheck(role.Customer, role.Moderator), app.GetAllFlights)                                              // Отфильтрованный список (интервал дат и статус)
			f.GET("/:flight_id", app.WithAuthCheck(role.Customer, role.Moderator), app.GetFlight)                                       // Один полет
			f.PUT("/:flight_id", app.WithAuthCheck(role.Customer, role.Moderator), app.UpdateFlight)                                    // Изменение полета (тип ракеты)
			f.DELETE("/:flight_id", app.WithAuthCheck(role.Customer, role.Moderator), app.DeleteFlight)                                 // Удаление полета
			f.DELETE("/:flight_id/delete_cargo/:cargo_id", app.WithAuthCheck(role.Customer, role.Moderator), app.DeleteFromFlight)      // Удаление груза из заявки
			f.PUT("/:flight_id/change_cargo/:cargo_id", app.WithAuthCheck(role.Customer, role.Customer), app.UpdateFlightCargoQuantity) // Изменение кол-ва груза в полете
			f.PUT("/user_confirm", app.WithAuthCheck(role.Customer, role.Moderator), app.UserConfirm)                                   // Сформировать создателем
			f.PUT("/:flight_id/moderator_confirm", app.WithAuthCheck(role.Moderator), app.ModeratorConfirm)                             // Завершить/отклонить модератором
			f.PUT("/:flight_id/shipment", app.Shipment)
		}
		// ПОЛЬЗОВАТЕЛИ (авторизация)
		u := api.Group("/user")
		{
			u.POST("/sign_up", app.Register)
			u.POST("/login", app.Login)
			u.POST("/logout", app.Logout)
		}
	}

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.Run(fmt.Sprintf("%s:%d", app.config.ServiceHost, app.config.ServicePort))

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

	app.minioClient, err = minio.New(app.config.Minio.Endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4("", "", ""),
		Secure: false,
	})
	if err != nil {
		return nil, err
	}

	app.redisClient, err = redis.New(app.config.Redis)
	if err != nil {
		return nil, err
	}

	return &app, nil
}
