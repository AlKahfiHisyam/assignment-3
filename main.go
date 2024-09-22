package main

import (
	"assignment-3/database"
	"assignment-3/repository"
	"assignment-3/service"
	"log"
)

func main() {
	database.StartDB()

	db := database.GetDB()
	dataRepo := repository.NewDataRepository(db)
	updater := service.NewDataUpdater(dataRepo)

	go updater.UpdateData()

	log.Println("Service is running...")
	select {}
}
