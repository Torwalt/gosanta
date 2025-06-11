package postgres_test

import (
	"context"
	"os"
	"testing"

	awards "gosanta/internal"
	"gosanta/internal/postgres"

	"github.com/stretchr/testify/assert"
	"github.com/uptrace/bun/dbfixture"
)

func TestUserGet(t *testing.T) {
	ctx := context.Background()
	db := getDb(t)
	resetDb(t, db, ctx)
	ur := postgres.NewUserRepository(db.DB)

	fixture := dbfixture.New(db)
	err := fixture.Load(ctx, os.DirFS("testdata"), "awards.yml")
	if err != nil {
		t.Errorf("could not load fixture: %v", err)
	}

	user, err := ur.Get(1)
	assert.Nil(t, err)
	assert.NotNil(t, user)
	assert.Equal(t, awards.UserId(1), user.Id)
	assert.Equal(t, awards.CompanyId(1), user.CompanyId)
	assert.Equal(t, 2, len(user.Awards))
	assert.Equal(t, awards.UserId(1), user.Awards[0].AssignedTo)
	assert.Equal(t, awards.UserId(1), user.Awards[1].AssignedTo)
}

func TestUserGetNotExists(t *testing.T) {
	ctx := context.Background()
	db := getDb(t)
	resetDb(t, db, ctx)
	ur := postgres.NewUserRepository(db.DB)

	user, err := ur.Get(1)
	assert.NotNil(t, err)
	assert.Nil(t, user)
}

func TestUserGetCompanyUsers(t *testing.T) {
	ctx := context.Background()
	db := getDb(t)
	resetDb(t, db, ctx)
	ur := postgres.NewUserRepository(db.DB)

	fixture := dbfixture.New(db)
	err := fixture.Load(ctx, os.DirFS("testdata"), "awards.yml")
	if err != nil {
		t.Errorf("could not load fixture: %v", err)
	}

	expCompID := awards.CompanyId(1)
	users, err := ur.GetCompanyUsers(expCompID)
	assert.Nil(t, err)
	assert.NotNil(t, users)
	assert.Equal(t, 2, len(users))
	assert.Equal(t, expCompID, users[0].CompanyId)
	assert.Equal(t, expCompID, users[1].CompanyId)
}
