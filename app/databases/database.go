package databases

import (
	"log"

	"github.com/surajNirala/rating_services/app/config"
	"github.com/surajNirala/rating_services/app/models"
)

func DatabaseUp() {
	DB := config.DB
	err := DB.AutoMigrate(
		&models.Rating{},
	)
	if err != nil {
		log.Fatalf("Error migrating the database: %v", err)
	}
}
