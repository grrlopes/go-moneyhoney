package helper

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/grrlopes/go-moneyhoney/src/domain/entity"
	"golang.org/x/crypto/bcrypt"
)

func CreatePassword(passw *entity.Users) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(passw.Password), 14)
	if err != nil {
		return err
	}
	passw.Password = string(bytes)
	return nil
}

func ValidPassword(passw *entity.Users, passwfromdb string) error {
	err := bcrypt.CompareHashAndPassword([]byte(passw.Password), []byte(passwfromdb))
	if err != nil {
		return err
	}
	return nil
}

func GenerateJwt(user *entity.Users) (string, error) {
	jwtKey := []byte(os.Getenv("JWTKEY"))
	expirationTime := time.Now().Add(1 * time.Hour)

	claims := &entity.Users{
		Email:  user.Email,
		Author: user.Author,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return tokenString, err
	}

	return tokenString, nil
}
