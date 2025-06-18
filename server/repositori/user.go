package repositori

import (
	"backend/pkg/models"
	"fmt"
	"log"
)

func (r *Provider) CreateUser(req *models.User) (int, error) {

	if err := r.dbGorm.Create(&req).Error; err != nil {
		log.Fatal("Gagal menyisipkan pengguna:", err)
		return 0, err
	} else {
		fmt.Println("Pengguna berhasil disisipkan:", req.Name)
	}

	return req.ID, nil
}

func (r *Provider) GetUserByEmail(email string) (string, string, int, string, error) {
	fmt.Println("email payload", email)
	var user models.User
	if err := r.dbGorm.Where("email=?", email).First(&user).Error; err != nil {
		fmt.Print("error get user")
		return "", "", 0, "", err
	}
	fmt.Println("email db", user.Email)

	return user.Email, user.Password, user.ID, user.Role, nil
}

func (r *Provider) GetUserAll() ([]models.User, error) {

	var user []models.User
	if err := r.dbGorm.Find(&user).Error; err != nil {
		return nil, err
	}

	return user, nil
}

func (r *Provider) UpdatePassword(password string, email string) error {

	var user models.User
	if err := r.dbGorm.Where("email=?", email).First(&user).Error; err != nil {
		return err
	}

	user.Password = password
	if err := r.dbGorm.Save(&user).Error; err != nil {
		return err
	}

	return nil
}

func (r *Provider) UpdateProfil(id int, req models.User) error {

	var user models.User
	if err := r.dbGorm.Where("id=?", id).First(&user).Error; err != nil {
		return err
	}

	user.Password = req.Password
	user.Email = req.Email

	if err := r.dbGorm.Save(&user).Error; err != nil {
		return err
	}

	return nil
}

func (r *Provider) UpdateRole(id int, roleUser string) error {

	var user models.User
	if err := r.dbGorm.Where("id=?", id).First(&user).Error; err != nil {
		return err
	}

	user.Role = roleUser

	if err := r.dbGorm.Save(&user).Error; err != nil {
		return err
	}

	return nil
}
