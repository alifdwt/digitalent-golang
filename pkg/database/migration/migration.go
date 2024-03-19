package migration

import (
	"mygram-api/models"

	"gorm.io/gorm"
)

func RunMigration(db *gorm.DB) error {
	err := db.AutoMigrate(
		&models.User{},
		&models.Photo{},
		&models.Comment{},
		&models.SocialMedia{},
	)
	if err != nil {
		panic(err)
	}

	return err
}
