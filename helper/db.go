package helper

import (
	"os"
	"rental-buku/book"
	"rental-buku/category"
	"rental-buku/loan"
	"rental-buku/payment"
	"rental-buku/user"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func SetupDB() *gorm.DB {
	dsn := os.Getenv("DSN_CONNECTION")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Error Connecting to Database")
	}
	db.AutoMigrate(&user.User{}, &category.Category{}, &book.Book{}, &loan.Loan{}, &payment.FinePayment{})

	return db
}
