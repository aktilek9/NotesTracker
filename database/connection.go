package database

import (
	"NotesTracker/modules"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DBConfig struct {
	Host     string
	User     string
	Password string
	DBName   string
	Port     string
	SSLMode  string
	TimeZone string
}

func NewDBConnection(config DBConfig) (*gorm.DB, error) {
	dsn := buildDSN(config)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	if err := runMigrations(db); err != nil {
		return nil, err
	}

	return db, nil
}

func buildDSN(config DBConfig) string {
	return "host=" + config.Host +
		" user=" + config.User +
		" password=" + config.Password +
		" dbname=" + config.DBName +
		" port=" + config.Port +
		" sslmode=" + config.SSLMode +
		" TimeZone=" + config.TimeZone
}

func runMigrations(db *gorm.DB) error {
	err := db.AutoMigrate(&modules.Note{})
	if err != nil {
		log.Printf("Failod to migrate database: %v", err)
		return err
	}
	return nil
}
