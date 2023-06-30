package sqlstore_test

import (
	"testing"

	"github.com/AlmasNurbayev/learn_go_crud/internal/app/model"
	"github.com/AlmasNurbayev/learn_go_crud/internal/app/store"
	"github.com/AlmasNurbayev/learn_go_crud/internal/app/store/sqlstore"
	"github.com/stretchr/testify/assert"
)

func TestUserRepository_Create(t *testing.T) {
	db, teardown := sqlstore.TestDB(t, databaseTestURL)
	defer teardown("users")
	s := sqlstore.New(db)
	u := model.TestUser(t)
	//t.Log(u)
	assert.NoError(t, s.User().Create(u))
	assert.NotNil(t, u)
}

func TestUserRepository_FindByEmail(t *testing.T) {
	db, teardown := sqlstore.TestDB(t, databaseTestURL)
	defer teardown("users")
	s := sqlstore.New(db)
	_, err := s.User().FindByEmail(model.TestUser(t).Email) // проверяем отсутствие
	assert.EqualError(t, err, store.ErrRecordNotFound.Error())

	s.User().Create(model.TestUser(t))
	u, err := s.User().FindByEmail(model.TestUser(t).Email)
	//t.Log(u)
	assert.NoError(t, err) // проверяем наличие
	assert.NotNil(t, u)
	assert.Equal(t, u.Email, "user@example.org")
}

func TestUserRepository_FindAll(t *testing.T) {
	db, teardown := sqlstore.TestDB(t, databaseTestURL)
	defer teardown("users")
	s := sqlstore.New(db)
	us, _ := s.User().FindAll() // проверяем отсутствие
	assert.Equal(t, len(us), 0)

	s.User().Create(model.TestUser(t))
	u, err := s.User().FindAll()
	//t.Log(u)
	assert.NoError(t, err) // проверяем наличие
	assert.NotNil(t, u)
	assert.GreaterOrEqual(t, len(u), 1)
}
