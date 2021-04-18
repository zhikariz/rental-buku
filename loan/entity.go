package loan

import (
	"rental-buku/payment"
	"time"

	"gorm.io/gorm"
)

type Loan struct {
	gorm.Model
	UserID       uint
	BookID       uint
	BorrowedDate time.Time
	DueDate      time.Time
	ReturnDate   time.Time
	FinePayment  payment.FinePayment
}
