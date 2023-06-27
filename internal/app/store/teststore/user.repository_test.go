package teststore_test

import (
	"testing"

	"github.com/AlmasNurbayev/learn_go_crud/internal/app/model"
	"github.com/AlmasNurbayev/learn_go_crud/internal/app/store"
	"github.com/AlmasNurbayev/learn_go_crud/internal/app/store/teststore"
	"github.com/stretchr/testify/assert"
)

func TestUserRepository_Create(t *testing.T) {

	s := teststore.New()
	u := model.TestUser(t)
	//t.Log(u)
	assert.NoError(t, s.User().Create(u))
	assert.NotNil(t, u)
}

func TestUserRepository_FindByEmail(t *testing.T) {
	s := teststore.New()
	_, err := s.User().FindByEmail(model.TestUser(t).Email) // проверяем отсутствие
	assert.EqualError(t, err, store.ErrRecordNotFound.Error())

	s.User().Create(model.TestUser(t))
	u, err := s.User().FindByEmail(model.TestUser(t).Email)
	//t.Log(u)
	assert.NoError(t, err) // проверяем наличие
	assert.NotNil(t, u)
	assert.Equal(t, u.Email, "user@example.org")
}
