package service

import (
	"errors"
	"learn-golang-unit-test/entity"
	"learn-golang-unit-test/repository"
)

type CategoryService struct {
	Repository repository.CategoryRespository
}

func (service CategoryService) Get(id string) (*entity.Category, error) {
	category := service.Repository.FindById(id)

	if category == nil {
		return nil, errors.New("category not found")
	} else {
		return category, nil
	}
}
