package apiserver

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/AlmasNurbayev/learn_go_crud/internal/app/model"
	"github.com/AlmasNurbayev/learn_go_crud/internal/app/store/teststore"
	"github.com/stretchr/testify/assert"
)

func TestServer_loadConfig(t *testing.T) {
	s := newServer(teststore.New(), NewConfig())
	assert.NotEmpty(t, s.config.BindAddr)
	assert.NotEmpty(t, s.config.LogLevel)
	assert.NotEmpty(t, s.config.KeyJwt)

}

func TestServer_HandleUsersCreate(t *testing.T) {
	s := newServer(teststore.New(), NewConfig())

	testCases := []struct {
		name         string
		payload      interface{}
		expectedCode int
	}{
		{
			name: "valid",
			payload: map[string]string{
				"email":    "user@example.com",
				"password": "strongWord",
			},
			expectedCode: http.StatusCreated,
		},
		{
			name:         "invalid payload",
			payload:      "invalid",
			expectedCode: http.StatusBadRequest,
		},
		{
			name: "invalid keys",
			payload: map[string]string{
				"email": "invalid",
				//"password": "strongWord",
			},
			expectedCode: http.StatusUnprocessableEntity,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			rec := httptest.NewRecorder()
			b := &bytes.Buffer{}
			json.NewEncoder(b).Encode(tc.payload)
			req, _ := http.NewRequest(http.MethodPost, "/users", b)
			s.ServeHTTP(rec, req)
			assert.Equal(t, tc.expectedCode, rec.Code)
		})
	}

}

func TestServer_HandleAuth(t *testing.T) {

	u := model.TestUser(t)
	store := teststore.New()

	s := newServer(store, NewConfig())
	store.User().Create(u)

	testCases := []struct {
		name         string
		payload      interface{}
		expectedCode int
	}{
		{
			name: "valid",
			payload: map[string]string{
				"email":    u.Email,
				"password": u.Password,
			},
			expectedCode: http.StatusOK,
		},
		{
			name:         "invalid payload",
			payload:      "invalid",
			expectedCode: http.StatusBadRequest,
		},
		{
			name: "invalid email",
			payload: map[string]string{
				"email":    "invalid",
				"password": u.Password,
			},
			expectedCode: http.StatusUnauthorized,
		},
		{
			name: "invalid password",
			payload: map[string]string{
				"email":    u.Email,
				"password": "invalid password",
			},
			expectedCode: http.StatusUnauthorized,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			//s.logger.Info(" ==== running test " + tc.name)
			rec := httptest.NewRecorder()
			b := &bytes.Buffer{}
			json.NewEncoder(b).Encode(tc.payload)
			req, _ := http.NewRequest(http.MethodPost, "/auth", b)
			s.ServeHTTP(rec, req)
			assert.Equal(t, tc.expectedCode, rec.Code)
		})
	}

}

func Test_CheckJWT(t *testing.T) {
	u := model.TestUser(t)
	store := teststore.New()

	s := newServer(store, NewConfig())
	store.User().Create(u)

	testCases := []struct {
		name         string
		payload      model.User
		expectedCode int
	}{
		{
			name: "valid",
			payload: model.User{
				Id:       u.Id,
				Email:    u.Email,
				Password: u.Password,
			},
			expectedCode: http.StatusOK,
		},
		{
			name: "not valid email",
			payload: model.User{
				Id:       u.Id,
				Email:    u.Email + "xxx",
				Password: u.Password,
			},
			expectedCode: http.StatusUnauthorized,
		},
		{
			name: "not valid password",
			payload: model.User{
				Id:       u.Id,
				Email:    u.Email,
				Password: u.Password + "xxx",
			},
			expectedCode: http.StatusUnauthorized,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			//s.logger.Info(" ==== running test " + tc.name)
			fmt.Println("email", tc.payload.Email)
			fmt.Println("password", tc.payload.Password)
			rec := httptest.NewRecorder()
			b := &bytes.Buffer{}
			json.NewEncoder(b).Encode(tc.payload)
			req, _ := http.NewRequest(http.MethodPost, "/auth", b)
			//req, _ := http.NewRequest(http.MethodPost, "api/users", b)
			s.ServeHTTP(rec, req)

			type tokenT struct {
				Token string `json:"Token"`
			}
			var bearer tokenT
			//fmt.Println("rec.body", rec.Body)
			err := json.NewDecoder(rec.Body).Decode(&bearer)
			if err != nil {
				fmt.Println(err)
				assert.Fail(t, "error on decode")
			}
			//fmt.Println("bearer", bearer)
			token_res := strings.TrimPrefix(bearer.Token, "Bearer ")

			req2, _ := http.NewRequest(http.MethodGet, "/api/users", nil)
			rec2 := httptest.NewRecorder()
			req2.Header.Add("Authorization", "Bearer "+token_res)
			s.ServeHTTP(rec2, req2)

			assert.Equal(t, tc.expectedCode, rec2.Code)
		})
	}

}
