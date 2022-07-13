package db

import (
	"fmt"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewDB() (*gorm.DB, error) {
	dbUser := os.Getenv("MYSQL_USER")
	if dbUser == "" {
		return nil, fmt.Errorf("empty env MYSQL_USER")
	}
	dbPassword := os.Getenv("MYSQL_USER_PASSWORD")
	if dbPassword == "" {
		return nil, fmt.Errorf("empty env MYSQL_USER_PASSWORD")
	}
	dbSrc := os.Getenv("MYSQL_HOST")
	if dbSrc == "" {
		return nil, fmt.Errorf("empty env MYSQL_HOST")
	}
	dbName := os.Getenv("MYSQL_DATABASE")
	if dbName == "" {
		return nil, fmt.Errorf("empty env MYSQL_DATABASE")
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbUser, dbPassword, dbSrc, dbName)
	engine, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	db, err := engine.DB()
	if err != nil {
		return nil, err
	}

	db.SetConnMaxLifetime(time.Duration(10) * time.Second)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	return engine, nil
}
