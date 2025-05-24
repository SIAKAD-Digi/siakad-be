package services

import (
	"siakad-digi/internal/exception"
	"siakad-digi/internal/models/api"
	"siakad-digi/internal/models/request"
	"siakad-digi/internal/repository"
	"siakad-digi/utils"
)

type TeacherService struct {
	UserRepository    *repository.UserRepository
	TeacherRepository *repository.TeacherRepository
}

func (s *TeacherService) Create(req *request.CreateTeacherRequest) {
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
	_, err := s.TeacherRepository.Create(req, password)

	if err != nil {
		panic(err)
	}
}

func (s *TeacherService) Update(req *request.UpdateTeacherRequest, id string) {
	user, _ := s.UserRepository.FindById(id)

	if user.ID == "" {
		panic(exception.NewNotFoundError("Guru tidak di temukan"))
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

	_, err := s.TeacherRepository.Update(req, id)

	if err != nil {
		panic(err)
	}
}

func (s *TeacherService) FindAll(req *request.FindAllTeacherRequest) (*[]api.TeacherApi, int) {
	teachers, total := s.TeacherRepository.FindAll(req)

	return teachers, total
}

func (s *TeacherService) FindById(id string) api.TeacherDetailApi {
	user, _ := s.UserRepository.FindById(id)

	if user.ID == "" {
		panic(exception.NewNotFoundError("Guru tidak di temukan"))
	}

	teacher := s.TeacherRepository.FindById(id)

	return teacher
}

func (s *TeacherService) DeleteById(id string) {
	user, _ := s.UserRepository.FindById(id)

	if user.ID == "" {
		panic(exception.NewNotFoundError("Guru tidak di temukan"))
	}

	s.UserRepository.DeleteById(id)
}
