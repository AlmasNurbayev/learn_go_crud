package teststore

import (
	"github.com/AlmasNurbayev/learn_go_crud/internal/app/model"
	"github.com/AlmasNurbayev/learn_go_crud/internal/app/store"
)

type Store struct {
	//config         *Config
	userRepository *UserRepository
}

func New() *Store {
	return &Store{}
}

// func (s *Store) Open() error {
// 	db, err := sql.Open("postgres", s.config.DatabaseURL)
// 	if err != nil {
// 		return err
// 	}

// 	if err := db.Ping(); err != nil {
// 		return err
// 	}

// 	s.db = db

// 	return nil
// }

// func (s *Store) Close() {
// 	s.db.Close()
// }

func (s *Store) User() store.UserRepository {
	if s.userRepository != nil {
		return s.userRepository
	}

	s.userRepository = &UserRepository{
		store: s,
		users: make(map[string]*model.User),
	}

	return s.userRepository
}
