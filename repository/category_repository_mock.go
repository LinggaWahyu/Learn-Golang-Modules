package repository

import (
	"learn-golang-unit-test/entity"

	"github.com/stretchr/testify/mock"
)

type CategoryRespositoryMock struct {
	Mock mock.Mock
}

func (repository *CategoryRespositoryMock) FindById(id string) *entity.Category {
	arguments := repository.Mock.Called(id)
	if arguments.Get(0) == nil {
		return nil
	} else {
		category := arguments.Get(0).(entity.Category)
		return &category
	}
}
