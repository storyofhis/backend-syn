package model

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string
	FullName string
	Email    string
	Address  string
	Password string
	Phone    string
}

type Borrow struct {
	gorm.Model
	UserId     int
	BookId     int
	BorrowDate time.Time
	ReturnDate time.Time
	Status     bool
}
