package repository

type DataRepository interface {
	UpdateData(water, wind int) error
}
