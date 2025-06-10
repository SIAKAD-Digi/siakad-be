package repository

import (
	"siakad-digi/internal/models/api"
	"siakad-digi/internal/models/database"
	"siakad-digi/internal/models/request"
	"siakad-digi/utils"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ClassRepository struct {
	DB *gorm.DB
}

func (r *ClassRepository) Create(req *request.CreateOrUpdateClassRequest) (string, error) {
	class := database.Class{
		ID:   uuid.NewString(),
		Name: req.Name,
	}

	err := r.DB.Create(&class).Error

	return class.ID, err
}

func (r *ClassRepository) FindById(id string) (*database.Class, error) {
	class := database.Class{}

	err := r.DB.First(&class, "id = ?", id).Error

	return &class, err
}

func (r *ClassRepository) FindByName(name string) (*database.Class, error) {
	class := database.Class{}

	err := r.DB.Unscoped().First(&class, "name = ?", name).Error

	return &class, err
}

func (r *ClassRepository) Update(req *request.CreateOrUpdateClassRequest, id string) (string, error) {
	class := database.Class{}

	err := r.DB.Model(&class).Where("id = ?", id).Update("name", req.Name).Error

	return class.ID, err
}

func (r *ClassRepository) FindAll(req *request.FindAllClassRequest) (*[]api.ClassAPi, int) {
	classes := []api.ClassAPi{}
	class := database.Class{}
	total := int64(0)

	query := r.DB.Model(&class)

	if req.Name != "" {
		query.Where("name like ?", req.Name+"%")
	}

	if req.StartDate != "" && req.EndDate != "" {
		query.Where("DATE(created_at) between (?::date - INTERVAL '1 day') and ?", req.StartDate, req.EndDate)
	}

	query.Order("created_at DESC")
	query.Count(&total)
	query.Scopes(utils.Paginate(req.Page, req.Limit)).Find(&classes)

	return &classes, int(total)

}

func (r *ClassRepository) DeleteById(id string) {
	class := database.Class{}
	r.DB.Delete(&class, "id = ?", id)
}
