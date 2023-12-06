package repository

import (
	"api-clean-archetecture/pkg/entity"

	"gorm.io/gorm"
)

//lida com a persistencia de dados, sem conhecimento de detalhes de implementacao

type UserReposity struct {
	DB *gorm.DB
}

func NewUserReposity(db *gorm.DB) *UserReposity {
	return &UserReposity{DB: db}
}

func (r *UserReposity) Create(user *entity.User) error {
	return r.DB.Create(user).Error
}

func (r *UserReposity) FindByID(id uint) (*entity.User, error) {
	var user entity.User
	err := r.DB.First(&user, id).Error
	return &user, err
}

func (r *UserReposity) Update(user *entity.User) error {
	return r.DB.Save(user).Error
}

func (r *UserReposity) Delete(id uint) error {
	return r.DB.Delete(&entity.User{}, id).Error
}
