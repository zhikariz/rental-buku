package seeds

import (
	"errors"
	"rental-buku/user"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func CreateUser(db *gorm.DB, name string, address string, email string, role string) error {
	var user user.User

	err := db.Where("email = ?", email).Find(&user).Error

	if err != nil {
		return err
	}

	if user.ID > 0 {
		return errors.New("User is exist")
	}

	passwordHash, _ := bcrypt.GenerateFromPassword([]byte("password"), bcrypt.MinCost)
	user.Name = name
	user.Address = address
	user.Email = email
	user.Password = string(passwordHash)
	user.Role = role

	return db.Create(&user).Error
}
