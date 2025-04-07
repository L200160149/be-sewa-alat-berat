package app

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/L200160149/be-sewa-alat-berat/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewDB() *gorm.DB {
	// Load environment variables
	config.InitEnv()

	// db connections
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	// db pool configs
	dbMaxIdleConnStr := os.Getenv("DB_MAX_IDLE_CONNECTIONS")
	dbMaxOpenConnStr := os.Getenv("DB_MAX_OPEN_CONNECTIONS")
	dbMaxConnLifetimeStr := os.Getenv("DB_MAX_CONN_LIFETIME")
	dbMaxIdleTimeStr := os.Getenv("DB_MAX_IDLE_TIME")

	dbMaxIdleConn, _ := strconv.Atoi(dbMaxIdleConnStr)
	dbMaxOpenConn, _ := strconv.Atoi(dbMaxOpenConnStr)
	dbMaxConnLifetime, _ := strconv.Atoi(dbMaxConnLifetimeStr)
	dbMaxIdleTime, _ := strconv.Atoi(dbMaxIdleTimeStr)

	if dbUser == "" || dbPassword == "" || dbHost == "" || dbPort == "" || dbName == "" {
		log.Fatal("Missing required database environment variables")
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbUser, dbPassword, dbHost, dbPort, dbName)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal("Failed to get sql.DB from gorm.DB:", err)
	}

	sqlDB.SetMaxIdleConns(dbMaxIdleConn)
	sqlDB.SetMaxOpenConns(dbMaxOpenConn)
	sqlDB.SetConnMaxLifetime(time.Duration(dbMaxConnLifetime) * time.Minute)
	sqlDB.SetConnMaxIdleTime(time.Duration(dbMaxIdleTime) * time.Minute)

	return db
}