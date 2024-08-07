package model

import "gorm.io/gorm"

type Book struct {
	gorm.Model
	Title       string
	Author      string
	Description string
}
