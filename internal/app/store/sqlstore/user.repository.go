package sqlstore

import (
	"database/sql"

	"github.com/AlmasNurbayev/learn_go_crud/internal/app/model"
	"github.com/AlmasNurbayev/learn_go_crud/internal/app/store"
)

type UserRepository struct {
	store *Store
}

func (r *UserRepository) Create(u *model.User) error {
	if err := u.Validate(); err != nil {
		return err
	}

	if err := u.BeforeCreate(); err != nil {
		return err
	}

	return r.store.db.QueryRow(
		"INSERT INTO users (email, encrypted_password) VALUES ($1, $2) RETURNING id",
		u.Email, u.Encrypted_password).Scan(&u.Id)
}

func (r *UserRepository) FindByEmail(email string) (*model.User, error) {

	u := &model.User{}
	if err := r.store.db.QueryRow(
		"select id, email, encrypted_password from users where email = $1",
		email).Scan(&u.Id, &u.Email, &u.Encrypted_password); err != nil {
		if err == sql.ErrNoRows {
			return nil, store.ErrRecordNotFound
		}

		return nil, err
	}

	return u, nil
}

func (r *UserRepository) FindAll() ([]*model.User, error) {

	//defer rows.Close()

	users := []*model.User{}
	rows, err := r.store.db.Query("select id, email, encrypted_password from users")

	if err != nil {
		return nil, err
	}
	if err == sql.ErrNoRows {
		return nil, store.ErrRecordNotFound
	}

	for rows.Next() {
		u := model.User{}
		if err := rows.Scan(&u.Id, &u.Email, &u.Encrypted_password); err != nil {
			return nil, store.ErrRecordNotFound
		}
		users = append(users, &u)

	}

	return users, nil

}
