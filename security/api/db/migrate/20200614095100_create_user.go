package migrate

import (
	"log"

	"github.com/deivisson/micro-sample/security/api/models"
	"github.com/jinzhu/gorm"
)

func CreateUserTable(db *gorm.DB) {
	err := db.Debug().AutoMigrate(&models.User{}).Error
	if err != nil {
		log.Fatalf("cannot migrate table: %v", err)
	}
}
