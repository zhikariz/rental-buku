package seeds

import (
	"rental-buku/helper"

	"gorm.io/gorm"
)

func All() []helper.Seed {
	return []helper.Seed{
		{
			Name: "CreateAdmin",
			Run: func(db *gorm.DB) error {
				return CreateUser(db, "Admin Rental", "Indonesia", "admin@gmail.com", "Admin")
			},
		},
		{
			Name: "CreateMember",
			Run: func(db *gorm.DB) error {
				return CreateUser(db, "Member Rental", "Indonesia", "member@gmail.com", "Member")
			},
		},
	}
}
