package apiserver

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

var (
	ErrIncorrectEmailPassword = errors.New("incorrect email or password")
	ErrInternalError          = errors.New("internal error")
	ErrBadRequest             = errors.New("bad request")
)

func (s *server) error(w http.ResponseWriter, r *http.Request, code int, err error) {
	fmt.Println("error ", r)
	s.respond(w, r, code, map[string]string{"error": err.Error()})

}

func (s *server) respond(w http.ResponseWriter, r *http.Request, code int, data interface{}) {
	fmt.Println("respond ", r)
	w.WriteHeader(code)
	fmt.Println("code ", code)
	if data != nil {
		if err := json.NewEncoder(w).Encode(data); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
	}
}
