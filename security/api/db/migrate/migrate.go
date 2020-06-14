package migrate

import (
	"github.com/jinzhu/gorm"
)

func Load(db *gorm.DB) {
	CreateUserTable(db)
}
