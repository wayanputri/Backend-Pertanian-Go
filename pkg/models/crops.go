package models

import "backend/pb"

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

	FarmAreaID int
	UserID     int
	FarmArea   FarmArea  `gorm:"foreignKey:FarmAreaID"`
	User       User      `gorm:"foreignKey:UserID"`
	Actions    []Actions `gorm:"foreignKey:TargetID"`
}

func Crop(crop Crops) *pb.CropsAll {
	return &pb.CropsAll{
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
