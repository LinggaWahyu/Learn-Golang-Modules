package repository

import "learn-golang-unit-test/entity"

type CategoryRespository interface {
	FindById(id string) *entity.Category
}
