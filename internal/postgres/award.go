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
	ctx := context.Background()
	query := ar.db.NewSelect().Where("id = ?", id)

	award, err := ar.getOne(ctx, query)
	if err != nil {
		return award, awards.ExtendError(
			err,
			fmt.Sprintf("no award with id: %v", id),
		)
	}

	return award, nil
}

func (ar *AwardRepository) GetUserAwards(id awards.UserId) ([]awards.PhishingAward, error) {
	ctx := context.Background()
	query := ar.db.NewSelect().Where("user_id = ?", id)

	awardS, err := ar.getList(ctx, query)
	if err != nil {
		return awardS, awards.ExtendError(err, fmt.Sprintf("could not retrieve for userId: %v", id))
	}

	return awardS, nil
}

func (ar *AwardRepository) GetByEmailRef(
	id awards.UserId,
	ref string,
) (*awards.PhishingAward, error) {
	ctx := context.Background()
	query := ar.db.NewSelect().Where("email_ref = ? AND user_id = ?", ref, id)

	award, err := ar.getOne(ctx, query)
	if err != nil {
		return award, awards.ExtendError(
			err,
			fmt.Sprintf("could not retrieve for userId: %v and emailRef: %v", id, ref),
		)
	}

	return award, nil
}

func (ar *AwardRepository) getOne(
	ctx context.Context,
	sQuery *bun.SelectQuery,
) (*awards.PhishingAward, error) {
	dba := new(DBAward)
	err := sQuery.Model(dba).Scan(ctx)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	pa := toAward(*dba)
	return &pa, nil
}

func (ar *AwardRepository) getList(
	ctx context.Context,
	sQuery *bun.SelectQuery,
) ([]awards.PhishingAward, error) {
	var awardS []awards.PhishingAward
	var dbAwards []DBAward

	err := sQuery.Model(&dbAwards).Scan(ctx)
	if err != nil {
		return awardS, err
	}
	if len(dbAwards) == 0 {
		return awardS, &awards.Error{
			Code: awards.DoesNotExistError,
			Err:  fmt.Errorf("no awards"),
		}
	}

	for _, a := range dbAwards {
		aw := toAward(a)
		awardS = append(awardS, aw)
	}
	return awardS, nil
}

func (ar *AwardRepository) Add(award *awards.PhishingAward) error {
	dbpe := &DBAward{
		UserID:   int64(award.AssignedTo),
		EarnedOn: award.EarnedOn,
		Reason:   int(award.Type),
		EmailRef: award.EmailRef,
	}
	ctx := context.Background()

	_, err := ar.db.NewInsert().Model(dbpe).Exec(ctx)
	if err != nil {
		return fmt.Errorf("could not insert PhishingAward for user: %v: %v", award.AssignedTo, err)
	}
	return nil
}

func (ar *AwardRepository) Delete(id int64) error {
	ctx := context.Background()
	_, err := ar.db.NewDelete().Model((*DBAward)(nil)).Where("id = ?", id).Exec(ctx)
	if err != nil {
		return fmt.Errorf("could not delete award with id: %v: %v", id, err)
	}
	return nil
}

func (ar *AwardRepository) UpdateExisting(existing, award *awards.PhishingAward) error {
	dbAward := &DBAward{
		ID:       existing.Id,
		EarnedOn: award.EarnedOn,
		EmailRef: award.EmailRef,
		Reason:   int(award.Type),
		UserID:   int64(award.AssignedTo),
	}
	ctx := context.Background()

	_, err := ar.db.NewUpdate().Model(dbAward).WherePK().Exec(ctx)
	if err != nil {
		return fmt.Errorf(
			"could not update existing award: %v with new award: %v: %v",
			existing,
			award,
			err,
		)
	}
	return nil
}

func toAward(dbAward DBAward) awards.PhishingAward {
	return awards.PhishingAward{
		Id:         dbAward.ID,
		AssignedTo: awards.UserId(dbAward.UserID),
		EarnedOn:   dbAward.EarnedOn,
		Type:       awards.AwardType(dbAward.Reason),
		EmailRef:   dbAward.EmailRef,
	}
}
