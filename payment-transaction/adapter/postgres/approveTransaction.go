package repository

import (
	"context"
	"log"
	"time"

	"github.com/renatospaka/payment-transaction/core/entity"
)

func (p *PostgresDatabase) approveTransaction(ctx context.Context, tr *entity.Transaction) error {
	log.Println("repository.transactions.approveTransaction")
	query := `
	UPDATE transactions
	SET status=$1,
			approved_at=$2,
			denied_at=$3,
			updated_at=$4,
			deleted_at=$5,
			authorization_id=$6
	WHERE id=$7
	`
	stmt, err := p.DB.PrepareContext(ctx, query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	var approvedAt, deniedAt, updatedAt, deletedAt interface{}	
	deniedAt = nil
	deletedAt = nil

	approvedAt = tr.ApprovedAt().Format(time.UnixDate)
	if tr.ApprovedAt().IsZero() {
		approvedAt = nil
	}

	updatedAt = tr.UpdatedAt().Format(time.UnixDate)
	if tr.UpdatedAt().IsZero() {
		updatedAt = nil
	}

	_, err = stmt.ExecContext(ctx, 
												tr.GetStatus(), 
												approvedAt, 
												deniedAt, 
												updatedAt, 
												deletedAt, 
												tr.GetAuthorizationID(),
												tr.GetID())
	if err != nil {
		return err
	}
	return nil
}
