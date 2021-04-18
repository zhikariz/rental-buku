package category

import (
	"rental-buku/book"

	"gorm.io/gorm"
)

type Category struct {
	gorm.Model
	Name  string `json:"name" gorm:"size:255"`
	Books []book.Book
}
