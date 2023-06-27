package auth

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateJWT(key_jwt string, id int, email string) (string, error) {
	//fmt.Println(key_jwt)
	secret := []byte(key_jwt)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"exp": time.Now().Add(120 * time.Minute),
			//claims["authorized"] = true
			"id":    id,
			"email": email,
		})
	tokenString, err := token.SignedString(secret)
	//fmt.Println("tokenString", tokenString)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	return "Bearer " + tokenString, nil

}
