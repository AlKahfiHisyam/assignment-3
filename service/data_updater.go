package service

import (
	"assignment-3/helper"
	"assignment-3/repository"
	"log"
	"math/rand"
	"time"
)

type DataUpdater struct {
	Repo repository.DataRepository
}

func NewDataUpdater(repo repository.DataRepository) *DataUpdater {
	return &DataUpdater{Repo: repo}
}

func (u *DataUpdater) UpdateData() {
	for {
		water := rand.Intn(100) + 1
		wind := rand.Intn(100) + 1

		err := u.Repo.UpdateData(water, wind)
		if err != nil {
			log.Println("Error saat update data:", err)
			continue
		}

		windStatus := helper.WindStatus(wind)
		waterStatus := helper.WaterStatus(water)

		log.Printf("Wind: %d m/s, Status: %s", wind, windStatus)
		log.Printf("Water: %d m, Status: %s", water, waterStatus)

		time.Sleep(15 * time.Second)
	}
}
