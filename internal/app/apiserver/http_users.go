package apiserver

import (
	"encoding/json"
	"net/http"

	"github.com/AlmasNurbayev/learn_go_crud/internal/app/auth"
	"github.com/AlmasNurbayev/learn_go_crud/internal/app/model"
	"github.com/AlmasNurbayev/learn_go_crud/other"
)

func (s *server) UserAuth() http.HandlerFunc {
	type request struct {
		Email    string `json: "email"`
		Password string `json: "password"`
	}
	return func(w http.ResponseWriter, r *http.Request) {
		req := &request{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			s.error(w, r, http.StatusBadRequest, ErrBadRequest)
			s.logger.Error("user body " + other.ToJSON(req) + " decode error " + err.Error())
			return
		}
		u, err := s.store.User().FindByEmail(req.Email)
		if err != nil {
			s.error(w, r, http.StatusUnauthorized, ErrIncorrectEmailPassword)
			s.logger.Info("user " + u.Email + " dont find email auth " + err.Error())
			return
		}
		if !u.ComparePassword(req.Password) {
			s.error(w, r, http.StatusUnauthorized, ErrIncorrectEmailPassword)
			s.logger.Info("user " + u.Email + " denied auth " + err.Error())
			return
		}
		// if err := u.Validate(); err != nil {
		// 	s.error(w, r, http.StatusUnauthorized, ErrIncorrectEmailPassword)
		// 	if err != nil {
		// 		s.logger.Info("user " + u.Email + " denied auth " + err.Error())
		// 	}
		// 	return
		// }

		jwt, err := auth.GenerateJWT(s.config.KeyJwt, u.Id, u.Email)
		if err != nil {
			s.error(w, r, http.StatusInternalServerError, ErrInternalError)
			s.logger.Error("user " + u.Email + " get JWT error " + err.Error())
			return
		}
		s.logger.Info("user " + u.Email + " success auth")
		w.Header().Set("Content-Type", "application/json")
		s.respond(w, r, http.StatusOK, jwt)
	}
}

func (s *server) UserCreate() http.HandlerFunc {
	type request struct {
		Email    string `json: "email"`
		Password string `json: "password"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		req := &request{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			s.error(w, r, http.StatusBadRequest, err)
			return
		}
		u := &model.User{
			Email:    req.Email,
			Password: req.Password,
		}
		s.logger.Info("users-create request", u)
		if err := s.store.User().Create(u); err != nil {
			s.error(w, r, http.StatusUnprocessableEntity, err)
			return
		}

		u.Sanitize()

		s.respond(w, r, http.StatusCreated, u)

	}
}
