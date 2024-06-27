package config

import (
	"log"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	dbUrl := os.Getenv("DB_URL")
	count := 10
	for retry := 0; retry < count; retry++ {
		database, err := gorm.Open(mysql.Open(dbUrl), &gorm.Config{})
		if err != nil {
			log.Printf("Failed to connect to database!")
			time.Sleep(5 * time.Second)
		} else {
			DB = database
			log.Printf("Connection success!")
			break
		}

	}

	if DB == nil {
		log.Fatal("Connection timeout to database")
	}
}
