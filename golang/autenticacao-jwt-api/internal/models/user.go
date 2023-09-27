package models

type User struct {
	ID       uint   `gorm:"primaryKey; autoIncrement:true"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
