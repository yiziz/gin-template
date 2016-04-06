package migrate

import (
	"time"

	"github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

// Sample table structure
type Sample struct {
	// Model
	ID        uint `gorm:"primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt mysql.NullTime
}

func init() {
	filename := "000001_sample_create_table"
	migrationList = append(migrationList, &Migration{
		Name: filename,
		Function: func(db *gorm.DB) error {
			return db.CreateTable(&Sample{}).Error
		},
	})
}
