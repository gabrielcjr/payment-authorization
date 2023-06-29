package usecase

import (
	"context"
	"errors"
	"log"

	"github.com/renatospaka/payment-transaction/core/dto"
	utils "github.com/renatospaka/payment-transaction/utils/entity"
)

func (t *TransactionUsecase) updateTransaction(ctx context.Context, tr *dto.TransactionUpdateDto) error {
	log.Println("usecase.transactions.update")
	if tr.ID == "" {
		return errors.New("id is required")
	}

	if _, err := utils.Parse(tr.ID); err != nil {
		return errors.New("invalid id")
	}

	transaction, err := t.repo.Find(tr.ID)
	if err != nil {
		return err
	} 
	if transaction == nil {
		return errors.New("transaction id not found")
	}

	// the only attribute allowed to be changed is the value. 
	// if the new value is equal to the current value, then an error is raised
	value := transaction.GetValue()
	if value == tr.Value {
		return errors.New("nothing to update in the transaction")
	}

	transaction.SetValue(tr.Value)
	if !transaction.IsValid() {
		return errors.New("transaction is not validated")
	}

	return t.repo.Update(transaction)
}
