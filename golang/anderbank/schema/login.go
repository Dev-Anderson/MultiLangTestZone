package schema

type Login struct {
	ID       uint   `gorm:"primaryKey; autoIncrement=true"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
