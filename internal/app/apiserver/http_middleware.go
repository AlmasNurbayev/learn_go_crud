package apiserver

import (
	"net/http"
	"strings"
	"time"

	"github.com/AlmasNurbayev/learn_go_crud/internal/app/auth"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

// CheckJWT - вытаскиваем токен из заголовка и проверяем
func (s *server) CheckJWT(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		bearer := r.Header.Get("Authorization")

		if bearer == "" {
			s.error(w, r, http.StatusUnauthorized, ErrNotAuthorized)
			s.logger.Error("token is empty")
			return
		}
		bearer = strings.TrimPrefix(bearer, "Bearer ")
		//fmt.Println(bearer)

		jwtPayload, _ := auth.VerifyJWT(bearer, s.config.KeyJwt)
		if jwtPayload == nil {
			s.error(w, r, http.StatusUnauthorized, ErrNotAuthorized)
			s.logger.Error("token is not valid")
			return
		}

		u, err := s.store.User().FindByEmail(jwtPayload.Email)
		if err != nil {
			s.error(w, r, http.StatusUnauthorized, ErrNotAuthorized)
			s.logger.Error("not found email of token")
			return
		}
		if u.Id != jwtPayload.Id {
			s.error(w, r, http.StatusUnauthorized, ErrNotAuthorized)
			s.logger.Error("not found Id of token")
			return
		}
		if jwtPayload.Exp.Before(time.Now()) {
			s.error(w, r, http.StatusUnauthorized, ErrNotAuthorized)
			s.logger.Error("token date is expired")
			return
		}

		//fmt.Println("jwtPayload", jwtPayload)
		//fmt.Println("err", err)

		next.ServeHTTP(w, r)

	})
}

// SetRequestID - добавляем в хедер уникальный id
func (s *server) SetRequestID(next http.Handler) http.Handler {
	return http.HandlerFunc((func(w http.ResponseWriter, r *http.Request) {
		id := uuid.New().String()
		w.Header().Set("X-Request-ID", id)
		next.ServeHTTP(w, r)
	}))
}

// LogRequest - пишем в лог о запросе и ответе
func (s *server) LogRequest(next http.Handler) http.Handler {
	return http.HandlerFunc((func(w http.ResponseWriter, r *http.Request) {
		//fmt.Println(r.Header)
		logger := s.logger.WithFields(logrus.Fields{
			"remote_addr": r.RemoteAddr,
			"request_id":  w.Header().Get("X-Request-Id"),
		})
		logger.Infof("started %s %s", r.Method, r.URL.RequestURI())
		start := time.Now()
		rw := &responseWriter{w, http.StatusOK}
		next.ServeHTTP(rw, r)

		logger.Infof("completed with %d %s in %v",
			rw.code, http.StatusText(rw.code),
			time.Since(start),
		)
	}))
}
