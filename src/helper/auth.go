package helper

import (
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/grrlopes/go-moneyhoney/src/domain/entity"
	"golang.org/x/crypto/bcrypt"
)

func CreatePassword(passw *entity.Users) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(passw.Password), 14)
	if err != nil {
		return string(bytes), err
	}

	return string(bytes), nil
}

func ValidPassword(passw *entity.Users, passwfromdb string) error {
	err := bcrypt.CompareHashAndPassword([]byte(passwfromdb), []byte(passw.Password))
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
			IssuedAt:  time.Now().Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return tokenString, err
	}

	return tokenString, nil
}

func VerifyJwt(tokenString string) error {
	token, err := jwt.ParseWithClaims(
		tokenString,
		&entity.Users{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("JWTKEY")), nil
		},
	)

	if err != nil {
		return err
	}

	if !token.Valid {
		return errors.New("token not valid")
	}

	claims, ok := token.Claims.(*entity.Users)
	if !ok {
		return errors.New("couldn't parse claims")
	}
	if claims.ExpiresAt < time.Now().UTC().Unix() {
		return errors.New("token is expired")
	}

	return err
}
