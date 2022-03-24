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

const DSN = "postgres://%v:%v@%v:%v/%v?sslmode=disable"

type DBAward struct {
	bun.BaseModel `bun:"table:awards"`

	ID       int64     `bun:"id,pk,autoincrement"`
	UserID   int64     `bun:"user_id,notnull"`
	User     DBUser    `bun:"rel:belongs-to,join:user_id=id"`
	EarnedOn time.Time `bun:"earned_on,notnull"`
	Reason   int       `bun:"reason,notnull"`
	EmailRef string    `bun:"email_ref,notnull"`
}

type AwardRepository struct {
	db *bun.DB
}

func NewAwardRepository(sqlDb *sql.DB) *AwardRepository {
	db := bun.NewDB(sqlDb, pgdialect.New())
	db.RegisterModel((*DBAward)(nil))
	return &AwardRepository{db: db}
}

func (ar *AwardRepository) Get(id int64) (*awards.PhishingAward, error) {
	dba := new(DBAward)
	ctx := context.Background()
	err := ar.db.NewSelect().Model(dba).Where("id = ?", id).Scan(ctx)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, &awards.Error{
				Code: awards.DoesNotExistError,
				Err:  fmt.Errorf("no award with id: %v", id),
			}
		}
		return nil, err
	}

	pa := toAward(*dba)
	return &pa, nil
}

func (ar *AwardRepository) GetByUserId(id awards.UserId) ([]awards.PhishingAward, error) {
	var dbAwards []DBAward
	var awardS []awards.PhishingAward

	ctx := context.Background()
	err := ar.db.NewSelect().Model(&dbAwards).Relation("User").Where("user_id = ?", id).Scan(ctx)
	if err != nil {
		return awardS, err
	}
	if len(dbAwards) == 0 {
		return awardS, &awards.Error{
			Code: awards.DoesNotExistError,
			Err:  fmt.Errorf("no awards for user with id: %v", id),
		}
	}

	for _, a := range dbAwards {
		aw := toAward(a)
		awardS = append(awardS, aw)
	}
	return awardS, nil
}

func toAward(dbAward DBAward) awards.PhishingAward {
	return awards.PhishingAward{
		Id:         dbAward.ID,
		AssignedTo: awards.UserId(dbAward.UserID),
		EarnedOn:   dbAward.EarnedOn,
		Reason:     awards.AwardType(dbAward.Reason),
		EmailRef:   dbAward.EmailRef,
	}
}
