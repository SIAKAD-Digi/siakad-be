package repository

import (
	"siakad-digi/internal/models/database"

	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

func (r *UserRepository) DeleteById(id string) {
	user := database.User{}
	r.DB.Delete(&user, "id = ?", id)
}

func (r *UserRepository) FindById(id string) (*database.User, error) {
	user := database.User{}

	res := r.DB.First(&user, "id = ?", id)

	return &user, res.Error
}

func (r *UserRepository) FindByNik(nik string) (*database.User, error) {
	user := database.User{}

	res := r.DB.Unscoped().First(&user, "nik = ?", nik)

	return &user, res.Error
}

func (r *UserRepository) FindByEmail(email string) (*database.User, error) {
	user := database.User{}

	res := r.DB.Unscoped().First(&user, "email = ?", email)

	return &user, res.Error
}

func (r *UserRepository) FindByPhoneNumber(phoneNumber string) (*database.User, error) {
	user := database.User{}

	res := r.DB.Unscoped().First(&user, "phone_number = ?", phoneNumber)

	return &user, res.Error
}
