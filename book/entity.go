package book

import (
	"rental-buku/loan"
	"time"

	"gorm.io/gorm"
)

type Book struct {
	gorm.Model
	Title       string    `json:"title" gorm:"type:text"`
	Description string    `json:"description" gorm:"type:text"`
	Author      string    `json:"author" gorm:"type:text"`
	Year        time.Time `json:"year" gorm:"type:year"`
	CategoryID  uint      `json:"category_id"`
	Loans       []loan.Loan
}
