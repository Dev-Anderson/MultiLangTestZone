package models

type Login struct {
	ID       uint   `gorm:"primaryKey; autoIncrement:true" swaggerignore:"true"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
