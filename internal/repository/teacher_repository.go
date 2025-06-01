package repository

import (
	"siakad-digi/internal/constant"
	"siakad-digi/internal/models/api"
	"siakad-digi/internal/models/database"
	"siakad-digi/internal/models/request"
	"siakad-digi/utils"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type TeacherRepository struct {
	DB *gorm.DB
}

func (r *TeacherRepository) Create(req *request.CreateTeacherRequest, password string) (string, error) {
	user := database.User{
		ID:          uuid.NewString(),
		Name:        req.Name,
		Nik:         req.Nik,
		Email:       req.Email,
		PhoneNumber: req.PhoneNumber,
		Password:    password,
		BirthOfDate: req.BirthOfDate,
		Gender:      req.Gender,
		Address:     req.Address,
		RoleId:      constant.TEACHER_ROLE,
		IsActive:    true,
	}
	teacher := database.Teacher{
		ID:        uuid.NewString(),
		IsMarried: *req.IsMarried,
	}

	err := r.DB.Transaction(func(tx *gorm.DB) error {

		if err := tx.Create(&user).Error; err != nil {
			return err
		}

		teacher.UsersID = user.ID

		if err := tx.Create(&teacher).Error; err != nil {
			return err
		}

		return nil
	})

	return user.ID, err
}

func (r *TeacherRepository) Update(req *request.UpdateTeacherRequest, id string) (string, error) {
	user := database.User{}
	teacher := database.Teacher{}

	userUpdate := map[string]any{
		"name":          req.Name,
		"nik":           req.Nik,
		"email":         req.Email,
		"phone_number":  req.PhoneNumber,
		"birth_of_date": req.BirthOfDate,
		"gender":        req.Gender,
		"address":       req.Address,
		"is_active":     *req.IsActive,
	}

	teacherUpdate := map[string]any{
		"is_married": *req.IsMarried,
	}

	err := r.DB.Transaction(func(tx *gorm.DB) error {

		userQuery := tx.Model(&user)
		userQuery.Where("id = ?", id)
		userQuery.Select("name", "nik", "email", "phone_number", "birth_of_date", "address", "is_active")

		if err := userQuery.Updates(&userUpdate).Error; err != nil {
			return err
		}

		teacherQuery := tx.Model(&teacher)
		teacherQuery.Where("users_id = ?", id)
		teacherQuery.Select("is_married")

		if err := teacherQuery.Updates(&teacherUpdate).Error; err != nil {
			return err
		}

		return nil
	})

	return user.ID, err
}

func (r *TeacherRepository) FindAll(req *request.FindAllTeacherRequest) (*[]api.TeacherApi, int) {
	teachers := []api.TeacherApi{}
	user := database.User{}
	total := int64(0)

	query := r.DB.Model(&user).Select("id, name, nik, email, is_active, created_at")

	if req.Name != "" {
		query.Where("users.name like ?", req.Name+"%")
	}

	if req.StartDate != "" && req.EndDate != "" {
		query.Where("DATE(created_at) between (?::date - INTERVAL '1 day') and ?", req.StartDate, req.EndDate)
	}

	if req.Status == "active" {
		query.Where("is_active", true)
	}

	if req.Status == "not_active" {
		query.Where("is_active", false)
	}

	query.Where("users.role_id = ?", constant.TEACHER_ROLE)
	query.Count(&total)
	query.Order("users.created_at DESC")
	query.Scopes(utils.Paginate(req.Page, req.Limit)).Scan(&teachers)

	return &teachers, int(total)

}

func (r *TeacherRepository) FindById(id string) api.TeacherDetailApi {
	user := database.User{}
	teacher := api.TeacherDetailApi{}

	query := r.DB.Model(user)
	query.Select("users.id, users.name , users.nik, users.phone_number, users.email,  users.profile_picture, users.birth_of_date, users.gender, users.address, users.is_active, teachers.is_married, users.created_at, users.updated_at")
	query.Joins("left join teachers on teachers.users_id = users.id")
	query.Where("users.id = ?", id)
	query.Scan(&teacher)

	return teacher
}
