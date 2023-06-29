package usecase

import (
	"context"
	"log"

	"github.com/renatospaka/payment-transaction/core/dto"
	"github.com/renatospaka/payment-transaction/core/entity"
	pkgTrail "github.com/renatospaka/payment-transaction/utils/entity"
)

func (t *TransactionUsecase) denyTransaction(tx context.Context, tr *dto.TransactionAuthorizeDto) error {
	log.Println("usecase.transactions.denyTransaction")

	mounting := &entity.TransactionMount{
		ID:              tr.ID,
		ClientID:        tr.ClientID,
		AuthorizationID: tr.AuthorizationId,
		Value:           tr.Value,
		TrailDate:       &pkgTrail.TrailDate{},
	}
	transaction, err := entity.MountTransaction(mounting)
	if err != nil {
		return err
	}
	transaction.SetStatusToDenied(tr.AuthorizationId)

	return t.repo.Deny(transaction)
}
