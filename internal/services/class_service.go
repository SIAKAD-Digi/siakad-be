package services

import (
	"siakad-digi/internal/exception"
	"siakad-digi/internal/models/api"
	"siakad-digi/internal/models/request"
	"siakad-digi/internal/repository"
)

type ClassService struct {
	ClassRepository *repository.ClassRepository
}

func (s *ClassService) Create(req *request.CreateOrUpdateClassRequest) {

	class, _ := s.ClassRepository.FindByName(req.Name)

	if class.Name != "" {
		panic(exception.NewBadRequestError("Kelas sudah ada"))
	}

	s.ClassRepository.Create(req)
}

func (s *ClassService) Update(req *request.CreateOrUpdateClassRequest, id string) {

	if c, _ := s.ClassRepository.FindById(id); c.ID == "" {
		panic(exception.NewNotFoundError("Kelas tidak di temukan"))
	}

	class, _ := s.ClassRepository.FindByName(req.Name)

	if class.Name != "" && class.ID != id {
		panic(exception.NewBadRequestError("Kelas sudah ada"))
	}

	s.ClassRepository.Update(req, id)
}

func (s *ClassService) FindAll(req *request.FindAllClassRequest) (*[]api.ClassApi, int) {
	classes, total := s.ClassRepository.FindAll(req)

	return classes, total
}

func (s *ClassService) DeleteById(id string) {
	class, _ := s.ClassRepository.FindById(id)

	if class.ID == "" {
		panic(exception.NewNotFoundError("Kelas tidak di temukan"))
	}

	s.ClassRepository.DeleteById(id)
}

func (s *ClassService) FindById(id string) *api.ClassApi {
	class, _ := s.ClassRepository.FindById(id)

	if class.ID == "" {
		panic(exception.NewNotFoundError("Kelas tidak di temukan"))
	}

	classApi, _ := s.ClassRepository.FindById(id)

	return classApi
}
