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

type StudentRepository struct {
	DB *gorm.DB
}

func (r *StudentRepository) Create(req *request.CreateStudentRequest, password string) (string, error) {
	user := database.User{
		ID:          uuid.NewString(),
		Name:        req.Name,
		Nik:         req.Nik,
		Email:       req.Email,
		PhoneNumber: req.PhoneNumber,
		Password:    password,
		BirthOfDate: req.BirthOfDate,
		Address:     req.Address,
		RoleId:      constant.STUDENT_ROLE,
		IsActive:    true,
	}
	student := database.Student{
		ID:              uuid.NewString(),
		StudentGuardian: req.StudentGuardian,
	}

	err := r.DB.Transaction(func(tx *gorm.DB) error {

		if err := tx.Create(&user).Error; err != nil {
			return err
		}

		student.UsersID = user.ID

		if err := tx.Create(&student).Error; err != nil {
			return err
		}

		return nil
	})

	return user.ID, err
}

func (r *StudentRepository) Update(req *request.UpdateStudentRequest, id string) (string, error) {
	user := database.User{}
	student := database.Student{}

	userUpdate := map[string]any{
		"name":          req.Name,
		"nik":           req.Nik,
		"email":         req.Email,
		"phone_number":  req.PhoneNumber,
		"birth_of_date": req.BirthOfDate,
		"address":       req.Address,
		"is_active":     *req.IsActive,
	}

	studentUpdate := map[string]any{
		"student_guardian": req.StudentGuardian,
	}

	err := r.DB.Transaction(func(tx *gorm.DB) error {

		userQuery := tx.Model(&user)
		userQuery.Where("id = ?", id)
		userQuery.Select("name", "nik", "email", "phone_number", "birth_of_date", "address", "is_active")

		if err := userQuery.Updates(&userUpdate).Error; err != nil {
			return err
		}

		studentQuery := tx.Model(student)
		studentQuery.Where("users_id = ?", id)
		studentQuery.Select("student_guardian")

		if err := studentQuery.Updates(&studentUpdate).Error; err != nil {
			return err
		}

		return nil
	})

	return user.ID, err
}

func (r *StudentRepository) FindAll(req *request.FindAllStudentRequest) (*[]api.StudentApi, int) {
	students := []api.StudentApi{}
	user := database.User{}
	total := int64(0)

	query := r.DB.Model(&user).Select("users.id, users.name, users.nik, classes.name as class_name, users.is_active, users.created_at")
	query.Joins("left join students on students.users_id = users.id").Joins("left join classes on classes.id = students.classes_id")

	if req.Name != "" {
		query.Where("users.name like ?", req.Name+"%")
	}

	if req.StartDate != "" && req.EndDate != "" {
		query.Where("DATE(users.created_at) between (?::date - INTERVAL '1 day') and ?", req.StartDate, req.EndDate)
	}

	if req.Status == "active" {
		query.Where("users.is_active", true)
	}

	if req.Status == "not_active" {
		query.Where("users.is_active", false)
	}

	query.Where("users.role_id = ?", constant.STUDENT_ROLE)
	query.Count(&total)
	query.Scopes(utils.Paginate(req.Page, req.Limit)).Scan(&students)

	return &students, int(total)

}

func (r *StudentRepository) FindById(id string) api.StudentDetailApi {
	user := database.User{}
	student := api.StudentDetailApi{}

	query := r.DB.Model(user)
	query.Select("users.id, users.name , users.nik, users.phone_number, users.email, classes.name as class_name,  users.profile_picture, users.birth_of_date, users.address, users.is_active, students.student_guardian, users.created_at, users.updated_at")
	query.Joins("left join students on students.users_id = users.id")
	query.Joins("left join classes on classes.id = students.classes_id")
	query.Where("users.id = ?", id)
	query.Scan(&student)

	return student
}
