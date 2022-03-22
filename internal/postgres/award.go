package postgres

import (
	"database/sql"
	awards "gosanta/internal"
	"time"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
)

const DSN = "postgres://%v:%v@%v:%v/%v?sslmode=disable"

type DBAward struct {
	bun.BaseModel `bun:"table:awards"`

	ID         int64 `bun:",pk,autoincrement"`
	AssignedTo int64
	EarnedOn   time.Time
	Reason     int
	EmailRef   string
}

type awardRepository struct {
	db *bun.DB
}

func NewAwardRepository(sqlDb *sql.DB) *awardRepository {
	db := bun.NewDB(sqlDb, pgdialect.New())
	return &awardRepository{db: db}
}

func (ar *awardRepository) Get(id int64) (*awards.PhishingAward, error) {
	dba := new(DBAward)
	err := ar.db.NewSelect().Model(dba).Where("id = ?", id).Scan(nil)
	// TODO
	if err != nil {
		return nil, nil
	}

	pa := toAward(*dba)
	return &pa, nil
}

func (ar *awardRepository) GetByUserId(id awards.UserId) ([]awards.PhishingAward, error) {
	var dbAwards []DBAward
	err := ar.db.NewSelect().Model(&dbAwards).Where("AssignedTo = ?", id).Scan(nil)
	// TODO
	if err != nil {
		return nil, nil
	}

	var awards []awards.PhishingAward
	for _, a := range dbAwards {
		aw := toAward(a)
		awards = append(awards, aw)
	}
	return awards, nil
}

func toAward(dbAward DBAward) awards.PhishingAward {
	return awards.PhishingAward{
		Id:         dbAward.ID,
		AssignedTo: awards.UserId(dbAward.AssignedTo),
		EarnedOn:   dbAward.EarnedOn,
		Reason:     awards.AwardType(dbAward.Reason),
		EmailRef:   dbAward.EmailRef,
	}
}
