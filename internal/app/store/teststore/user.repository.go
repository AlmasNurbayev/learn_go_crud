package teststore

import (
	"github.com/AlmasNurbayev/learn_go_crud/internal/app/model"
	"github.com/AlmasNurbayev/learn_go_crud/internal/app/store"
)

type UserRepository struct {
	store *Store
	users []*model.User
}

func (r *UserRepository) Create(u *model.User) error {
	if err := u.Validate(); err != nil {
		return err
	}

	if err := u.BeforeCreate(); err != nil {
		return err
	}

	r.users = append(r.users, u)
	u.Id = len(r.users) + 1

	return nil
}

func (r *UserRepository) FindByEmail(email string) (*model.User, error) {
	//u, ok := r.users[email]

	for _, pp := range r.users {
		if pp.Email == email {
			return pp, nil
		}
	}
	return nil, store.ErrRecordNotFound
}

func (r *UserRepository) FindAll() ([]*model.User, error) {

	return r.users, nil
}
