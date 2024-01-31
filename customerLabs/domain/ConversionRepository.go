package domain

import (
	"github.com/jmoiron/sqlx"
)

type GetDbClient func() *sqlx.DB
type RepositoryDb struct {
	client GetDbClient
}

func NewConversionRepositoryDb() RepositoryDb {
	return RepositoryDb{}
}
