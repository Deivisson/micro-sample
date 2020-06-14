package migrate

import (
	"github.com/jinzhu/gorm"
)

// var users = []models.User{
// 	models.User{
// 		Nickname: "Steven victor",
// 		Email:    "steven@gmail.com",
// 		Password: "password",
// 	},
// 	models.User{
// 		Nickname: "Martin Luther",
// 		Email:    "luther@gmail.com",
// 		Password: "password",
// 	},
// }

func Load(db *gorm.DB) {
	CreateUserTable(db)
	CreateStoreTable(db)
	CreateUserStoresTable(db)
}
