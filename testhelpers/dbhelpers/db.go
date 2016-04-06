package dbhelpers

import (
	"log"

	"github.com/yiziz/gin-template/db"
	"github.com/yiziz/gin-template/db/migrate"
	"github.com/jinzhu/gorm"
)

// DB opens a db connection and returns the &db
func DB() *gorm.DB {
	db, err := db.InitDB(db.Config("test"))
	if err != nil {
		log.Fatal(err)
	}
	// schema migrations
	migrate.RunMigrations(db)
	return db
}

// ClearTable deletes data from a table
func ClearTable(db *gorm.DB, model interface{}) {
	db.Unscoped().Delete(model)
}

// ClearDB deletes data from tables in db
func ClearDB(db *gorm.DB) {
	tx := db.Begin()
	// Using .Table instead of .Delete(models.Model{}) to avoid importing models
	// and causing circular imports when testing the models package
	// tx.Table("users").Unscoped().Delete("")
	tx.Commit()
}

// CloseDB clears the db and closes the connection
func CloseDB(db *gorm.DB) {
	ClearDB(db)
	db.Close()
}
