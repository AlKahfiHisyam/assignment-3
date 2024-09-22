package repository

import (
	"assignment-3/model"

	"gorm.io/gorm"
)

type DataRepositoryImpl struct {
	DB *gorm.DB
}

func NewDataRepository(db *gorm.DB) DataRepository {
	return &DataRepositoryImpl{DB: db}
}

func (r *DataRepositoryImpl) UpdateData(water, wind int) error {
	data := model.DataModel{Water: water, Wind: wind}
	result := r.DB.Create(&data)
	return result.Error
}
