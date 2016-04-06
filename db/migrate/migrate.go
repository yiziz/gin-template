package migrate

import (
	"github.com/jinzhu/gorm"

	"fmt"
	"time"
)

// SchemaMigration used to see which migrations have run by storing those that have run
type SchemaMigration struct {
	MigrationName string    `sql:"unique_index"`
	CreatedAt     time.Time `sql:"DEFAULT:current_timestamp"`
}

// MigrationFunctionType is func that is called to run the migration for a table
type MigrationFunctionType func(db *gorm.DB) error

// Migration model
type Migration struct {
	Name     string
	Function MigrationFunctionType
}

var migrationList []*Migration

func createSchemaMigrationTable(db *gorm.DB) {
	db.AutoMigrate(&SchemaMigration{})
}

func runStandardMigrations(db *gorm.DB) {
	for _, migration := range migrationList {
		migrationName := migration.Name
		migrationFunction := migration.Function
		if !isAlreadyMigrated(db, migrationName, migrationFunction) {
			migrate(db, migrationName, migrationFunction)
		}
	}
}

func isAlreadyMigrated(db *gorm.DB, migrationName string, migrationFunction MigrationFunctionType) bool {
	var rowCount int
	db.Model(SchemaMigration{}).Where("migration_name = ?", migrationName).Count(&rowCount)
	return rowCount > 0
}

func migrate(db *gorm.DB, migrationName string, migrationFunction MigrationFunctionType) {
	fmt.Println("Migrating", migrationName)
	migrationError := migrationFunction(db)
	if migrationError == nil {
		fmt.Println("Finished migrating", migrationName)
		db.Create(&SchemaMigration{MigrationName: migrationName})
	} else {
		fmt.Println("Error migrating", migrationName)
	}
}

// RunMigrations runs all migrations
func RunMigrations(db *gorm.DB) {
	createSchemaMigrationTable(db)
	runStandardMigrations(db)
}
