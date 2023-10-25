package main

import (
	"log"

	"InternetApps_5sem/internal/pkg/app"
)

func main() {
	app, err := app.New()
	if err != nil {
		log.Println("Application can not be started!", err)
		return
	}
	app.StartServer()
}
