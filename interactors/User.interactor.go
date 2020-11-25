package interactors

import (
	"hexagonal/entities"
)

type UserItr struct{}

func (*UserItr) FindAll() ([]entities.User, error) {
	return UserRepo.FindAll()
}

func (*UserItr) Create(u *entities.User) error {
	return UserRepo.Create(u)
}

func (*UserItr) GetById(id uint64) (entities.User, error) {
	return UserRepo.GetById(id)
}
