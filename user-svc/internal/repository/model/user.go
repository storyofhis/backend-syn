package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string
	FullName string
	Email    string
	Address  string
	Password string
	Phone    string
}
