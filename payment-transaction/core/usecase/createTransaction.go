package usecase

import (
	"context"
	"log"

	"github.com/renatospaka/payment-transaction/adapter/grpc/pb"
	"github.com/renatospaka/payment-transaction/core/dto"
	"github.com/renatospaka/payment-transaction/core/entity"
)

func (t *TransactionUsecase) createTransaction(ctx context.Context, tr *dto.TransactionCreateDto) (*dto.TransactionDto, error) {
	log.Println("usecase.transactions.create")

	transaction, err := entity.NewTransaction(tr.ClientID, tr.Value)
	if err != nil {
		return nil, err
	}
	id := transaction.GetID()
	clientId := transaction.GetClientID()
	value := transaction.GetValue()
	tr.ID = id

	err = t.repo.Create(transaction)
	if err != nil {
		return nil, err
	}

	// Call the Authotization service and update status accordingly
	auth := &pb.AuthorizationRequest{
		ClientId:      clientId,
		TransactionId: id,
		Value:         value,
	}
	authorize, err := t.authorizeTransaction(ctx, auth)
	if err != nil {
		return nil, err
	}

	// request the gRPC server to authorize the transaction
	tran := &dto.TransactionAuthorizeDto{
		ID:              id,
		ClientID:        clientId,
		AuthorizationId: authorize.AuthorizationId,
		Value:           value,
	}
	if authorize.Status == entity.TR_APPROVED {
		err = t.approveTransaction(ctx, tran)
	} else 
	if authorize.Status == entity.TR_DENIED {
		err = t.denyTransaction(ctx, tran)
	}
	if err != nil {
		return nil, err
	}

	// Find the new Transaction the reply to the caller
	created, err := t.findTransaction(ctx, id)
	if err != nil {
		return nil, err
	}
	return created, nil
}
