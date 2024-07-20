package db

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

type Config struct {
	Client   string `required:"true"`
	Database string `required:"true"`
	Username string `required:"true"`
	Password string `required:"true"`
	Host     string `required:"true"`
	Port     string `required:"true"`
}

func InitDB(config Config) (*gorm.DB, error) {
	var db *gorm.DB
	var err error

	switch config.Client {
	case "postgresql":
		dsn := fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%s", config.Username, config.Password, config.Database, config.Host, config.Port)
		db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	default:
		return nil, fmt.Errorf("unsupported database client: %s", config.Client)
	}

	if err != nil {
		return nil, err
	}

	DB = db
	return db, nil
}

func CloseDB() {
	sqlDB, _ := DB.DB()
	sqlDB.Close()
}
