package book

import "gorm.io/gorm"

type Book struct {
	gorm.DB
	Title       string
	Author      string
	AuthorId    string
	Description string
}
