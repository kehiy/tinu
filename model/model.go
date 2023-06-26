package model

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB
type Tinu struct{
	ID      string `json:"id" gorm:"primaryKey"`
	URL     string `json:"url" gorm:"not null"`
	Clicked uint64 `json:"clicked"`
}

func loadEnv() (string, string, string, string, string) {
	err := godotenv.Load()
    if err != nil {
        fmt.Println("Error loading .env file:", err)
        os.Exit(1)
    }

    // Get the value of an environment variable and print it
    dbHost := os.Getenv("DBHOST")
    dbPort := os.Getenv("DBPORT")
    dbName := os.Getenv("DBNAME")
    dbPass := os.Getenv("DBPASS")
    dbUser := os.Getenv("DBUSER")
    return dbHost, dbPort, dbName, dbPass, dbUser
}

func Setup() {
	dbHost, dbPort, dbName, dbPass, dbUser := loadEnv()
	// set your own data:
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", dbHost, dbUser, dbPass, dbName, dbPort)

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