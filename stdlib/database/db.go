package database

import (
	"fmt"

	log "github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDB(logger *log.Logger, option Option) *gorm.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", option.User, option.Password, option.Host, option.Port, option.DatabaseName)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		logger.Fatal("failed to connect to database:", err)
	}

	// Test the connection
	sqlDB, err := db.DB()
	if err != nil {
		logger.Fatal("failed to get database object:", err)
	}
	defer sqlDB.Close()

	return db
}

type Option struct {
	User         string `env:"DB_USER"`
	Password     string `env:"DB_PASSWORD"`
	Host         string `env:"DB_HOST"`
	Port         string `env:"DB_PORT"`
	DatabaseName string `env:"DB_NAME"`
}
