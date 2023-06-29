package test

import (
	"testing"
	"time"

	"github.com/renatospaka/payment-transaction/core/entity"
	"github.com/stretchr/testify/assert"
)

const (
	AUTHORIZATION_ID = "6993f877-f0f3-4713-abbd-7ffc40be4eaf"
	CLIENT_ID        = "ebf77a6e-e236-441a-8ef6-4a3afdc5d4b1"
)

func TestNewTransaction(t *testing.T) {
	tr, err := entity.NewTransaction(CLIENT_ID, 200.00)
	assert.Nil(t, err)
	assert.NotNil(t, tr)
	assert.NotEmpty(t, tr.GetID())
	assert.Equal(t, CLIENT_ID, tr.GetClientID())
	assert.Equal(t, float32(200.00), tr.GetValue())
	assert.True(t, tr.DeniedAt().IsZero())
	assert.True(t, tr.ApprovedAt().IsZero())
	assert.Equal(t, entity.TR_PENDING, tr.GetStatus())

	err = tr.Validate()
	assert.Nil(t, err)
}

func TestNewTransactionInvalidlientID(t *testing.T) {
	tr, err := entity.NewTransaction("X"+CLIENT_ID, 0)
	assert.NotNil(t, err)
	assert.Nil(t, tr)
	assert.EqualError(t, err, entity.ErrInvalidClientID.Error())

	tr2, err2 := entity.NewTransaction("", 0)
	assert.NotNil(t, err2)
	assert.Nil(t, tr2)
	assert.EqualError(t, err2, entity.ErrInvalidClientID.Error())
}

func TestNewTransactionValueIsZero(t *testing.T) {
	tr, err := entity.NewTransaction(CLIENT_ID, 0)
	assert.NotNil(t, err)
	assert.Nil(t, tr)
	assert.EqualError(t, err, entity.ErrValueIsZero.Error())
}

func TestNewTransactionValueIsNegative(t *testing.T) {
	tr, err := entity.NewTransaction(CLIENT_ID, -200.00)
	assert.NotNil(t, err)
	assert.Nil(t, tr)
	assert.EqualError(t, err, entity.ErrValueIsNegative.Error())
}

func TestSetValue(t *testing.T) {
	tr, err := entity.NewTransaction(CLIENT_ID, 200.00)
	assert.Nil(t, err)
	assert.NotNil(t, tr)

	err = tr.SetValue(100)
	assert.Nil(t, err)
	assert.Equal(t, float32(100), tr.GetValue())

	err = tr.SetValue(0)
	assert.NotNil(t, err)
	assert.Equal(t, float32(100), tr.GetValue())

	err = tr.SetValue(-200)
	assert.NotNil(t, err)
	assert.Equal(t, float32(100), tr.GetValue())

	err = tr.Validate()
	assert.Nil(t, err)
}

func TestSetStatusToApproved(t *testing.T) {
	tr, err := entity.NewTransaction(CLIENT_ID, 200.00)
	assert.Nil(t, err)
	assert.NotNil(t, tr)
	assert.Equal(t, entity.TR_PENDING, tr.GetStatus())
	assert.Equal(t, "", tr.GetAuthorizationID())
	assert.True(t, tr.DeniedAt().IsZero())
	assert.True(t, tr.ApprovedAt().IsZero())

	err = tr.SetStatusToApproved(AUTHORIZATION_ID)
	assert.Nil(t, err)
	assert.Equal(t, entity.TR_APPROVED, tr.GetStatus())
	assert.Equal(t, AUTHORIZATION_ID, tr.GetAuthorizationID())
	assert.True(t, tr.DeniedAt().IsZero())
	assert.False(t, tr.ApprovedAt().IsZero())

	tr2, _ := entity.NewTransaction(CLIENT_ID, 200.00)
	err = tr2.SetStatusToApproved("")
	assert.NotNil(t, err)
	assert.Equal(t, entity.TR_PENDING, tr2.GetStatus())
	assert.Equal(t, "", tr2.GetAuthorizationID())
}

func TestSetStatusToApprovedOnDate(t *testing.T) {
	tr, err := entity.NewTransaction(CLIENT_ID, 200.00)
	assert.Nil(t, err)
	assert.NotNil(t, tr)
	assert.Equal(t, entity.TR_PENDING, tr.GetStatus())
	assert.Equal(t, "", tr.GetAuthorizationID())
	assert.True(t, tr.DeniedAt().IsZero())
	assert.True(t, tr.ApprovedAt().IsZero())

	date := time.Now().Add(time.Hour * 24 * -10)
	err = tr.SetStatusToApprovedOnDate(AUTHORIZATION_ID, date)
	assert.Nil(t, err)
	assert.Equal(t, entity.TR_APPROVED, tr.GetStatus())
	assert.Equal(t, AUTHORIZATION_ID, tr.GetAuthorizationID())
	assert.True(t, tr.DeniedAt().IsZero())
	assert.False(t, tr.ApprovedAt().IsZero())
	assert.Equal(t, date.Unix(), tr.ApprovedAt().Unix())

	tr2, _ := entity.NewTransaction(CLIENT_ID, 200.00)
	err = tr2.SetStatusToApprovedOnDate("", date)
	assert.NotNil(t, err)
	assert.Equal(t, entity.TR_PENDING, tr2.GetStatus())
	assert.Equal(t, "", tr2.GetAuthorizationID())
}

func TestSetStatusToDenied(t *testing.T) {
	tr, err := entity.NewTransaction(CLIENT_ID, 200.00)
	assert.Nil(t, err)
	assert.NotNil(t, tr)
	assert.Equal(t, entity.TR_PENDING, tr.GetStatus())
	assert.Equal(t, "", tr.GetAuthorizationID())
	assert.True(t, tr.DeniedAt().IsZero())
	assert.True(t, tr.ApprovedAt().IsZero())

	err = tr.SetStatusToDenied(AUTHORIZATION_ID)
	assert.Nil(t, err)
	assert.Equal(t, entity.TR_DENIED, tr.GetStatus())
	assert.Equal(t, AUTHORIZATION_ID, tr.GetAuthorizationID())
	assert.False(t, tr.DeniedAt().IsZero())
	assert.True(t, tr.ApprovedAt().IsZero())

	tr2, _ := entity.NewTransaction(CLIENT_ID, 200.00)
	err = tr2.SetStatusToDenied("")
	assert.NotNil(t, err)
	assert.Equal(t, entity.TR_PENDING, tr2.GetStatus())
	assert.Equal(t, "", tr2.GetAuthorizationID())
}

func TestSetStatusToDeniedOnDate(t *testing.T) {
	tr, err := entity.NewTransaction(CLIENT_ID, 200.00)
	assert.Nil(t, err)
	assert.NotNil(t, tr)
	assert.Equal(t, entity.TR_PENDING, tr.GetStatus())
	assert.Equal(t, "", tr.GetAuthorizationID())
	assert.True(t, tr.DeniedAt().IsZero())
	assert.True(t, tr.ApprovedAt().IsZero())

	date := time.Now().Add(time.Hour * 24 * -10)
	err = tr.SetStatusToDeniedOnDate(AUTHORIZATION_ID, date)
	assert.Nil(t, err)
	assert.Equal(t, entity.TR_DENIED, tr.GetStatus())
	assert.Equal(t, AUTHORIZATION_ID, tr.GetAuthorizationID())
	assert.False(t, tr.DeniedAt().IsZero())
	assert.True(t, tr.ApprovedAt().IsZero())
	assert.Equal(t, date.Unix(), tr.DeniedAt().Unix())

	tr2, _ := entity.NewTransaction(CLIENT_ID, 200.00)
	err = tr2.SetStatusToDeniedOnDate("", date)
	assert.NotNil(t, err)
	assert.Equal(t, entity.TR_PENDING, tr2.GetStatus())
	assert.Equal(t, "", tr2.GetAuthorizationID())
}

func TestSetStatusToDeleted(t *testing.T) {
	tr, err := entity.NewTransaction(CLIENT_ID, 200.00)
	assert.Nil(t, err)
	assert.NotNil(t, tr)
	assert.Equal(t, entity.TR_PENDING, tr.GetStatus())
	assert.True(t, tr.DeniedAt().IsZero())
	assert.True(t, tr.ApprovedAt().IsZero())

	tr.SetStatusToDeleted()
	assert.Equal(t, entity.TR_DELETED, tr.GetStatus())
	assert.True(t, tr.DeniedAt().IsZero())
	assert.True(t, tr.ApprovedAt().IsZero())
	assert.False(t, tr.TrailDate.DeletedAt().IsZero())
}
