package repositori

import (
	"backend/pb"
	"backend/pkg/models"
	"errors"
	"fmt"
	"log"
	"time"
)

func (r *Provider) InsertCrops(req *models.Crops) error {

	req.CreatedAt = time.Now()
	req.DeletedAt = nil
	if err := r.dbGorm.Create(&req).Error; err != nil {
		log.Fatal("Gagal menyisipkan pengguna:", err)
		return err
	} else {
		fmt.Println("Insert crops berhasil ditambahkan:", req.Spesies)
	}

	return nil
}

func (r *Provider) UpdateCrops(req models.Crops, id int, role string, wide string) error {
	query := r.dbGorm

	var corp models.Crops
	if role == "user" {
		query = query.Where("user_id=?", id)
	}
	if err := query.Model(&corp).Where("id = ?", req.ID).Updates(&req).Error; err == nil {

		if err = r.dbGorm.First(&corp, req.ID).Error; err != nil {
			return errors.New("corp not found")
		}
		fmt.Println("cropsjjj", corp)
		fmt.Println("corp.FarmAreaID", corp.FarmAreaID)
		if err = r.dbGorm.Model(&models.FarmArea{}).Where("id = ?", corp.FarmAreaID).Updates(&models.FarmArea{
			Wide: wide,
		}).Error; err != nil {
			return err
		}
		return nil
	}
	return errors.New("update failed")

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
		query = query.Where("user_id=? and deleted_at IS NULL", id)
	} else {
		query = query.Where("deleted_at IS NULL").Preload("User")
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
	if err := query.Where("reference=? and deleted_at IS NULL", reference).Preload("FarmArea").First(&data).Error; err != nil {
		return nil, err
	}

	response := models.CropDetail(data)
	return response, nil
}

func (r *Provider) DeleteCrops(id int, userId int, role string) error {
	query := r.dbGorm

	var corp models.Crops
	if role == "user" {
		query = query.Where("user_id=?", userId)
	}
	now := time.Now()
	if err := query.Model(&corp).Where("id = ?", id).Updates(&models.Crops{
		DeletedAt: &now,
	}).Error; err == nil {
		return err
	}
	return nil

}
