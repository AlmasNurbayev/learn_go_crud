package apiserver

import (
	"encoding/json"
	"errors"
	"net/http"
)

var (
	ErrIncorrectEmailPassword = errors.New("incorrect email or password")
	ErrNotAuthorized          = errors.New("no authorize data")
	ErrInternalError          = errors.New("internal error")
	ErrBadRequest             = errors.New("bad request")
)

func (s *server) error(w http.ResponseWriter, r *http.Request, code int, err error) {
	//fmt.Println("error ", err.Error())
	s.respond(w, r, code, map[string]string{"error": err.Error()})

}

func (s *server) respond(w http.ResponseWriter, r *http.Request, code int, data interface{}) {
	w.WriteHeader(code)
	// fmt.Println("code ", code)
	// fmt.Println("request ", r)
	// fmt.Println("writer ", w)
	if data != nil {
		err := json.NewEncoder(w).Encode(data)
		// fmt.Println("writer2 ", w)
		// fmt.Println("err2 ", err)
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
	}
	// fmt.Println("exit respond")
}
