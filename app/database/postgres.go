package database

import (
	"backend/app/config"
	"backend/pkg/models"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ProstgresConfig() *sql.DB {

	config.LoadConfig()

	cfg := config.AppConfig
	connStr := fmt.Sprintf("user=%s dbname=%s password=%s host=%s port=%s sslmode=disable", cfg.PostgresUser, cfg.PostgresName, cfg.PostgresPassword, cfg.PostgresHost, cfg.PostgresPort)
	// Menghubungkan ke PostgreSQL
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Error membuka koneksi ke database: ", err)
	}

	// Mengecek koneksi ke database
	err = db.Ping()
	if err != nil {
		log.Fatal("Gagal menghubungkan ke database: ", err)
	}

	fmt.Println("Berhasil terhubung ke database PostgreSQL!")
	return db
}

func InitPostgresGorm() *gorm.DB {
	config.LoadConfig()
	cfg := config.AppConfig

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		cfg.PostgresHost,
		cfg.PostgresUser,
		cfg.PostgresPassword,
		cfg.PostgresName,
		cfg.PostgresPort,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to Postgres: %v", err)
	}

	err = db.AutoMigrate(&models.User{}, &models.Cctv{}, &models.Actions{}, &models.FarmArea{}, &models.Crops{}, &models.Livestocks{}, &models.Sensor{}, &models.SensorReading{})
	if err != nil {
		log.Fatalf("AutoMigrate error: %v", err)
	}

	return db
}
