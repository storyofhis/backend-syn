package model

import "gorm.io/gorm"

type Book struct {
	gorm.Model
	Title       string
	CategoryId  string
	Author      string
	Description string
	Stock       int32
}
