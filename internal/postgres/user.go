package postgres

import (
	"database/sql"
	awards "gosanta/internal"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
)

type DBUser struct {
	bun.BaseModel `bun:"table:users"`

	ID        int64 `bun:",pk,autoincrement"`
	CompanyId int64
}

type userRepository struct {
	db *bun.DB
}

func NewUserRepository(sqlDb *sql.DB) *userRepository {
	db := bun.NewDB(sqlDb, pgdialect.New())
	return &userRepository{db: db}
}

func (ar *userRepository) Get(id awards.UserId) (*awards.User, error) {
	dbu := new(DBUser)
	err := ar.db.NewSelect().Model(dbu).Where("id = ?", id).Scan(nil)
	// TODO
	if err != nil {
		return nil, nil
	}

	pa := toUser(*dbu)
	return &pa, nil
}

func toUser(dbUser DBUser) awards.User {
	return awards.User{
		Id:        awards.UserId(dbUser.ID),
		CompanyId: awards.CompanyId(dbUser.CompanyId),
	}
}
