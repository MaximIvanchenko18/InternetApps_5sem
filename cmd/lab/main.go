package main

import (
	"log"

	"InternetApps_5sem/internal/api"
)

func main() {
	log.Println("Application start!")
	api.StartServer()
	log.Println("Application terminated!")
}
