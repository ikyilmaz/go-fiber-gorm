package models

type UserModel struct {
	BaseModel
	FirstName string `gorm:"type:varchar(32);"`
	LastName  string `gorm:"type:varchar(32);"`
	Username  string `gorm:"type:varchar(32);not null;unique;"`
	Email     string `gorm:"type:varchar(128);not null;unique;"`
	Password  string `gorm:"type:varchar(255);not null;unique;"`
}

func (u UserModel) TableName() string { return "users" }
