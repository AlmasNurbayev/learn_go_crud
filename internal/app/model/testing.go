package model

import "testing"

func TestUser(t *testing.T) *User {
	return &User{
		Id:       9,
		Email:    "user@example.org",
		Password: "password",
	}
}
