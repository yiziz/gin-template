package models

import (
	"time"

	"github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

// ScopeFunc is used to decorate db queries for .Scopes
type ScopeFunc func(db *gorm.DB) *gorm.DB

var modelDB *gorm.DB

// SetDB sets the globalDB
func SetDB(db *gorm.DB) {
	modelDB = db
}

// DB returns globalDB
func DB() *gorm.DB {
	return modelDB
}

// Model is the base model every model should inherit from
type Model struct {
	ID        uint64         `gorm:"primary_key" json:"id"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt mysql.NullTime `json:"-"`
}
