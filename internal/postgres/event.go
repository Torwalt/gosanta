package postgres

import (
	"context"
	"database/sql"
	"fmt"
	awards "gosanta/internal"
	"time"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
)

type DBPhishingEvent struct {
	bun.BaseModel `bun:"table:phishing_events"`

	ID          int64      `bun:"id,pk,autoincrement"`
	UserID      int64      `bun:"user_id,notnull"`
	User        DBUser     `bun:"rel:belongs-to,join:user_id=id"`
	Action      int        `bun:"action,notnull"`
	CreatedAt   time.Time  `bun:"created_at,notnull"`
	EmailRef    string     `bun:"email_ref,notnull"`
	ProcessedAt *time.Time `bun:"processed_at"`
}

type PhishingEventRepository struct {
	db *bun.DB
}

func NewPhishingEventRepository(sqlDb *sql.DB) *PhishingEventRepository {
	db := bun.NewDB(sqlDb, pgdialect.New())
	db.RegisterModel((*DBPhishingEvent)(nil))
	return &PhishingEventRepository{db: db}
}

func (per *PhishingEventRepository) GetUnprocessed() ([]awards.UserPhishingEvent, error) {
	var dbPhishingEvents []DBPhishingEvent
	var phishingEvents []awards.UserPhishingEvent

	ctx := context.Background()
	err := per.db.NewSelect().Model(&dbPhishingEvents).Where("processed_at is Null").Scan(ctx)
	if err != nil {
		return phishingEvents, fmt.Errorf("could not retrieve unprocessed phishing events: %v", err)
	}

	for _, event := range dbPhishingEvents {
		pe := toPhishingEvent(event)
		phishingEvents = append(phishingEvents, pe)
	}
	return phishingEvents, nil
}

func (per *PhishingEventRepository) Write(upe awards.UserPhishingEvent) error {
	dbpe := &DBPhishingEvent{
		UserID:    int64(upe.UserID),
		EmailRef:  upe.EmailRef,
		CreatedAt: upe.CreatedAt,
		Action:    int(upe.Action),
	}
	ctx := context.Background()
	_, err := per.db.NewInsert().Model(dbpe).Exec(ctx)
	if err != nil {
		return fmt.Errorf("could not insert UserPhishingEvent for user: %v: %v", upe.UserID, err)
	}
	return nil
}

func toPhishingEvent(dbpe DBPhishingEvent) awards.UserPhishingEvent {
	return awards.UserPhishingEvent{
		ID:          dbpe.ID,
		UserID:      awards.UserId(dbpe.UserID),
		Action:      awards.PhishingAction(dbpe.Action),
		CreatedAt:   dbpe.CreatedAt,
		EmailRef:    dbpe.EmailRef,
		ProcessedAt: dbpe.ProcessedAt,
	}
}
