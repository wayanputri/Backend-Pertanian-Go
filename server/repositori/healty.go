package repositori

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/go-redis/redis/v8"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func (r *Provider) GetUsers() {

	// Query untuk mengambil data
	rows, err := r.DBSql().Query("SELECT id, name FROM users")
	if err != nil {
		log.Fatal("Gagal mengeksekusi query: ", err)
	}
	defer rows.Close()

	// Loop melalui hasil query
	for rows.Next() {
		var id int
		var name string
		if err := rows.Scan(&id, &name); err != nil {
			log.Fatal("Gagal memindahkan data: ", err)
		}
		fmt.Printf("ID: %d, Name: %s\n", id, name)
	}

	// Cek jika ada error setelah loop
	if err := rows.Err(); err != nil {
		log.Fatal("Error setelah looping rows: ", err)
	}
}

func (r *Provider) SetData(ctx context.Context, myKey string, value string) {
	// Menyimpan data ke Redis
	err := r.DBRedis().Set(ctx, myKey, value, 10*time.Minute).Err()
	if err != nil {
		log.Fatalf("Gagal menyimpan data: %v", err)
	}
	fmt.Println("Data berhasil disimpan!")
}

func (r *Provider) GetData(ctx context.Context, myKey string) string {
	// Mengambil data dari Redis
	val, err := r.DBRedis().Get(ctx, myKey).Result()
	if err == redis.Nil {
		fmt.Println("Key tidak ditemukan!")
	} else if err != nil {
		log.Fatalf("Gagal mengambil data: %v", err)
	} else {
		fmt.Printf("Data dari Redis: %s\n", val)
	}
	return val
}

func (r *Provider) CreateUserMongo(ctx context.Context) {
	// Menambah data user ke koleksi MongoDB
	user := User{Name: "John Doe4", Email: "johndoe4@example.com"}
	insertResult, err := r.DBMongo().InsertOne(ctx, user)
	if err != nil {
		log.Fatalf("Gagal menambah data user: %v", err)
	}
	fmt.Printf("User berhasil ditambahkan dengan ID mongo: %v\n", insertResult.InsertedID)
}

func (r *Provider) GetUserMongo(ctx context.Context) {
	// Mengambil satu data user dari koleksi MongoDB
	var result User
	err := r.DBMongo().FindOne(ctx, bson.D{{Key: "name", Value: "John Doe4"}}).Decode(&result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			fmt.Println("Data tidak ditemukan")
		} else {
			log.Fatalf("Gagal mengambil data user mongo: %v", err)
		}
	} else {
		fmt.Printf("User ditemukan: %v\n", result)
	}
}

type User struct {
	Name  string `bson:"name"`
	Email string `bson:"email"`
}
