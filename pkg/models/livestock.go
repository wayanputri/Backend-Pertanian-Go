package models

import (
	"backend/pb"
	"time"
)

type Livestocks struct {
	ID             int `gorm:"primaryKey"`
	Spesies        string
	Age            string
	HealthStatus   string
	KebutuhanMakan string
	TypeOfFood     string
	Quantity       int
	FarmAreaID     int
	UserID         int
	Reference      string
	CreatedAt      time.Time
	UpdatedAt      time.Time
	DeletedAt      *time.Time
	Actions        []Actions `gorm:"foreignKey:TargetID"`
	FarmArea       FarmArea  `gorm:"foreignKey:FarmAreaID"`
	User           User      `gorm:"foreignKey:UserID"`
}

func LifeStockAll(lifeStock []Livestocks) []*pb.LifeStockAll {
	var lifeStockAll []*pb.LifeStockAll
	for _, v := range lifeStock {
		lifeStockAll = append(lifeStockAll, LifeStock(v))
	}
	return lifeStockAll
}

func LifeStock(lifeStock Livestocks) *pb.LifeStockAll {
	return &pb.LifeStockAll{
		Reference:  lifeStock.Reference,
		JenisHewan: lifeStock.Spesies,
		UmurHewan:  lifeStock.Age,
		NamaUser:   lifeStock.User.Name,
		Id:         int64(lifeStock.ID),
	}
}

func LifeStockDetail(lifeStock Livestocks) *pb.LifeStockDetail {
	return &pb.LifeStockDetail{
		Id:                  int64(lifeStock.ID),
		Reference:           lifeStock.Reference,
		NamaHewan:           lifeStock.Spesies,
		UmurHewan:           lifeStock.Age,
		Quantity:            int64(lifeStock.Quantity),
		LuasKandang:         lifeStock.FarmArea.Wide,
		KebutuhanMakan:      lifeStock.KebutuhanMakan,
		PersentaseKesehatan: lifeStock.HealthStatus,
		JenisMakanan:        lifeStock.TypeOfFood,
	}
}
