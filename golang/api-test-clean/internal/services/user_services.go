package services

import (
	"api-test/internal/repository"
	"api-test/pkg/models"
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	UserRepository *repository.UserRepository
}

// Exemplo de criptografica de senha novo usuario
func (s *UserService) CreateUser(user *models.User) error {
	// Criptografar a senha antes de salvar no banco de dados
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user.Password = string(hashedPassword)
	//chamar o metodo create userRepository para salvar o usuario no banco de dados
	return s.UserRepository.Create(user)
}

func (s *UserService) GetUserByID(id int) (*models.User, error) {
	//chamar o metodo FindByID do UserRepository para obter o usuario pelo ID
	return s.UserRepository.FindByID(id)
}

func (s *UserService) UpdateUser(user *models.User) error {
	//chamar o metod update do UserReposotyr para atulizar o user
	return s.UserRepository.Update(user)
}

func (s *UserService) DeleteUser(id int) error {
	// chamar o metod delete do UserRepository para excluir o usuario pelo ID
	return s.UserRepository.Delete(id)
}

func (s *UserService) CompareHashAndPassword(hash, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
}

func (s *UserService) AuthenticateUser(email, password string) (*models.User, error) {
	//buscar o user pelo email
	user, err := s.UserRepository.FindByEmail(email)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, errors.New("Usuario nao encontrado")
	}

	//Comparar a senha fornecida com a senha armazenada
	err = s.CompareHashAndPassword(user.Password, password)
	if err != nil {
		return nil, errors.New("Senha incorreta")
	}

	return user, nil
}
