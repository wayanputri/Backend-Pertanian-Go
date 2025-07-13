package models

import (
	"backend/pb"
	"time"
)

type Crops struct {
	ID                 int `gorm:"primaryKey"`
	Spesies            string
	Age                string
	GrowStage          string
	HealthStatus       string
	WaterNeeds         string
	PestControl        string
	OrganicFertilizer  string
	ChemicalFertilizer string
	Reference          string
	CreatedAt          time.Time
	UpdatedAt          time.Time
	DeletedAt          *time.Time `gorm:"index"`

	FarmAreaID int
	UserID     int
	FarmArea   FarmArea  `gorm:"foreignKey:FarmAreaID"`
	User       User      `gorm:"foreignKey:UserID"`
	Actions    []Actions `gorm:"foreignKey:TargetID"`
}

func Crop(crop Crops) *pb.CropsAll {
	return &pb.CropsAll{
		Id:            int64(crop.ID),
		Reference:     crop.Reference,
		JenisTumbuhan: crop.Spesies,
		Umur:          crop.Age,
		NamaUser:      crop.User.Name,
	}
}
func CropAll(crop []Crops) []*pb.CropsAll {
	var cropAll []*pb.CropsAll
	for _, v := range crop {
		cropAll = append(cropAll, Crop(v))
	}
	return cropAll
}

func CropDetail(crop Crops) *pb.CropDetail {
	return &pb.CropDetail{
		Id:                        int64(crop.ID),
		Reference:                 crop.Reference,
		Jenis:                     crop.Spesies,
		Umur:                      crop.Age,
		KebutuhanPupukKimia:       crop.ChemicalFertilizer,
		KebutuhanPupukOrganik:     crop.OrganicFertilizer,
		LuasLahan:                 crop.FarmArea.Wide,
		KebutuhanPengendalianHama: crop.PestControl,
		KebutuhanAir:              crop.WaterNeeds,
	}
}

func CropUpdate(crop *Crops) Crops {
	return Crops{
		ID:                 crop.ID,
		Spesies:            crop.Spesies,
		Age:                crop.Age,
		GrowStage:          crop.GrowStage,
		HealthStatus:       crop.HealthStatus,
		WaterNeeds:         crop.WaterNeeds,
		PestControl:        crop.PestControl,
		OrganicFertilizer:  crop.OrganicFertilizer,
		ChemicalFertilizer: crop.ChemicalFertilizer,
		Reference:          crop.Reference,
		UpdatedAt:          time.Now(),
		FarmAreaID:         crop.FarmAreaID,
		UserID:             crop.UserID,
	}
}
