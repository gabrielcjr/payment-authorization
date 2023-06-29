package usecase

import (
	"context"
	"errors"
	"log"

	"github.com/renatospaka/payment-transaction/core/dto"
	"github.com/renatospaka/payment-transaction/utils/dateTime"
)

func (t *TransactionUsecase) findAllTransactions(ctx context.Context, page int, limit int) (dto.TransactionFindAllResponseDto, error) {
	log.Println("usecase.transactions.findAll")

	if page < 0 {
		return dto.TransactionFindAllResponseDto{}, errors.New("pagination cannot be negative")
	} else if limit < 0 {
		return dto.TransactionFindAllResponseDto{}, errors.New("limitation cannot be negative")
	}

	transactions, err := t.repo.FindAll(page, limit)
	if err != nil {
		return dto.TransactionFindAllResponseDto{}, err
	}

	var response dto.TransactionFindAllResponseDto
	for _, tr := range transactions {
		transaction := &dto.TransactionDto{
			ID:              tr.GetID(),
			ClientID:        tr.GetClientID(),
			AuthorizationID: tr.GetAuthorizationID(),
			Status:          tr.GetStatus(),
			Value:           tr.GetValue(),
			DeniedAt:        dateTime.FormatDateToNull(tr.DeniedAt()),
			ApprovedAt:      dateTime.FormatDateToNull(tr.ApprovedAt()),
			CreatedAt:       dateTime.FormatDateToNull(tr.CreatedAt()),
			UpdatedAt:       dateTime.FormatDateToNull(tr.UpdatedAt()),
			DeletedAt:       dateTime.FormatDateToNull(tr.DeletedAt()),
		}
		response.Transactions = append(response.Transactions, transaction)
	}
	return response, nil
}
