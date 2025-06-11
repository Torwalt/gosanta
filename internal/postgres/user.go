package postgres

import (
	"context"
	"database/sql"
	"fmt"

	awards "gosanta/internal"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
)

type DBUser struct {
	bun.BaseModel `bun:"table:users"`

	ID        int64      `bun:"id,pk,autoincrement"`
	CompanyId int64      `bun:"company_id,notnull"`
	Awards    []*DBAward `bun:"rel:has-many,join:id=user_id"`
}

type userRepository struct {
	db *bun.DB
}

func NewUserRepository(sqlDb *sql.DB) *userRepository {
	db := bun.NewDB(sqlDb, pgdialect.New())
	db.RegisterModel((*DBUser)(nil))
	return &userRepository{db: db}
}

func (ar *userRepository) Get(id awards.UserId) (*awards.User, error) {
	dbu := new(DBUser)
	ctx := context.Background()
	err := ar.db.NewSelect().Model(dbu).Relation("Awards").Where("id = ?", id).Scan(ctx)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, &awards.Error{
				Code: awards.DoesNotExistError,
				Err:  fmt.Errorf("no user with id: %v", id),
			}
		}
		return nil, err
	}

	pa := toUser(*dbu)
	return &pa, nil
}

func (ar *userRepository) GetCompanyUsers(cId awards.CompanyId) ([]awards.User, error) {
	var dbUsers []DBUser
	var users []awards.User

	ctx := context.Background()
	err := ar.db.NewSelect().
		Model(&dbUsers).
		Relation("Awards").
		Where("company_id = ?", cId).
		Scan(ctx)
	if err != nil {
		if err == sql.ErrNoRows {
			return users, &awards.Error{
				Code: awards.DoesNotExistError,
				Err:  fmt.Errorf("no user with id: %v", cId),
			}
		}
		return users, err
	}

	for _, dbUser := range dbUsers {
		users = append(users, toUser(dbUser))
	}
	return users, nil
}

func toUser(dbUser DBUser) awards.User {
	var awardS []awards.PhishingAward
	for _, dba := range dbUser.Awards {
		awardS = append(awardS, toAward(*dba))
	}

	return awards.User{
		Id:        awards.UserId(dbUser.ID),
		CompanyId: awards.CompanyId(dbUser.CompanyId),
		Awards:    awardS,
	}
}
