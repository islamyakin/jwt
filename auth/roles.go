package auth

import (
	"fmt"
	"github.com/islamyakin/jwt/models"
	"gorm.io/gorm"
)

func GetUserRoles(username string) (string, error) {
	var user models.User
	result := db.Where("username = ?", username).First(&user)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return "", fmt.Errorf("user not found")
		}
		return "", result.Error
	}
	return user.Roles, nil
}
