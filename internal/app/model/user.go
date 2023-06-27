package model

import (
	"github.com/go-playground/validator/v10"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Id                 int    `json:"id"`
	Email              string `json:"Email" validate:"required,email"`
	Password           string `json:"Password,omitempty" validate:"required,min=6,max=100"`
	Encrypted_password string `json:"-"`
}

func (u *User) Sanitize() {
	u.Password = ""

}

func (u *User) ComparePassword(password string) bool {
	compare := bcrypt.CompareHashAndPassword([]byte(u.Encrypted_password), []byte(password))
	if compare == nil {
		return true
	} else {
		return false
	}

}

func (u *User) BeforeCreate() error {

	if len(u.Password) > 0 {
		enc, err := encryptString(u.Password)
		if err != nil {
			return err
		}
		u.Encrypted_password = enc

	}
	return nil
}

func encryptString(s string) (string, error) {
	b, err := bcrypt.GenerateFromPassword([]byte(s), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

func (u *User) Validate() error {

	validate := validator.New()
	//fmt.Println(u)
	err := validate.Struct(u)
	//validationErrors := err.(validator.ValidationErrors)

	return err

}
