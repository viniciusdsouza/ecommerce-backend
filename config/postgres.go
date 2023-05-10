package config

import (
	"database/sql"
	"os"

	"github.com/viniciusdsouza/ecommerce-backend/schemas"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitializePostgreSQL() (*gorm.DB, error) {
	logger := GetLogger("postgresql")
	dbPath := "./db/main.db"

	// Check if the database file already exists
	_, err := os.Stat(dbPath)
	if os.IsNotExist(err) {
		logger.Info("database file not found, creating...")

		err = os.MkdirAll("./db", os.ModePerm)
		if err != nil {
			return nil, err
		}

		file, err := os.Create(dbPath)
		if err != nil {
			return nil, err
		}

		file.Close()
	}

	// Create and Connet Database
	dsn := "host=localhost user=postgres password=postgres dbname=ecommerce-backend port=5433 sslmode=disable"
	sqlDB, err := sql.Open("pgx", dsn)
	if err != nil {
		logger.Errorf("postgres opening error: %v", err)
		return nil, err
	}
	db, err := gorm.Open(postgres.New(postgres.Config{
		Conn: sqlDB,
	}), &gorm.Config{})
	if err != nil {
		logger.Errorf("postgres opening error: %v", err)
		return nil, err
	}

	// Migrate the Schemas
	err = db.AutoMigrate(&schemas.Address{})
	if err != nil {
		logger.Errorf("postgres automigrate error: %v", err)
		return nil, err
	}

	err = db.AutoMigrate(&schemas.Cart{})
	if err != nil {
		logger.Errorf("postgres automigrate error: %v", err)
		return nil, err
	}

	err = db.AutoMigrate(&schemas.CartItem{})
	if err != nil {
		logger.Errorf("postgres automigrate error: %v", err)
		return nil, err
	}

	err = db.AutoMigrate(&schemas.Customer{})
	if err != nil {
		logger.Errorf("postgres automigrate error: %v", err)
		return nil, err
	}

	err = db.AutoMigrate(&schemas.Order{})
	if err != nil {
		logger.Errorf("postgres automigrate error: %v", err)
		return nil, err
	}

	err = db.AutoMigrate(&schemas.Product{})
	if err != nil {
		logger.Errorf("postgres automigrate error: %v", err)
		return nil, err
	}

	err = db.AutoMigrate(&schemas.Seller{})
	if err != nil {
		logger.Errorf("postgres automigrate error: %v", err)
		return nil, err
	}

	err = db.AutoMigrate(&schemas.UserSession{})
	if err != nil {
		logger.Errorf("postgres automigrate error: %v", err)
		return nil, err
	}

	// Return the Datebase
	return db, nil
}
