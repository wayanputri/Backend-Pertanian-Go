package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	ServerPort string
	GRPCPort   string
	GRPCHost   string
	JWTSecret  string

	RedisAddress  string
	RedisPassword string
	RedisDB       int

	PostgresHost     string
	PostgresUser     string
	PostgresName     string
	PostgresPort     string
	PostgresPassword string

	MongoURI        string
	MongoCollection string
	MongoDB         string

	MinioEndpoint  string
	MinioAccessKey string
	MinioSecretKey string
	MinioSecure    bool
	MinioBuckect   string
}

var AppConfig *Config

// LoadConfig menginisialisasi dan memuat konfigurasi dari file .env
func LoadConfig() {
	// Memuat file .env (jika ada)
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	redisDB, err := strconv.Atoi(getEnv("REDIS_DB", "0"))
	if err != nil {
		log.Fatalln("failed convert redis db", err)
		return
	}
	AppConfig = &Config{
		ServerPort: getEnv("SERVER_PORT", ""),
		GRPCPort:   getEnv("GRPC_PORT", ""),
		GRPCHost:   getEnv("GRPC_HOST", ""),
		JWTSecret:  getEnv("JWT_SECRET", ""),

		RedisAddress:  getEnv("REDIS_ADDRESS", ""),
		RedisPassword: getEnv("REDIS_PASSWORD", ""),
		RedisDB:       redisDB,

		PostgresHost:     getEnv("POSTGRES_HOST", ""),
		PostgresUser:     getEnv("POSTGRES_USER", ""),
		PostgresName:     getEnv("POSTGRES_NAME", ""),
		PostgresPort:     getEnv("POSTGRES_PORT", ""),
		PostgresPassword: getEnv("POSTGRES_PASSWORD", ""),

		MongoURI:        getEnv("MONGO_URI", ""),
		MongoCollection: getEnv("MONGO_COLLECTION", ""),
		MongoDB:         getEnv("MONGO_DB", ""),
		MinioEndpoint:   getEnv("MINIO_ENDPOINT", ""),
		MinioAccessKey:  getEnv("MINIO_ACCESKEY", ""),
		MinioSecretKey:  getEnv("MINIO_SECRETKEY", ""),
		MinioBuckect:    getEnv("MINIO_BUCKET", ""),
	}
}

// getEnv untuk mengambil nilai dari environment variable, dengan nilai default
func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}
