package migrate

import (
	"log"

	"github.com/deivisson/apstore/api/models"
	"github.com/jinzhu/gorm"
)

func CreateUserTable(db *gorm.DB) {
	err := db.Debug().AutoMigrate(&models.User{}).Error
	if err != nil {
		log.Fatalf("cannot migrate table: %v", err)
	}
}
