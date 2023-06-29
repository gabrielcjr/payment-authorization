package repository

import (
	"context"
	"log"
)

func (p *PostgresDatabase) deleteTransaction(ctx context.Context, id string) error {
	log.Println("repository.transactions.deleteTransaction")

	query := `
	DELETE FROM transactions
	WHERE id=$1
	`
	stmt, err := p.DB.PrepareContext(ctx, query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.ExecContext(ctx, id)
	if err != nil {
		return err
	}
	return nil
}
