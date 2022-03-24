package postgres_test

import (
	"context"
	awards "gosanta/internal"
	"gosanta/internal/postgres"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/uptrace/bun/dbfixture"
)

func TestPERGetUnprocessed(t *testing.T) {
	ctx := context.Background()
	db := getDb(t)
	resetDb(t, db, ctx)
	per := postgres.NewPhishingEventRepository(db.DB)

	fixture := dbfixture.New(db, dbfixture.WithRecreateTables())
	err := fixture.Load(ctx, os.DirFS("testdata"), "events.yml")
	if err != nil {
		t.Errorf("could not load fixture: %v", err)
	}

	pEvents, err := per.GetUnprocessed()

	assert.Nil(t, err)
	assert.NotNil(t, pEvents)
	assert.Equal(t, 1, len(pEvents))
	assert.Equal(t, awards.Opened, pEvents[0].Action)
	assert.Nil(t, pEvents[0].ProcessedAt)
	assert.Equal(t, awards.UserId(1), pEvents[0].UserID)
}

func TestPERWrite(t *testing.T) {
	ctx := context.Background()
	db := getDb(t)
	resetDb(t, db, ctx)
	per := postgres.NewPhishingEventRepository(db.DB)

	fixture := dbfixture.New(db, dbfixture.WithRecreateTables())
	err := fixture.Load(ctx, os.DirFS("testdata"), "events.yml")
	if err != nil {
		t.Errorf("could not load fixture: %v", err)
	}

	upe := awards.UserPhishingEvent{
		UserID:      awards.UserId(1),
		Action:      awards.Opened,
		CreatedAt:   time.Now(),
		EmailRef:    "f20416ef-15d5-4159-9bef-de150edfa970",
		ProcessedAt: nil,
	}
	err = per.Write(upe)

	assert.Nil(t, err)
}
