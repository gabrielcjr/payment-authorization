package entity

import (
	"time"

	pkgEntity "github.com/renatospaka/payment-transaction/utils/entity"
)

const (
	TR_APPROVED = "approved"
	TR_DENIED   = "denied"
	TR_PENDING  = "pending"
	TR_DELETED  = "deleted"
)

type Transaction struct {
	deniedAt   time.Time
	approvedAt time.Time
	*pkgEntity.TrailDate
	id              pkgEntity.ID
	clientId        pkgEntity.ID
	authorizationId pkgEntity.ID
	status          string
	value           float32
	valid           bool
}

type TransactionMount struct {
	ID              string
	ClientID        string
	AuthorizationID string
	Value           float32
	Status          string
	DeniedAt        time.Time
	ApprovedAt      time.Time
	*pkgEntity.TrailDate
}

// Create a new transaction
func NewTransaction(clientId string, value float32) (*Transaction, error) {
	uuid, err := pkgEntity.Parse(clientId)
	if err != nil {
		return nil, ErrInvalidClientID
	}

	transaction := &Transaction{
		id:         pkgEntity.NewID(),
		clientId:   uuid,
		value:      value,
		status:     TR_PENDING,
		deniedAt:   time.Time{},
		approvedAt: time.Time{},
		TrailDate:  &pkgEntity.TrailDate{},
		valid:      false,
	}
	transaction.TrailDate.SetCreationToToday()

	// deliver the new transaction validated
	err = transaction.Validate()
	if err != nil {
		return nil, err
	}
	return transaction, nil
}

// func MountTransaction(id string, client string, authorizationId string, status string, value float32, deniedAt time.Time, approvedAt time.Time, createdAt time.Time, updatedAt time.Time, deletedAt time.Time) (*Transaction, error) {
func MountTransaction(mounting *TransactionMount) (*Transaction, error) {
	uuid, err := pkgEntity.Parse(mounting.ID)
	if err != nil {
		return nil, err
	}

	uuidClient, err := pkgEntity.Parse(string(mounting.ClientID))
	if err != nil {
		return nil, err
	}

	uuidCAuthorization, err := pkgEntity.Parse(mounting.AuthorizationID)
	if err != nil {
		return nil, err
	}

	status := mounting.Status
	if status!= TR_APPROVED &&
		status!= TR_DELETED &&
		status!= TR_DENIED &&
		status!= TR_PENDING {
		status = TR_PENDING
	}

	transaction := &Transaction{
		id:              uuid,
		clientId:        uuidClient,
		authorizationId: uuidCAuthorization,
		value:           mounting.Value,
		status:          status,
		deniedAt:        mounting.DeniedAt,
		approvedAt:      mounting.ApprovedAt,
		TrailDate:       &pkgEntity.TrailDate{},
		valid:           false,
	}
	transaction.TrailDate.SetCreationToDate(mounting.CreatedAt())
	transaction.TrailDate.SetAlterationToToday()
	if !mounting.DeletedAt().IsZero() {
		transaction.TrailDate.SetDeletionToDate(mounting.DeletedAt())
	}

	// deliver the new transaction validated
	err = transaction.Validate()
	if err != nil {
		return nil, err
	}
	return transaction, nil
}

// Get the ID of the transaction
func (t *Transaction) GetID() string {
	uuid := t.id.String()
	if uuid == "00000000-0000-0000-0000-000000000000" {
		uuid = ""
	}
	return uuid
}

// Get the Client ID of the transaction
func (t *Transaction) GetClientID() string {
	uuid := t.clientId.String()
	if uuid == "00000000-0000-0000-0000-000000000000" {
		uuid = ""
	}
	return uuid
}

// Get the Authorization ID of the transaction
func (t *Transaction) GetAuthorizationID() string {
	uuid := t.authorizationId.String()
	if uuid == "00000000-0000-0000-0000-000000000000" {
		uuid = ""
	}
	return uuid
}

// Change the status to approved and the approved day is today
func (t *Transaction) SetStatusToApproved(authorization string) error {
	uuid, err := pkgEntity.Parse(authorization)
	if err != nil {
		return ErrInvalidAuthorizationID
	}

	t.authorizationId = uuid
	t.approvedAt = time.Now()
	t.deniedAt = time.Time{}
	t.TrailDate.SetAlterationToToday()
	t.status = TR_APPROVED
	return nil
}

// Change the status to approved on a specific date
func (t *Transaction) SetStatusToApprovedOnDate(authorization string, date time.Time) error {
	uuid, err := pkgEntity.Parse(authorization)
	if err != nil {
		return ErrInvalidAuthorizationID
	}

	t.authorizationId = uuid
	t.approvedAt = date
	t.deniedAt = time.Time{}
	t.TrailDate.SetAlterationToToday()
	t.status = TR_APPROVED
	return nil
}

// Change the status to denied and the denied day is today
func (t *Transaction) SetStatusToDenied(authorization string) error {
	uuid, err := pkgEntity.Parse(authorization)
	if err != nil {
		return ErrInvalidAuthorizationID
	}

	t.authorizationId = uuid
	t.approvedAt = time.Time{}
	t.deniedAt = time.Now()
	t.TrailDate.SetAlterationToToday()
	t.status = TR_DENIED
	return nil
}

// Change the status to denied on a specific date
func (t *Transaction) SetStatusToDeniedOnDate(authorization string, date time.Time) error {
	uuid, err := pkgEntity.Parse(authorization)
	if err != nil {
		return ErrInvalidAuthorizationID
	}

	t.authorizationId = uuid
	t.approvedAt = time.Time{}
	t.deniedAt = date
	t.TrailDate.SetAlterationToToday()
	t.status = TR_DENIED
	return nil
}

// Change the status to deleted and the deleted day is today
func (t *Transaction) SetStatusToDeleted() {
	t.approvedAt = time.Time{}
	t.deniedAt = time.Time{}
	t.TrailDate.SetDeletionToToday()
	t.status = TR_DELETED
}

// Get the current status of the transacion
func (t *Transaction) GetStatus() string {
	return t.status
}

// Get when the transacion was denied (if it was)
func (t *Transaction) DeniedAt() time.Time {
	return t.deniedAt
}

// Get when the transacion was approved (if it was)
func (t *Transaction) ApprovedAt() time.Time {
	return t.approvedAt
}

// Change the transaction value and validate it before committing it
func (t *Transaction) SetValue(value float32) (err error) {
	current := t.value
	
	t.value = value
	if err = t.Validate(); err != nil {
		t.value = current
		return err
	}
	t.SetAlterationToToday()
	
	return nil
}

// Get the current value of the transaction
func (t *Transaction) GetValue() float32 {
	return t.value
}

// Validates all business rules for this transaction
func (t *Transaction) Validate() error {
	t.valid = false
	
	if t.id.String() == "" {
		return ErrIDIsRequired
	}

	if _, err := pkgEntity.Parse(t.id.String()); err != nil {
		return ErrInvalidID
	}

	if t.clientId.String() == "" {
		return ErrIDIsRequired
	}

	if _, err := pkgEntity.Parse(t.clientId.String()); err != nil {
		return ErrInvalidClientID
	}

	if t.value < 0 {
		return ErrValueIsNegative
	}

	if t.value == 0 {
		return ErrValueIsZero
	}

	if t.status != TR_APPROVED &&
		t.status != TR_DELETED &&
		t.status != TR_DENIED &&
		t.status != TR_PENDING {
		return ErrInvalidStatus
	}

	t.valid = true
	return nil
}

// Return whether the structure is valid or not
func (t *Transaction) IsValid() bool {
	return t.valid
}
