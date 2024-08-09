package model

import "gorm.io/gorm"

type Author struct {
	gorm.Model
	Name      string
	Biography string
}
