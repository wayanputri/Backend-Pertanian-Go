package repositori

import (
	"backend/pb"
	"backend/pkg/models"
	"errors"
	"fmt"
	"log"
	"time"
)

func (r *Provider) InsertLifeStock(req *models.Livestocks) error {

	req.CreatedAt = time.Now()
	req.DeletedAt = nil
	if err := r.dbGorm.Create(&req).Error; err != nil {
		log.Fatal("Gagal menyisipkan pengguna:", err)
		return err
	} else {
		fmt.Println("Insert life stock berhasil ditambahkan:", req.Spesies)
	}

	return nil
}

func (r *Provider) GetLifeStockAll(id uint64, role string) ([]*pb.LifeStockAll, error) {
	query := r.dbGorm

	if role == "user" {
		query = query.Where("user_id=? and deleted_at IS NULL", id)
	} else {
		query = query.Where("deleted_at IS NULL").Preload("User")
	}
	var data []models.Livestocks
	if err := query.Find(&data).Error; err != nil {
		return nil, err
	}

	lifeStockAll := models.LifeStockAll(data)
	return lifeStockAll, nil

}

func (r *Provider) GetLifeStockDetail(id int, role string, reference string) (*pb.LifeStockDetail, error) {
	query := r.dbGorm
	if role == "user" {
		query = query.Where("user_id=?", id)
	}

	var data models.Livestocks
	if err := query.Where("reference=? and deleted_at IS NULL", reference).Preload("FarmArea").First(&data).Error; err != nil {
		return nil, err
	}

	response := models.LifeStockDetail(data)
	return response, nil
}
func (r *Provider) DeleteLifeStock(id int, userId int, role string) error {
	query := r.dbGorm

	var corp models.Livestocks
	if role == "user" {
		query = query.Where("user_id=?", userId)
	}
	now := time.Now()
	if err := query.Model(&corp).Where("id = ?", id).Updates(&models.Livestocks{
		DeletedAt: &now,
	}).Error; err == nil {
		return err
	}
	return nil

}

func (r *Provider) UpdateLifeStock(req models.Livestocks, id int, role string, wide string) error {
	query := r.dbGorm

	var lifeStock models.Livestocks
	if role == "user" {
		query = query.Where("user_id=?", id)
	}
	if err := query.Model(&lifeStock).Where("id = ?", req.ID).Updates(&req).Error; err == nil {

		if err = r.dbGorm.First(&lifeStock, req.ID).Error; err != nil {
			return errors.New("life stock not found")
		}

		if err = r.dbGorm.Model(&models.FarmArea{}).Where("id = ?", lifeStock.FarmAreaID).Updates(&models.FarmArea{
			Wide: wide,
		}).Error; err != nil {
			return err
		}
		return nil
	}
	return errors.New("update failed")

}
