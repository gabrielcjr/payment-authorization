package repository

import (
	"context"
	"log"

	"github.com/renatospaka/payment-transaction/core/entity"
)

func (p *PostgresDatabase) updateTransaction(ctx context.Context, tr *entity.Transaction) error {
	log.Println("repository.transactions.updateTransaction")

	query := `
	UPDATE transactions
	SET value=$2,
			updated_at=$3
	WHERE id = $1`
	stmt, err := p.DB.PrepareContext(ctx, query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_ = stmt.QueryRow(tr.GetID(), tr.GetValue(), tr.UpdatedAt())
	if err != nil {
		return err
	}	
	return nil
}
