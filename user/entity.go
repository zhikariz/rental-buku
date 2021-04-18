package user

import (
	"rental-buku/loan"
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name            string    `json:"name" gorm:"size:255"`
	Address         string    `json:"address" gorm:"type:text"`
	Photo           string    `json:"photo" gorm:"type:text"`
	Email           string    `json:"email" gorm:"size:255"`
	EmailVerifiedAt time.Time `json:"email_verified_at"`
	Password        string    `json:"password" gorm:"type:text"`
	Role            string    `json:"role" gorm:"type:ENUM('Admin', 'Member')"`
	Loans           []loan.Loan
}
