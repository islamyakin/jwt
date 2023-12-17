package auth

import (
	"fmt"
	"github.com/islamyakin/jwt/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var (
	secretKey = []byte("theworldiyourhand")
	db        *gorm.DB
)

const (
	host     = "localhost"
	port     = 4502
	user     = "postgres"
	password = "12345678"
	dbname   = "batman"
)

func init() {
	dsn := fmt.Sprint("host=", host, " user=", user, " password=", password, " dbname=", dbname, " port=", port, " sslmode=disable TimeZone=Asia/Jakarta")

	var err error
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&models.User{})
	fmt.Println("Successfully connected!")
}

func CreateUser(username, password string) error {
	user := models.User{Username: username, Password: password}
	result := db.Create(&user)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func AuthenticateUser(username, password string) bool {
	var user models.User
	result := db.Where("username = ?", username).First(&user)
	if result.Error != nil || result.RowsAffected == 0 {
		return false
	}
	return user.Password == password
}

func CreateToken(username string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"username": username,
			"exp":      time.Now().Add(time.Hour * 1).Unix(),
		})
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", nil
	}
	return tokenString, nil
}

func verifyToken(tokenString string) error {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})
	if err != nil {
		return err
	}
	if !token.Valid {
		return fmt.Errorf("invalid token")
	}
	return nil
}
