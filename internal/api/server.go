package api

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type cargo struct {
	Photo       string  `json:"photo src"`
	Category    string  `json:"category"`
	Name        string  `json:"name"`
	Price       int32   `json:"price"`
	Unit        string  `json:"unit"`
	Description string  `json:"description"`
	Rating      float32 `json:"rating"`
}

var goods = []cargo{
	{Photo: "/image/oxygen.jpg", Category: "Кислород", Name: "Кислород", Price: 390, Unit: "1 литр", Description: "Кислород в баллонах под высоким давлением", Rating: 5.0},
	{Photo: "/image/water.jpg", Category: "Напитки", Name: "Вода", Price: 72, Unit: "1 литр", Description: "Вода родниковая очищенная", Rating: 4.9},
	{Photo: "/image/tea.jpg", Category: "Напитки", Name: "Чай черный", Price: 750, Unit: "3 грамма", Description: "Чай черный цейлонский без сахара в специализированном пакете", Rating: 5.0},
	{Photo: "/image/coffee.jpg", Category: "Напитки", Name: "Кофе с молоком и сахаром", Price: 1140, Unit: "30 граммов", Description: "Кофе натуральный Arabica, натуральное молоко, сахар-песок", Rating: 4.8},
	{Photo: "/image/kisel.jpg", Category: "Напитки", Name: "Кисель вишневый", Price: 25, Unit: "120 граммов", Description: "Сахар-песок, крахмал, сок вишневый концентрированный, витаминная смесь, кислота лимонная", Rating: 4.2},
	{Photo: "/image/bread.jpg", Category: "Еда", Name: "Хлеб пшеничный", Price: 380, Unit: "30 граммов", Description: "Мука пшеничная, вода, маргарин, сахар, дрожжи, соль, молоко сухое", Rating: 5.0},
	{Photo: "/image/meat_grecha.jpg", Category: "Еда", Name: "Каша гречневая с мясом", Price: 15, Unit: "60 граммов", Description: "Крупа гречневая, соль, жир, фарш говяжий сушеный, лук сушеный, аромат говядины", Rating: 4.8},
	{Photo: "/image/borsh.jpg", Category: "Еда", Name: "Борщ с мясом", Price: 1100, Unit: "30 граммов", Description: "Мясо-говядина крупнокусковое, картофель, капуста свежая, свекла, морковь, лук репчатый, корень петрушки, томатная паста, пюре из перца, сахар-песок, масло топленое, соль поваренная, лимонная кислота, лист лавровый, перец черный молотый, бульон", Rating: 4.7},
	{Photo: "/image/biscuits.jpg", Category: "Еда", Name: "Печенье Восток", Price: 510, Unit: "30 граммов", Description: "Мука высшего сорта, крахмал маисовый, сахарная пудра, инвертный сироп, маргарин, молоко цельное, ванильная пудра, соль, сода, аммоний, эссенция", Rating: 5.0},
	{Photo: "/image/personal.jpg", Category: "Личные вещи", Name: "Посылка от родных", Price: 0, Unit: "1 штука", Description: "Личные посылки от родственников и друзей космонавтов", Rating: 5.0},
}

func StartServer() {
	log.Println("Server start up")

	r := gin.Default()

	r.Static("/image", "./resources")  // images
	r.Static("/styles", "./templates") // css-files
	r.LoadHTMLGlob("templates/*")

	r.GET("/cargo", func(c *gin.Context) {
		c.HTML(http.StatusOK, "start.tmpl", goods)
	})

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")

	log.Println("Server down")
}
