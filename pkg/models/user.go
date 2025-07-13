package models

import (
	"backend/pb"
	"time"
)

type User struct {
	ID              int `gorm:"primaryKey"`
	Name            string
	Email           string `gorm:"uniqueIndex"`
	Role            string
	Password        string
	Address         string
	LandArea        string
	TypeOfLivestock string
	Profile         string
	CreatedAt       time.Time
	UpdatedAt       time.Time
	DeletedAt       time.Time

	Cctv       []Cctv       `gorm:"foreignKey:UserID"`
	Action     []Actions    `gorm:"foreignKey:UserID"`
	Livestocks []Livestocks `gorm:"foreignKey:UserID"`
	Crops      []Crops      `gorm:"foreignKey:UserID"`
}

func UserProfile(user User) *pb.User {
	return &pb.User{
		Id:              int64(user.ID),
		Name:            user.Name,
		Address:         user.Address,
		LandArea:        user.LandArea,
		Email:           user.Email,
		TypeOfLivestock: user.TypeOfLivestock,
		Ficture:         user.Profile,
	}
}
