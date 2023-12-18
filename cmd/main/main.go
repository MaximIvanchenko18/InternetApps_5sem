package main

import (
	"log"

	"InternetApps_5sem/internal/pkg/app"
)

// TODO: change
// @title Cargo transfer to ISS
// @version 1.0

// @host 127.0.0.1:8000
// @schemes http
// @BasePath /

func main() {
	app, err := app.New()
	if err != nil {
		log.Println("application can not be started", err)
		return
	}
	app.StartServer()
}
