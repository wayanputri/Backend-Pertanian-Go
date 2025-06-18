package repositori

import (
	"database/sql"

	"github.com/go-redis/redis/v8"
	"go.mongodb.org/mongo-driver/mongo"
	"gorm.io/gorm"
)

type Provider struct {
	db     *sql.DB
	dbr    *redis.Client
	dbm    *mongo.Collection
	dbGorm *gorm.DB
}

func New(db *sql.DB, dbr *redis.Client, dbm *mongo.Collection, dbGorm *gorm.DB) *Provider {

	return &Provider{
		db:     db,
		dbr:    dbr,
		dbm:    dbm,
		dbGorm: dbGorm,
	}
}

func (p *Provider) DBSql() *sql.DB {
	return p.db
}

func (p *Provider) DBRedis() *redis.Client {
	return p.dbr
}
func (p *Provider) DBMongo() *mongo.Collection {
	return p.dbm
}

func (p *Provider) DBGorm() *gorm.DB {
	return p.dbGorm
}
