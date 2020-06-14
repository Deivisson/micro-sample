package migrate

import (
	"log"

	"github.com/deivisson/apstore/api/models"
	"github.com/jinzhu/gorm"
)

func CreateStoreTable(db *gorm.DB) {
	var err error

	err = db.Debug().AutoMigrate(&models.Store{}).Error
	if err != nil {
		log.Fatalf("cannot migrate table: %v", err)
	}

	err = db.Debug().Model(&models.Store{}).AddForeignKey("user_id", "users(id)", "RESTRICT", "RESTRICT").Error
	if err != nil {
		log.Fatalf("attaching foreign key error: %v", err)
	}
}
