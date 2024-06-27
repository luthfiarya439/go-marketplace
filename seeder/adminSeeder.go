package seeder

import (
	"go-marketplace/config"
	"go-marketplace/models"

	"golang.org/x/crypto/bcrypt"
)

func AdminSeeder() error {
	var user models.User
	if err := config.DB.Model(&user).Where("email = ?", "admin@mail.com").First(&user).RowsAffected; err == 0 {
		passwordHash, err := bcrypt.GenerateFromPassword([]byte("admin@mail.com"), bcrypt.DefaultCost)
		if err != nil {
			return err
		}

		user = models.User{
			Name:     "Admin",
			Email:    "admin@mail.com",
			Password: string(passwordHash),
			Role:     "admin",
		}
		if err := config.DB.Create(&user).Error; err != nil {
			return err
		}
		return err
	}
	return nil
}
