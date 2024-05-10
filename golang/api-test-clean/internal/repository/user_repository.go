package repository

import (
	"api-test/pkg/models"
	"database/sql"
)

type UserRepository struct {
	DB *sql.DB
}

func (r *UserRepository) FindByID(id int) (*models.User, error) {
	query := "SELECT id, name, email, password FROM users where id = $1"
	row := r.DB.QueryRow(query, id)
	user := &models.User{}
	err := row.Scan(&user.ID, &user.Name, &user.Email, &user.Password)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return user, nil
}

func (r *UserRepository) FindByEmail(email string) (*models.User, error) {
	query := "select id, name, email, password from users where email = $1"
	row := r.DB.QueryRow(query, email)
	user := &models.User{}
	err := row.Scan(&user.ID, &user.Name, &user.Email, &user.Password)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return user, nil

}

func (r *UserRepository) Create(user *models.User) error {
	query := "insert into users (name, email, password) values($1, $2, $3) returning id"
	err := r.DB.QueryRow(query, user.Name, user.Email, user.Password).Scan(&user.ID)
	if err != nil {
		return err
	}
	return nil
}

func (r *UserRepository) Update(user *models.User) error {
	query := "update users set name = $2, email = $3 where id = $1"
	_, err := r.DB.Exec(query, user.ID, user.Name, user.Email)
	if err != nil {
		return err
	}
	return nil
}

func (r *UserRepository) Delete(id int) error {
	query := "delete from users where id = $1"
	_, err := r.DB.Exec(query, id)
	if err != nil {
		return err
	}
	return nil
}
