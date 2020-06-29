package models

type UserModel struct {
	BaseModel
	FirstName string
	LastName  string
	Username  string
	Email     string
	Password  string
}

func (u UserModel) TableName() string { return "users" }
