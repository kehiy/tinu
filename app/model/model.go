package model

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB
type Tinu struct{
	ID      string `json:"id" gorm:"primaryKey"`
	URL     string `json:"url"`
	Clicked uint64 `json:"clicked"`
}

func Setup() {
	// set your own data:
	dsn := "host=luca.iran.liara.ir user=root password=jXU5nAEA8ln70yvEybX8CNKs dbname=tinu port=34930 sslmode=disable"
	gorm.Open(postgres.Open(dsn), &gorm.Config{})
}