package repositori

import (
	"backend/pb"
	"backend/pkg/models"
	"fmt"
	"log"
)

func (r *Provider) InsertCrops(req *models.Crops) error {

	if err := r.dbGorm.Create(&req).Error; err != nil {
		log.Fatal("Gagal menyisipkan pengguna:", err)
		return err
	} else {
		fmt.Println("Insert crops berhasil ditambahkan:", req.Spesies)
	}

	return nil
}

func (r *Provider) InsertArea(req *models.FarmArea) (int, error) {

	if err := r.dbGorm.Create(&req).Error; err != nil {
		log.Fatal("Gagal menyisipkan pengguna:", err)
		return 0, err
	} else {
		fmt.Println("Insert farm area berhasil ditambahkan:", req.ID)
	}

	return req.ID, nil
}

func (r *Provider) GetCropsAll(id uint64, role string) ([]*pb.CropsAll, error) {
	query := r.dbGorm

	if role == "user" {
		query = query.Where("user_id=?", id)
	} else {
		query = query.Preload("User")
	}
	var data []models.Crops
	if err := query.Find(&data).Error; err != nil {
		return nil, err
	}

	cropAll := models.CropAll(data)
	return cropAll, nil

}

func (r *Provider) GetCropDetail(id int, role string, reference string) (*pb.CropDetail, error) {
	query := r.dbGorm
	if role == "user" {
		query = query.Where("user_id=?", id)
	}

	var data models.Crops
	if err := query.Where("reference=?", reference).Preload("FarmArea").First(&data).Error; err != nil {
		return nil, err
	}

	response := models.CropDetail(data)
	return response, nil
}
