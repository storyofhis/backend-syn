package category

import "gorm.io/gorm"

type Category struct {
	gorm.Model
	Name        string
	Description string
}
