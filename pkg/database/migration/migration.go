package migration

import (
	"log"

	"github.com/alifdwt/digitalent-golang/models"
	"gorm.io/gorm"
)

func RunMigration(db *gorm.DB) error {
	err := db.AutoMigrate(
		&models.Order{},
		&models.Items{},
	)
	if err != nil {
		log.Fatal("Migration failed", err)
	}

	log.Println("Migration success")

	return err
}
