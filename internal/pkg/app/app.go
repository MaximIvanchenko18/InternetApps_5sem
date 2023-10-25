package app

import (
	"log"
	"strconv"

	"net/http"

	"github.com/gin-gonic/gin"

	"InternetApps_5sem/internal/app/config"
	"InternetApps_5sem/internal/app/dsn"
	"InternetApps_5sem/internal/app/repository"
)

type Application struct {
	repo   *repository.Repository
	config *config.Config
}

func (a *Application) StartServer() {
	log.Println("Server start up")

	r := gin.Default()

	r.Static("/image", "./resources")         // images
	r.Static("/styles", "./templates/styles") // css-files
	r.LoadHTMLGlob("templates/*.tmpl")        // html-files

	r.GET("/cargo", func(c *gin.Context) {
		name := c.Query("Name")
		lowprice_str := c.Query("LowPrice")
		highprice_str := c.Query("HighPrice")

		lowprice, err1 := strconv.Atoi(lowprice_str)
		highprice, err2 := strconv.Atoi(highprice_str)
		if err1 != nil {
			lowprice = 0
		}
		if err2 != nil {
			highprice = 0
		}

		goods, err := a.repo.GetFilteredCargo(name, lowprice, highprice)

		if err != nil {
			log.Println("Goods can not be filtered!")
			c.Error(err)
			return
		}
		c.HTML(http.StatusOK, "cargo.tmpl", gin.H{
			"Cargo":      goods,
			"searchName": name,
			"lowPrice":   lowprice,
			"highPrice":  highprice,
		})
	})

	r.GET("/cargo/:name", func(c *gin.Context) {
		name := c.Param("name")
		good, err := a.repo.GetCargoByEnName(name)
		if err != nil {
			log.Println("Good can not be displayed!")
			c.Error(err)
			return
		}
		c.HTML(http.StatusOK, "info.tmpl", *good)
	})

	r.POST("/cargo", func(c *gin.Context) {
		id_str := c.PostForm("delete")

		id, err := strconv.Atoi(id_str)
		if err != nil {
			log.Println("Good can not be deleted!")
			c.Error(err)
			return
		}

		a.repo.DeleteCargoById(id)

		goods, err := a.repo.GetFilteredCargo("", 0, 0)

		if err != nil {
			log.Println("Goods can not be filtered!")
			c.Error(err)
			return
		}
		c.HTML(http.StatusOK, "cargo.tmpl", gin.H{
			"Cargo":      goods,
			"searchName": "",
			"lowPrice":   0,
			"highPrice":  0,
		})
	})

	r.Run("0.0.0.0:9000") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")

	log.Println("Server down")
}

func New() (*Application, error) {
	var err error
	app := Application{}
	app.config, err = config.NewConfig()
	if err != nil {
		return nil, err
	}

	app.repo, err = repository.New(dsn.FromEnv())
	if err != nil {
		return nil, err
	}

	return &app, nil
}
