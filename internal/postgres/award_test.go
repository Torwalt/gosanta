package postgres_test

import (
	"context"
	"os"
	"testing"
	"time"

	awards "gosanta/internal"
	"gosanta/internal/postgres"

	"github.com/stretchr/testify/assert"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dbfixture"
	"github.com/uptrace/bun/dialect/pgdialect"
)

func getPostgresConf() postgres.Config {
	return postgres.Config{
		Host:   os.Getenv("POSTGRES_HOST"),
		Port:   os.Getenv("POSTGRES_PORT"),
		User:   os.Getenv("POSTGRES_USER"),
		Secret: os.Getenv("POSTGRES_SECRET"),
		Name:   os.Getenv("POSTGRES_NAME"),
	}
}

func getDb(t *testing.T) *bun.DB {
	conf := getPostgresConf()
	sqlDb := postgres.NewDb(conf)
	return bun.NewDB(sqlDb, pgdialect.New())
}

func resetDb(t *testing.T, db *bun.DB, ctx context.Context) {
	err := db.ResetModel(
		ctx,
		(*postgres.DBAward)(nil),
		(*postgres.DBUser)(nil),
		(*postgres.DBPhishingEvent)(nil),
	)
	if err != nil {
		t.Errorf("could not reset DB: %v", err)
	}
}

func TestAwardGet(t *testing.T) {
	ctx := context.Background()
	db := getDb(t)
	resetDb(t, db, ctx)
	ar := postgres.NewAwardRepository(db.DB)

	fixture := dbfixture.New(db, dbfixture.WithRecreateTables())
	err := fixture.Load(ctx, os.DirFS("testdata"), "awards.yml")
	if err != nil {
		t.Errorf("could not load fixture: %v", err)
	}

	a, err := ar.Get(1)
	assert.Nil(t, err)
	assert.NotNil(t, a)

	assert.Equal(t, int64(1), a.Id)
	assert.Equal(t, awards.OpenAward, a.Type)
	assert.Equal(t, awards.UserId(1), a.AssignedTo)
}

func TestAwardGetNotExists(t *testing.T) {
	ctx := context.Background()
	db := getDb(t)
	resetDb(t, db, ctx)
	ar := postgres.NewAwardRepository(db.DB)

	a, err := ar.Get(1)
	assert.Nil(t, err)
	assert.Nil(t, a)
}

func TestAwardGetUserAwards(t *testing.T) {
	ctx := context.Background()
	db := getDb(t)
	resetDb(t, db, ctx)
	ar := postgres.NewAwardRepository(db.DB)

	fixture := dbfixture.New(db)
	err := fixture.Load(ctx, os.DirFS("testdata"), "awards.yml")
	if err != nil {
		t.Errorf("could not load fixture: %v", err)
	}

	uId := awards.UserId(1)
	awards, err := ar.GetUserAwards(uId)
	assert.Nil(t, err)
	assert.NotNil(t, awards)

	assert.Equal(t, 2, len(awards))
	assert.Equal(t, awards[0].AssignedTo, uId)
	assert.Equal(t, awards[1].AssignedTo, uId)
}

func TestAwardGetByUserIdNotExists(t *testing.T) {
	ctx := context.Background()
	db := getDb(t)
	resetDb(t, db, ctx)
	ar := postgres.NewAwardRepository(db.DB)

	uId := awards.UserId(1)
	awards, err := ar.GetUserAwards(uId)
	assert.Nil(t, awards)
	assert.NotNil(t, err)
}

func TestAwardAdd(t *testing.T) {
	ctx := context.Background()
	db := getDb(t)
	resetDb(t, db, ctx)
	ar := postgres.NewAwardRepository(db.DB)

	fixture := dbfixture.New(db)
	err := fixture.Load(ctx, os.DirFS("testdata"), "awards.yml")
	if err != nil {
		t.Errorf("could not load fixture: %v", err)
	}

	pa := &awards.PhishingAward{
		AssignedTo: awards.UserId(1),
		EarnedOn:   time.Now(),
		Type:       awards.OpenAward,
		EmailRef:   "f20416ef-15d5-4159-9bef-de150edfa970",
	}
	err = ar.Add(pa)

	assert.Nil(t, err)
}

func TestAwardDelete(t *testing.T) {
	ctx := context.Background()
	db := getDb(t)
	resetDb(t, db, ctx)
	ar := postgres.NewAwardRepository(db.DB)

	fixture := dbfixture.New(db, dbfixture.WithRecreateTables())
	err := fixture.Load(ctx, os.DirFS("testdata"), "awards.yml")
	if err != nil {
		t.Errorf("could not load fixture: %v", err)
	}

	err = ar.Delete(1)
	assert.Nil(t, err)

	a, err := ar.Get(1)
	assert.Nil(t, err)
	assert.Nil(t, a)
}

func TestAwardGetByEmailRef(t *testing.T) {
	ctx := context.Background()
	db := getDb(t)
	resetDb(t, db, ctx)
	ar := postgres.NewAwardRepository(db.DB)

	fixture := dbfixture.New(db, dbfixture.WithRecreateTables())
	err := fixture.Load(ctx, os.DirFS("testdata"), "awards.yml")
	if err != nil {
		t.Errorf("could not load fixture: %v", err)
	}

	a, err := ar.GetByEmailRef(1, "f20416ef-15d5-4159-9bef-de150edfa970")
	assert.Nil(t, err)
	assert.NotNil(t, a)

	assert.Equal(t, int64(1), a.Id)
	assert.Equal(t, awards.OpenAward, a.Type)
	assert.Equal(t, awards.UserId(1), a.AssignedTo)
	assert.Equal(t, "f20416ef-15d5-4159-9bef-de150edfa970", a.EmailRef)
}

func TestAwardUpdateExisting(t *testing.T) {
	ctx := context.Background()
	db := getDb(t)
	resetDb(t, db, ctx)
	ar := postgres.NewAwardRepository(db.DB)

	fixture := dbfixture.New(db, dbfixture.WithRecreateTables())
	err := fixture.Load(ctx, os.DirFS("testdata"), "awards.yml")
	if err != nil {
		t.Errorf("could not load fixture: %v", err)
	}

	existingAward := &awards.PhishingAward{
		Id:         1,
		AssignedTo: awards.UserId(1),
		EarnedOn:   time.Now(),
		Type:       awards.OpenAward,
		EmailRef:   "f20416ef-15d5-4159-9bef-de150edfa970",
	}
	newAward := &awards.PhishingAward{
		Id:         1,
		AssignedTo: awards.UserId(1),
		EarnedOn:   time.Now().Add(time.Duration(100)),
		Type:       awards.ReportAward,
		EmailRef:   "f20416ef-15d5-4159-9bef-de150edfa970",
	}
	err = ar.UpdateExisting(existingAward, newAward)

	assert.Nil(t, err)
}
