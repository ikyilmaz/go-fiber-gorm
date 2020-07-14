package models

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserModel struct {
	BaseModel
	FirstName string `gorm:"type:varchar(32);"`
	LastName  string `gorm:"type:varchar(32);"`
	Username  string `gorm:"type:varchar(32);not null;unique;"`
	Email     string `gorm:"type:varchar(128);not null;unique;"`
	Password  string `gorm:"type:varchar(255);not null;unique;"`
}

func (u UserModel) TableName() string { return "users" }

func (u *UserModel) BeforeSave(db *gorm.DB) error {
	if u.Password != "" {
		byteHashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), 12)

		if err != nil {
			return err
		}

		u.Password = string(byteHashedPassword)
	}

	return nil
}
