package database

import (
	"backend/app/config"
	"context"
	"fmt"
	"log"

	"github.com/go-redis/redis/v8"
)

// var rdb *redis.Client
// var ctx = context.Background()

func InitRedis(ctx context.Context) *redis.Client {
	// Mengonfigurasi client Redis
	config.LoadConfig()
	cfg := config.AppConfig
	rdb := redis.NewClient(&redis.Options{
		Addr:     cfg.RedisAddress,  // Alamat Redis server
		Password: cfg.RedisPassword, // Tidak ada password (jika ada, masukkan di sini)
		DB:       cfg.RedisDB,       // Pilih database yang digunakan (default: 0)
	})

	// Cek koneksi ke Redis
	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		log.Fatalf("Koneksi Redis gagal: %v", err)
	} else {
		fmt.Println("Koneksi Redis berhasil!")
	}
	return rdb
}
