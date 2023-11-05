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
		button_str := c.Query("Filter")

		var lowprice uint
		var highprice uint
		if button_str != "" {
			lowprice64, err1 := strconv.ParseUint(lowprice_str, 10, 64)
			highprice64, err2 := strconv.ParseUint(highprice_str, 10, 64)
			if err1 != nil {
				lowprice, _ = a.repo.GetLowestPrice()
			} else {
				lowprice = uint(lowprice64)
			}
			if err2 != nil {
				highprice, _ = a.repo.GetHighestPrice()
			} else {
				highprice = uint(highprice64)
			}
		} else {
			name = ""
			lowprice, _ = a.repo.GetLowestPrice()
			highprice, _ = a.repo.GetHighestPrice()
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

		lowprice, _ := a.repo.GetLowestPrice()
		highprice, _ := a.repo.GetHighestPrice()

		goods, err := a.repo.GetFilteredCargo("", lowprice, highprice)

		if err != nil {
			log.Println("Goods can not be filtered!")
			c.Error(err)
			return
		}
		c.HTML(http.StatusOK, "cargo.tmpl", gin.H{
			"Cargo":      goods,
			"searchName": "",
			"lowPrice":   lowprice,
			"highPrice":  highprice,
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
