package services

import "gorm.io/gorm"

type UserService struct{ db *gorm.DB }

func NewUserService(db *gorm.DB) *UserService { return &UserService{db: db} }

func (u *UserService) GetManyUser() {
	panic("implement me")
}

func (u *UserService) CreateUser() {
	panic("implement me")
}
