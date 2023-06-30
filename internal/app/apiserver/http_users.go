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
		//w.Header().Set("Access-Control-Allow-Origin", "*")
		req := &request{}
		// fmt.Println("point1")
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			s.error(w, r, http.StatusBadRequest, ErrBadRequest)
			s.logger.Error("user body " + other.ToJSON(req) + " decode error " + err.Error())
			return
		}
		// fmt.Println("point2")
		u, err := s.store.User().FindByEmail(req.Email)
		if err != nil {
			s.error(w, r, http.StatusUnauthorized, ErrIncorrectEmailPassword)
			//fmt.Println("point21")
			s.logger.Info("user " + req.Email + " dont find email auth " + err.Error())
			//fmt.Println("point22")
			return
		}
		// fmt.Println("point3")
		if !u.ComparePassword(req.Password) {
			s.error(w, r, http.StatusUnauthorized, ErrIncorrectEmailPassword)
			s.logger.Info("user " + other.ToJSON(req) + " incorrect hash password")
			return
		}
		// fmt.Println("point4")
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
		// fmt.Println("point5")
		s.logger.Info("user " + u.Email + " success auth")
		w.Header().Set("Content-Type", "application/json")

		type tokenT struct {
			Token string `json: "Token"`
		}

		body := tokenT{Token: jwt}
		//b := &bytes.Buffer{}

		//json.NewEncoder(w).Encode(body)
		s.respond(w, r, http.StatusOK, body)
	}
}

func (s *server) UserCreate() http.HandlerFunc {
	type request struct {
		Email    string `json: "email"`
		Password string `json: "password"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		//w.Header().Set("Access-Control-Allow-Origin", "*")
		req := &request{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			s.logger.Error("user create decode error")
			s.error(w, r, http.StatusBadRequest, err)
			return
		}
		u := &model.User{
			Email:    req.Email,
			Password: req.Password,
		}
		s.logger.Info("users-create request", u)
		if err := s.store.User().Create(u); err != nil {
			s.logger.Error("user create store error")
			s.error(w, r, http.StatusUnprocessableEntity, err)
			return
		}

		u.Sanitize()

		s.respond(w, r, http.StatusCreated, u)

	}
}

func (s *server) UserGet() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		//w.Header().Set("Access-Control-Allow-Origin", "*")
		//email := r.URL.Query().Get("email")
		//fmt.Println("email", email)

		users, err := s.store.User().FindAll()
		if err != nil {
			s.logger.Error("user get from store error")
			s.error(w, r, http.StatusInternalServerError, err)
			return
		}

		s.logger.Info("users-get request done")
		s.respond(w, r, http.StatusOK, users)
	}
}
