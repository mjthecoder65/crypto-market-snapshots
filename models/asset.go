package models

type Assets struct {
	ID   uint   `db:"id" json:"id" gorm:"primary_key"`
	Name string `db:"name" json:"name"`
}
