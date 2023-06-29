package repository

import (
	"context"
	"database/sql"
	"log"

	"github.com/renatospaka/payment-transaction/core/dto"
	"github.com/renatospaka/payment-transaction/core/entity"
	"github.com/renatospaka/payment-transaction/utils/dateTime"
	"github.com/renatospaka/payment-transaction/utils/strings"
	utils "github.com/renatospaka/payment-transaction/utils/entity"
)

func (p *PostgresDatabase) findTransaction(ctx context.Context, id string) (*entity.Transaction, error) {
	log.Println("repository.transactions.findTransaction")

	query := `
	SELECT id, client_id, authorization_id, status, value , approved_at, denied_at, created_at, updated_at, deleted_at
	FROM transactions
	WHERE id = $1`
	stmt, err := p.DB.PrepareContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows := stmt.QueryRow(id)
	if err != nil {
		return nil, err
	}

	var (
		tr            dto.TransactionDto
		approved      sql.NullTime
		denied        sql.NullTime
		created       sql.NullTime
		updated       sql.NullTime
		deleted       sql.NullTime
		client        sql.NullString
		authorization sql.NullString
	)
	err = rows.Scan(
		&tr.ID,
		&client,
		&authorization,
		&tr.Status,
		&tr.Value,
		&approved,
		&denied,
		&created,
		&updated,
		&deleted,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	approvedAt := dateTime.FormatNullDate(approved)
	deniedAt := dateTime.FormatNullDate(denied)
	createdAt := dateTime.FormatNullDate(created)
	updatedAt := dateTime.FormatNullDate(updated)
	deletedAt := dateTime.FormatNullDate(deleted)
	clientId, _ := utils.Parse(strings.FormatNullString(client))
	authorizationId, _ := utils.Parse(strings.FormatNullString(authorization))

	existing := &entity.TransactionMount{
		ID:              tr.ID,
		ClientID:        clientId.String(),
		AuthorizationID: authorizationId.String(),
		Value:           tr.Value,
		Status:          tr.Status,
		DeniedAt:        deniedAt,
		ApprovedAt:      approvedAt,
		TrailDate:       &utils.TrailDate{},
	}
	existing.SetCreationToDate(createdAt)
	existing.SetAlterationToDate(updatedAt)
	existing.SetDeletionToDate(deletedAt)

	transaction, err := entity.MountTransaction(existing)
	if err != nil || !transaction.IsValid() {
		return nil, err
	}
	return transaction, nil
}
