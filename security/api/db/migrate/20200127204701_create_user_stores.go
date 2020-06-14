package migrate

import (
	"log"

	"github.com/deivisson/apstore/api/models"
	"github.com/jinzhu/gorm"
)

func CreateUserStoresTable(db *gorm.DB) {
	var err error
	err = db.Debug().AutoMigrate(&models.StoresUsers{}).Error
	if err != nil {
		log.Fatalf("cannot migrate table: %v", err)
	}

	err = db.Debug().Model(&models.StoresUsers{}).AddForeignKey("user_id", "users(id)", "RESTRICT", "RESTRICT").Error
	if err != nil {
		log.Fatalf("attaching foreign key error: %v", err)
	}

	err = db.Debug().Model(&models.StoresUsers{}).AddForeignKey("store_id", "stores(id)", "RESTRICT", "RESTRICT").Error
	if err != nil {
		log.Fatalf("attaching foreign key error: %v", err)
	}

	err = db.Debug().Model(&models.StoresUsers{}).AddUniqueIndex("user_store_unique_index", []string{"user_id", "store_id"}...).Error
	if err != nil {
		log.Fatalf("attaching unique index error: %v", err)
	}
}
