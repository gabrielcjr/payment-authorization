package usecase

import (
	"context"
	"errors"
	"log"
)

func (t *TransactionUsecase) deleteTransaction(ctx context.Context, id string) error {
	if id == "" {
		return errors.New("id is required")
	}
	log.Printf("Excluindo id %s\n", id)
	
	tr, err := t.repo.Find(id)
	if err != nil {
		return err
	} else
	if tr == nil {
		return errors.New("transaction id not found")
	}

	err = t.repo.Delete(id)
	if err != nil {
		return err
	}

	return nil
}
