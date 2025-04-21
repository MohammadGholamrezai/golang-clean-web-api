package db

import (
	"fmt"
	"log"
	"time"

	"github.com/MohammadGholamrezai/golang-clean-web-api/config"
	"github.com/MohammadGholamrezai/golang-clean-web-api/pkg/logging"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var dbClient *gorm.DB
var logger = logging.NewLogger(config.GetConfig())

func InitDb(cfg *config.Config) error {
	cnn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=Asia/Tehran",
		cfg.Postgres.Host,
		cfg.Postgres.User,
		cfg.Postgres.Password,
		cfg.Postgres.DbName,
		cfg.Postgres.Port,
		cfg.Postgres.SSLMode)

	dbClient, err := gorm.Open(postgres.Open(cnn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	sqlDb, _ := dbClient.DB()
	if err = sqlDb.Ping(); err != nil {
		panic(err)
	}

	sqlDb.SetMaxIdleConns(cfg.Postgres.MaxIdleConns)
	sqlDb.SetMaxOpenConns(cfg.Postgres.MaxOpenConns)
	sqlDb.SetConnMaxLifetime(time.Duration(cfg.Postgres.ConnMaxLifetime) * time.Minute)

	logger.Info(logging.Postgres, logging.Startup, "Database connection initialized successfully.", nil)
	// log.Println("Database connection initialized successfully.")
	return nil
}

func GetDb() *gorm.DB {
	return dbClient
}

func CloseDb() {
	sqlDb, err := dbClient.DB()
	if err != nil {
		log.Println("Error getting database instance:", err)
		return
	}
	sqlDb.Close()
	log.Println("Database connection closed successfully.")
}
