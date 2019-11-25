package sqlstore_test

import (
	"testing"

	"github.com/DalerBakhriev/http-rest-api/internal/app/model"
	"github.com/DalerBakhriev/http-rest-api/internal/app/store"
	"github.com/DalerBakhriev/http-rest-api/internal/app/store/sqlstore"
	"github.com/stretchr/testify/assert"
)

func TestUserRepository_Create(t *testing.T) {

	db, teardown := sqlstore.TestDB(t, dataBaseURL)
	defer teardown("public.users")

	s := sqlstore.New(db)
	u := model.TestUser(t)

	assert.NoError(t, s.User().Create(u))
	assert.NotNil(t, u)
}

func TestUserRepository_Find(t *testing.T) {

	db, teardown := sqlstore.TestDB(t, dataBaseURL)
	defer teardown("public.users")

	s := sqlstore.New(db)
	u1 := model.TestUser(t)

	s.User().Create(u1)
	u2, err := s.User().Find(u1.ID)
	assert.NoError(t, err)
	assert.NotNil(t, u2)
}

func TestUserRepository_FindByEmail(t *testing.T) {

	db, teardown := sqlstore.TestDB(t, dataBaseURL)
	defer teardown("public.users")

	s := sqlstore.New(db)
	u1 := model.TestUser(t)
	_, err := s.User().FindByEmail(u1.Email)
	assert.EqualError(t, err, store.ErrRecordNotFound.Error())

	s.User().Create(u1)
	u2, err := s.User().FindByEmail(u1.Email)
	assert.NoError(t, err)
	assert.NotNil(t, u2)
}
