package initializers

import (
	"log"

	"github.com/jinzhu/gorm"
	"github.com/yiziz/gin-template/app/models"
	"github.com/yiziz/gin-template/db"
)

// InitializeDatabase opens and sets the DB
func InitializeDatabase(appEnv string) {
	db := ConnectDB(appEnv)
	models.SetDB(db)
}

// ConnectDB opens up a connection with db and returns it
func ConnectDB(appEnv string) *gorm.DB {
	db, err := db.InitDB(db.Config(appEnv))
	if err != nil {
		log.Fatal(err)
	}
	return db
}
