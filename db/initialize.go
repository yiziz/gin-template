package db

import (
	_ "github.com/go-sql-driver/mysql" // Needed for opening connection to mysql
	"github.com/jinzhu/gorm"
)

var globalDB *gorm.DB

// OpenDatabaseConnection opens a db connection to a database
func OpenDatabaseConnection(adapter, parameters string) (*gorm.DB, error) {
	db, err := gorm.Open(adapter, parameters)
	return db, err
}

// SetDB sets the models.db variable
func SetDB(initializedDB *gorm.DB) {
	globalDB = initializedDB
}

// InitDB opens db connection and set globalDB
func InitDB(adapter, parameters string) (*gorm.DB, error) {
	db, err := OpenDatabaseConnection(adapter, parameters)
	SetDB(db)
	return db, err
}

// DB returns the globalDB
func DB() *gorm.DB {
	return globalDB
}
