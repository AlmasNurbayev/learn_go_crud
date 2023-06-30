package store

import "github.com/AlmasNurbayev/learn_go_crud/internal/app/model"

type UserRepository interface {
	Create(*model.User) error
	FindByEmail(string) (*model.User, error)
	FindAll() ([]*model.User, error)
}
