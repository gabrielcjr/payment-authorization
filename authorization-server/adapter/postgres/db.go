package repository

import (
	"context"
	"database/sql"

	_ "github.com/lib/pq"
	"github.com/renatospaka/authorization-server/core/entity"
)

type PostgresDatabase struct {
	DB *sql.DB
}

func NewPostgresDatabase(db *sql.DB) *PostgresDatabase {
	return &PostgresDatabase{
		DB: db,
	}
}

func (p *PostgresDatabase) Process(auth *entity.Authorization) error {
	ctx := context.Background()
	return p.processAuthorization(ctx, auth)
}
