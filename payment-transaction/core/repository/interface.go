package repository

import (
	"github.com/renatospaka/payment-transaction/core/entity"
)

type TransactionInterface interface {
	Create(*entity.Transaction) error
	Delete(string) error
	Update(*entity.Transaction) error
	Find(string) (*entity.Transaction, error)
	FindAll(page, limit int) ([]*entity.Transaction, error)
	Approve(*entity.Transaction) error
	Deny(*entity.Transaction) error
}
