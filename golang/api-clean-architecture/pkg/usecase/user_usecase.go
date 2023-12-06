package usecase

import "api-clean-archetecture/pkg/entity"

//contem a logica de negocio independente do framework e da camada de entrega

type UserReposity interface {
	Create(user *entity.User) error
	FindByID(id uint) (*entity.User, error)
	Update(user *entity.User) error
	Delete(id uint) error
}

type UserUseCase struct {
	UserReposity UserReposity
}

func (uc *UserUseCase) CreateUser(user *entity.User) error {
	return uc.UserReposity.Create(user)
}

func (uc *UserUseCase) GetUserByID(id uint) (*entity.User, error) {
	return uc.UserReposity.FindByID(id)
}

func (uc *UserUseCase) UpdateUser(user *entity.User) error {
	return uc.UserReposity.Update(user)
}

func (uc *UserUseCase) DeleteUser(id uint) error {
	return uc.UserReposity.Delete(id)
}
