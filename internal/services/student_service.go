package services

import (
	"siakad-digi/internal/exception"
	"siakad-digi/internal/models/api"
	"siakad-digi/internal/models/request"
	"siakad-digi/internal/repository"
	"siakad-digi/utils"
)

type StudentService struct {
	UserRepository    *repository.UserRepository
	StudentRepository *repository.StudentRepository
}

func (s *StudentService) Create(req *request.CreateStudentRequest) {
	e, _ := s.UserRepository.FindByEmail(req.Email)

	if e.Email != "" {
		panic(exception.NewBadRequestError("Email sudah terdaftar"))
	}

	n, _ := s.UserRepository.FindByNik(req.Nik)

	if n.Nik != "" {
		panic(exception.NewBadRequestError("Nik sudah terdaftar"))
	}

	p, _ := s.UserRepository.FindByPhoneNumber(req.PhoneNumber)

	if p.PhoneNumber != "" {
		panic(exception.NewBadRequestError("Nomer telepon sudah terdaftar"))
	}

	password := utils.HashPassword(req.Nik)
	_, err := s.StudentRepository.Create(req, password)

	if err != nil {
		panic(err)
	}
}

func (s *StudentService) Update(req *request.UpdateStudentRequest, id string) {
	user, _ := s.UserRepository.FindById(id)

	if user.ID == "" {
		panic(exception.NewNotFoundError("Siswa tidak di temukan"))
	}

	e, _ := s.UserRepository.FindByEmail(req.Email)

	if e.Email != "" && e.ID != id {
		panic(exception.NewBadRequestError("Email sudah terdaftar"))
	}

	n, _ := s.UserRepository.FindByNik(req.Nik)

	if n.Nik != "" && n.ID != id {
		panic(exception.NewBadRequestError("Nik sudah terdaftar"))
	}

	p, _ := s.UserRepository.FindByPhoneNumber(req.PhoneNumber)

	if p.PhoneNumber != "" && p.ID != id {
		panic(exception.NewBadRequestError("Nomer telepon sudah terdaftar"))
	}

	_, err := s.StudentRepository.Update(req, id)

	if err != nil {
		panic(err)
	}
}

func (s *StudentService) FindAll(req *request.FindAllStudentRequest) (*[]api.StudentApi, int) {
	students, total := s.StudentRepository.FindAll(req)

	return students, total
}

func (s *StudentService) FindById(id string) api.StudentDetailApi {
	user, _ := s.UserRepository.FindById(id)

	if user.ID == "" {
		panic(exception.NewNotFoundError("Siswa tidak di temukan"))
	}

	student := s.StudentRepository.FindById(id)

	return student
}

func (s *StudentService) DeleteById(id string) {
	user, _ := s.UserRepository.FindById(id)

	if user.ID == "" {
		panic(exception.NewNotFoundError("Siswa tidak di temukan"))
	}

	s.UserRepository.DeleteById(id)
}
