package model

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB
type Tinu struct{
	ID      string `json:"id" gorm:"primaryKey"`
	URL     string `json:"url" gorm:"not null"`
	Clicked uint64 `json:"clicked"`
}

func Setup() {
	// set your own data:
	dsn := "host=host user=user password=password dbname=name port=port sslmode=disable"

	var err error
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	err = db.AutoMigrate(&Tinu{})
	if err != nil {
		fmt.Println(err)
	}
}