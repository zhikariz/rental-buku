package payment

import "gorm.io/gorm"

type FinePayment struct {
	gorm.Model
	Receipt string
	Amount  float64
	LoanID  uint
}
