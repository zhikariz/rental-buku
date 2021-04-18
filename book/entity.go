package book

import (
	"rental-buku/loan"

	"gorm.io/gorm"
)

type Book struct {
	gorm.Model
	Title       string `json:"title" gorm:"type:text"`
	Description string `json:"description" gorm:"type:text"`
	Author      string `json:"author" gorm:"type:text"`
	Year        int    `json:"year" gorm:"type:year"`
	CategoryID  uint   `json:"category_id"`
	Loans       []loan.Loan
}
