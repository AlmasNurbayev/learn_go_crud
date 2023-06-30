package auth

import (
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JwtClaims struct {
	jwt.RegisteredClaims
	Exp   time.Time `json:"exp"`
	Id    int       `json:"id"`
	Email string    `json:"email"`
}

func GenerateJWT(key_jwt string, id int, email string) (string, error) {
	//fmt.Println(key_jwt)
	secret := []byte(key_jwt)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		JwtClaims{
			RegisteredClaims: jwt.RegisteredClaims{},
			Exp:              time.Now().Add(120 * time.Minute),
			//claims["authorized"] = true
			Id:    id,
			Email: email,
		})
	tokenString, err := token.SignedString(secret)
	//fmt.Println("tokenString", tokenString)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	return tokenString, nil
}

func VerifyJWT(tokenString string, key_jwt string) (*JwtClaims, error) {
	// pass your custom claims to the parser function
	var jwtPayload JwtClaims

	token, err := jwt.ParseWithClaims(tokenString, &jwtPayload, func(token *jwt.Token) (interface{}, error) {
		return []byte(key_jwt), nil
	})
	//fmt.Println("jwtPayload", jwtPayload)
	if err != nil {
		return nil, err
	}
	if !token.Valid {
		return nil, errors.New("token is not valid")
	}

	// type-assert `Claims` into a variable of the appropriate type
	return &jwtPayload, nil
}
