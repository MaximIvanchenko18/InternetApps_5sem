package api

import (
	"log"

	"strconv"
	"strings"

	"net/http"

	"github.com/gin-gonic/gin"
)

type cargo struct {
	EnglishName string  `json:"english name"` // for page with more info
	Photo       string  `json:"photo src"`
	Category    string  `json:"category"`
	Name        string  `json:"name"`
	Price       int32   `json:"price"`    // Rubles
	Weight      float32 `json:"weight"`   // kg
	Capacity    float32 `json:"capacity"` // m^3
	Description string  `json:"description"`
}

var goods = []cargo{
	{EnglishName: "oxygen", Photo: "/image/oxygen.jpg", Category: "Кислород", Name: "Кислород", Price: 19500, Weight: 60, Capacity: 0.05, Description: "Кислород в баллонах под высоким давлением"},
	{EnglishName: "water", Photo: "/image/water.jpg", Category: "Напитки", Name: "Вода", Price: 110, Weight: 0.5, Capacity: 0.0005, Description: "Вода родниковая очищенная"},
	{EnglishName: "tea", Photo: "/image/tea.jpg", Category: "Напитки", Name: "Чай черный", Price: 750, Weight: 0.003, Capacity: 0.0003, Description: "Чай черный цейлонский без сахара в специализированном пакете"},
	{EnglishName: "coffee", Photo: "/image/coffee.jpg", Category: "Напитки", Name: "Кофе с молоком и сахаром", Price: 1140, Weight: 0.03, Capacity: 0.0002, Description: "Кофе натуральный Arabica, натуральное молоко, сахар-песок"},
	{EnglishName: "kisel", Photo: "/image/kisel.jpg", Category: "Напитки", Name: "Кисель вишневый", Price: 25, Weight: 0.12, Capacity: 0.00005, Description: "Сахар-песок, крахмал, сок вишневый концентрированный, витаминная смесь, кислота лимонная"},
	{EnglishName: "bread", Photo: "/image/bread.jpg", Category: "Еда", Name: "Хлеб пшеничный", Price: 380, Weight: 0.03, Capacity: 0.00015, Description: "Мука пшеничная, вода, маргарин, сахар, дрожжи, соль, молоко сухое"},
	{EnglishName: "meat_and_grecha", Photo: "/image/meat_grecha.jpg", Category: "Еда", Name: "Каша гречневая с мясом", Price: 15, Weight: 0.06, Capacity: 0.00025, Description: "Крупа гречневая, соль, жир, фарш говяжий сушеный, лук сушеный, аромат говядины"},
	{EnglishName: "borsh", Photo: "/image/borsh.jpg", Category: "Еда", Name: "Борщ с мясом", Price: 1100, Weight: 0.03, Capacity: 0.0003, Description: "Мясо-говядина крупнокусковое, картофель, капуста свежая, свекла, морковь, лук репчатый, корень петрушки, томатная паста, пюре из перца, сахар-песок, масло топленое, соль поваренная, лимонная кислота, лист лавровый, перец черный молотый, бульон"},
	{EnglishName: "biscuits", Photo: "/image/biscuits.jpg", Category: "Еда", Name: "Печенье Восток", Price: 510, Weight: 0.03, Capacity: 0.0002, Description: "Мука высшего сорта, крахмал маисовый, сахарная пудра, инвертный сироп, маргарин, молоко цельное, ванильная пудра, соль, сода, аммоний, эссенция"},
	{EnglishName: "personal", Photo: "/image/personal.jpg", Category: "Личные вещи", Name: "Посылка от родных", Price: 0, Weight: 1, Capacity: 1, Description: "Личные посылки от родственников и друзей космонавтов"},
}

func StartServer() {
	log.Println("Server start up")

	r := gin.Default()

	r.Static("/image", "./resources")         // images
	r.Static("/styles", "./templates/styles") // css-files
	r.LoadHTMLGlob("templates/*.tmpl")

	r.GET("/cargo", func(c *gin.Context) {
		name := c.Query("Name")
		lowprice := c.Query("LowPrice")
		highprice := c.Query("HighPrice")
		filtered_goods := filterGoods(goods, name, lowprice, highprice)
		c.HTML(http.StatusOK, "cargo.tmpl", gin.H{
			"Cargo":      filtered_goods,
			"searchName": name,
			"lowPrice":   lowprice,
			"highPrice":  highprice,
		})
	})

	r.GET("/cargo/:name", func(c *gin.Context) {
		name := c.Param("name")
		for _, good := range goods {
			if good.EnglishName == name {
				c.HTML(http.StatusOK, "info.tmpl", good)
				return
			}
		}
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Good is not found"})
	})

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")

	log.Println("Server down")
}

func filterGoods(goods []cargo, name_str string, lowprice_str string, highprice_str string) []cargo {
	lowprice, err1 := strconv.Atoi(lowprice_str)
	highprice, err2 := strconv.Atoi(highprice_str)
	nameparts := strings.Split(name_str, " ")
	if (err1 != nil || lowprice <= 0) && err2 != nil && len(nameparts) <= 0 {
		return goods
	}
	var result []cargo
	for _, good := range goods {
		containAll := true
		for _, part := range nameparts {
			if !strings.Contains(good.Name, part) {
				containAll = false
				break
			}
		}
		if containAll && (err1 != nil || lowprice <= int(good.Price)) && (err2 != nil || highprice >= int(good.Price)) {
			result = append(result, good)
		}
	}
	return result
}
