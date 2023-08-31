package db

import (
	"go-wedding/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"log"
)

func Init() *gorm.DB {

	dbURL := "postgres://Burnwood1911:1tpQ9axqmuKg@ep-snowy-wave-23053803.us-east-2.aws.neon.tech/neondb"

	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	db.AutoMigrate(&models.Invite{})

	return db

}
