package seeds

import (
	"rental-buku/user"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func CreateUser(db *gorm.DB, name string, address string, email string, role string) error {
	passwordHash, _ := bcrypt.GenerateFromPassword([]byte("password"), bcrypt.MinCost)
	return db.Create(&user.User{Name: name, Address: address, Email: email, Password: string(passwordHash), Role: role}).Error
}
