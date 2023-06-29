package repository

import (
	"context"
	"log"
	"time"

	"github.com/renatospaka/payment-transaction/core/entity"
)

func (p *PostgresDatabase) createTransaction(ctx context.Context, tr *entity.Transaction) error {
	log.Println("repository.transactions.createTransaction")

	var approvedAt, deniedAt, createdAt, updatedAt, deletedAt interface{}

	approvedAt = tr.ApprovedAt().Format(time.UnixDate)
	if tr.ApprovedAt().IsZero() {
		approvedAt = nil
	}

	deniedAt = tr.DeniedAt().Format(time.UnixDate)
	if tr.DeniedAt().IsZero() {
		deniedAt = nil
	}

	createdAt = tr.CreatedAt().Format(time.UnixDate)
	if tr.CreatedAt().IsZero() {
		createdAt = nil
	}

	updatedAt = tr.UpdatedAt().Format(time.UnixDate)
	if tr.UpdatedAt().IsZero() {
		updatedAt = nil
	}

	deletedAt = tr.DeletedAt().Format(time.UnixDate)
	if tr.DeletedAt().IsZero() {
		deletedAt = nil
	}

	query := `
	INSERT INTO transactions
		(id, client_id, status, value, approved_at, denied_at, created_at, updated_at, deleted_at) 
	VALUES
		($1, $2, $3, $4, $5, $6, $7, $8, $9)
	`
	stmt, err := p.DB.PrepareContext(ctx, query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.ExecContext(ctx,
		tr.GetID(),
		tr.GetClientID(),
		tr.GetStatus(),
		tr.GetValue(),
		approvedAt,
		deniedAt,
		createdAt,
		updatedAt,
		deletedAt,
	)
	if err != nil {
		return err
	}
	return nil
}
