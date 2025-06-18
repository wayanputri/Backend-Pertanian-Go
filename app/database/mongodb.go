package database

import (
	"backend/app/config"
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func InitMongo(ctx context.Context) (*mongo.Client, *mongo.Collection) {
	// Inisialisasi koneksi ke MongoDB
	config.LoadConfig()
	cfg := config.AppConfig
	clientOptions := options.Client().ApplyURI(cfg.MongoURI)
	var err error
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatalf("Gagal terkoneksi ke MongoDB: %v", err)
	}

	// Verifikasi koneksi
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatalf("Koneksi gagal: %v", err)
	} else {
		fmt.Println("Koneksi MongoDB berhasil!")
	}

	// Pilih database dan koleksi
	database := client.Database(cfg.MongoDB)               // Nama database
	collection := database.Collection(cfg.MongoCollection) // Nama koleksi
	return client, collection
}
