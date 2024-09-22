package model

type DataModel struct {
	Id    int `gorm:"primaryKey" json:"id"`
	Water int `json:"water"`
	Wind  int `json:"wind"`
}
