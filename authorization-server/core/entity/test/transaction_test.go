package test

import (
	"testing"

	"github.com/renatospaka/authorization-server/core/entity"
	"github.com/stretchr/testify/assert"
)

const (
	CLIENT_ID        = "ebf77a6e-e236-441a-8ef6-4a3afdc5d4b1"
)

func TestNewAuthorization(t *testing.T) {
	auth, err := entity.NewAuthorization(CLIENT_ID, 200.00)
	assert.Nil(t, err)
	assert.NotNil(t, auth)
	assert.NotEmpty(t, auth.GetID())
	assert.Equal(t, CLIENT_ID, auth.GetClientID())
	assert.Equal(t, float32(200.00), auth.GetValue())
	assert.True(t, auth.DeniedAt().IsZero())
	assert.True(t, auth.ApprovedAt().IsZero())
	assert.Equal(t, entity.TR_PENDING, auth.GetStatus())

	err = auth.Validate()
	assert.Nil(t, err)
}

func TestNewAuthorizationInvalidlientID(t *testing.T) {
	auth, err := entity.NewAuthorization("X"+CLIENT_ID, 0)
	assert.NotNil(t, err)
	assert.Nil(t, auth)
	assert.EqualError(t, err, entity.ErrInvalidClientID.Error())

	auth2, err2 := entity.NewAuthorization("", 0)
	assert.NotNil(t, err2)
	assert.Nil(t, auth2)
	assert.EqualError(t, err2, entity.ErrInvalidClientID.Error())
}

func TestNewAuthorizationValueIsZero(t *testing.T) {
	auth, err := entity.NewAuthorization(CLIENT_ID, 0)
	assert.NotNil(t, err)
	assert.Nil(t, auth)
	assert.EqualError(t, err, entity.ErrValueIsZero.Error())
}

func TestNewAuthorizationValueIsNegative(t *testing.T) {
	auth, err := entity.NewAuthorization(CLIENT_ID, -200.00)
	assert.NotNil(t, err)
	assert.Nil(t, auth)
	assert.EqualError(t, err, entity.ErrValueIsNegative.Error())
}

func TestApprove(t *testing.T) {
	auth, err := entity.NewAuthorization(CLIENT_ID, 200.00)
	assert.Nil(t, err)
	assert.NotNil(t, auth)
	assert.Equal(t, entity.TR_PENDING, auth.GetStatus())
	assert.True(t, auth.DeniedAt().IsZero())
	assert.True(t, auth.ApprovedAt().IsZero())

	auth.Approve()
	assert.Equal(t, entity.TR_APPROVED, auth.GetStatus())
	assert.True(t, auth.DeniedAt().IsZero())
	assert.False(t, auth.ApprovedAt().IsZero())
}

func TestDeny(t *testing.T) {
	auth, err := entity.NewAuthorization(CLIENT_ID, 200.00)
	assert.Nil(t, err)
	assert.NotNil(t, auth)
	assert.Equal(t, entity.TR_PENDING, auth.GetStatus())
	assert.True(t, auth.DeniedAt().IsZero())
	assert.True(t, auth.ApprovedAt().IsZero())

	auth.Deny()
	assert.Equal(t, entity.TR_DENIED, auth.GetStatus())
	assert.False(t, auth.DeniedAt().IsZero())
	assert.True(t, auth.ApprovedAt().IsZero())
}

func TestDeleteIt(t *testing.T) {
	auth, err := entity.NewAuthorization(CLIENT_ID, 200.00)
	assert.Nil(t, err)
	assert.NotNil(t, auth)
	assert.Equal(t, entity.TR_PENDING, auth.GetStatus())
	assert.True(t, auth.DeniedAt().IsZero())
	assert.True(t, auth.ApprovedAt().IsZero())

	auth.DeleteIt()
	assert.Equal(t, entity.TR_DELETED, auth.GetStatus())
	assert.True(t, auth.DeniedAt().IsZero())
	assert.True(t, auth.ApprovedAt().IsZero())
	assert.False(t, auth.TrailDate.DeletedAt().IsZero())
}
