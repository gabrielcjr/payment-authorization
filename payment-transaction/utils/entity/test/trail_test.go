package test

import (
	"testing"
	"time"

	"github.com/renatospaka/payment-transaction/utils/entity"
	"github.com/stretchr/testify/assert"
)

var (
	date time.Time
)

func init() {
	location, _ := time.LoadLocation("America/Sao_Paulo")
	date = time.Date(2019, 8, 23, 13, 44, 4, 221, location)
}

func TestNewTrial(t *testing.T) {
	trail := entity.NewTrailDate()
	assert.NotNil(t, trail)

	createdAt := trail.CreatedAt()
	assert.False(t, createdAt.IsZero())

	updatedAt := trail.UpdatedAt()
	assert.False(t, updatedAt.IsZero())

	deletedAt := trail.DeletedAt()
	assert.Equal(t, time.Time{}, deletedAt)
	assert.True(t, deletedAt.IsZero())
}

// CreatedAt tests
func TestSetCreationToToday(t *testing.T) {
	trail := entity.NewTrailDate()
	assert.NotNil(t, trail)

	time.Sleep(50 * time.Millisecond)
	trail.SetCreationToToday()
	now := time.Now().Format("2006-01-02T15:04:05 -07:00:00")
	createdAt := trail.CreatedAt().Format("2006-01-02T15:04:05 -07:00:00")
	assert.NotNil(t, createdAt)
	assert.Equal(t, now, createdAt)
	assert.NotEqual(t, date, createdAt)

	updatedAt := trail.UpdatedAt().Format("2006-01-02T15:04:05 -07:00:00")
	assert.NotNil(t, updatedAt)
	assert.Equal(t, now, updatedAt)

	deletedAt := trail.DeletedAt().Format("2006-01-02T15:04:05 -07:00:00")
	assert.NotEqual(t, now, deletedAt)
	assert.True(t, trail.DeletedAt().IsZero())
}

func TestSetCreationToTodayFail(t *testing.T) {
	trail := entity.NewTrailDate()
	assert.NotNil(t, trail)

	time.Sleep(50 * time.Millisecond)
	trail.SetCreationToToday()
	now := time.Now()
	createdAt := trail.CreatedAt()
	assert.NotNil(t, createdAt)
	assert.NotEqual(t, now, createdAt)
}

func TestSetCreationToDate(t *testing.T) {
	trail := entity.NewTrailDate()
	assert.NotNil(t, trail)

	trail.SetCreationToDate(date)
	now := time.Now().Format("2006-01-02T15:04:05 -07:00:00")
	createdAt := trail.CreatedAt().Format("2006-01-02T15:04:05 -07:00:00")
	dateStr := date.Format("2006-01-02T15:04:05 -07:00:00")
	assert.NotNil(t, createdAt)
	assert.NotEqual(t, now, createdAt)
	assert.Equal(t, dateStr, createdAt)

	updatedAt := trail.UpdatedAt().Format("2006-01-02T15:04:05 -07:00:00")
	assert.NotNil(t, updatedAt)
	assert.Equal(t, dateStr, updatedAt)

	deletedAt := trail.DeletedAt().Format("2006-01-02T15:04:05 -07:00:00")
	assert.NotEqual(t, dateStr, deletedAt)
	assert.True(t, trail.DeletedAt().IsZero())
}

func TestSetCreationToDateFail(t *testing.T) {
	trail := entity.NewTrailDate()
	assert.NotNil(t, trail)

	trail.SetCreationToDate(date)
	now := time.Now()
	createdAt := trail.CreatedAt()
	assert.NotNil(t, createdAt)
	assert.NotEqual(t, now, createdAt)
}

// UpdatedAt tests
func TestSetAlterationToToday(t *testing.T) {
	trail := entity.NewTrailDate()
	createdAtOrig := trail.CreatedAt().Format("2006-01-02T15:04:05 -07:00:00")
	assert.NotNil(t, trail)

	time.Sleep(1 * time.Second)
	trail.SetAlterationToToday()
	now := time.Now().Format("2006-01-02T15:04:05 -07:00:00")

	updatedAt := trail.UpdatedAt().Format("2006-01-02T15:04:05 -07:00:00")
	assert.NotNil(t, updatedAt)
	assert.Equal(t, now, updatedAt)

	createdAt := trail.CreatedAt().Format("2006-01-02T15:04:05 -07:00:00")
	assert.NotNil(t, createdAt)
	assert.Equal(t, createdAtOrig, createdAt)
	assert.NotEqual(t, now, createdAt)
	assert.NotEqual(t, date, createdAt)

	deletedAt := trail.DeletedAt().Format("2006-01-02T15:04:05 -07:00:00")
	assert.NotEqual(t, now, deletedAt)
	assert.True(t, trail.DeletedAt().IsZero())

	wanted, _ := time.ParseDuration("1s")
	diff := trail.UpdatedAt().Sub(trail.CreatedAt())
	assert.Equal(t, wanted, diff)
}

func TestAlterationdToDate(t *testing.T) {
	trail := entity.NewTrailDate()
	createdAtOrig := trail.CreatedAt().Format("2006-01-02T15:04:05 -07:00:00")
	assert.NotNil(t, trail)

	time.Sleep(1 * time.Second)
	trail.SetAlterationToDate(date)
	now := time.Now().Format("2006-01-02T15:04:05 -07:00:00")
	updatedAt := trail.UpdatedAt().Format("2006-01-02T15:04:05 -07:00:00")
	dateStr := date.Format("2006-01-02T15:04:05 -07:00:00")
	assert.NotNil(t, updatedAt)
	assert.NotEqual(t, now, updatedAt)
	assert.Equal(t, dateStr, updatedAt)

	createdAt := trail.CreatedAt().Format("2006-01-02T15:04:05 -07:00:00")
	assert.NotNil(t, createdAt)
	assert.Equal(t, createdAtOrig, createdAt)

	deletedAt := trail.DeletedAt().Format("2006-01-02T15:04:05 -07:00:00")
	assert.NotEqual(t, now, deletedAt)
	assert.True(t, trail.DeletedAt().IsZero())
}

// DeleteddAt tests
func TestSetDeletionToToday(t *testing.T) {
	trail := entity.NewTrailDate()
	assert.NotNil(t, trail)
	createdAtOrig := trail.CreatedAt().Format("2006-01-02T15:04:05 -07:00:00")
	updatedAtOrig := trail.CreatedAt().Format("2006-01-02T15:04:05 -07:00:00")

	time.Sleep(1 * time.Second)
	trail.SetDeletionToToday()
	now := time.Now().Format("2006-01-02T15:04:05 -07:00:00")
	deletedAt := trail.DeletedAt().Format("2006-01-02T15:04:05 -07:00:00")
	assert.NotNil(t, deletedAt)
	assert.Equal(t, now, deletedAt)
	assert.NotEqual(t, date, deletedAt)

	createdAt := trail.CreatedAt().Format("2006-01-02T15:04:05 -07:00:00")
	assert.NotNil(t, createdAt)
	assert.Equal(t, createdAtOrig, createdAt)

	updatedAt := trail.UpdatedAt().Format("2006-01-02T15:04:05 -07:00:00")
	assert.NotNil(t, updatedAt)
	assert.Equal(t, updatedAtOrig, updatedAt)
}

func TestDeletiondToDate(t *testing.T) {
	trail := entity.NewTrailDate()
	assert.NotNil(t, trail)
	createdAtOrig := trail.CreatedAt().Format("2006-01-02T15:04:05 -07:00:00")
	updatedAtOrig := trail.CreatedAt().Format("2006-01-02T15:04:05 -07:00:00")

	time.Sleep(1 * time.Second)
	trail.SetDeletionToDate(date)
	now := time.Now().Format("2006-01-02T15:04:05 -07:00:00")
	deletedAt := trail.DeletedAt().Format("2006-01-02T15:04:05 -07:00:00")
	dateStr := date.Format("2006-01-02T15:04:05 -07:00:00")
	assert.NotNil(t, deletedAt)
	assert.NotEqual(t, now, deletedAt)
	assert.Equal(t, dateStr, deletedAt)

	createdAt := trail.CreatedAt().Format("2006-01-02T15:04:05 -07:00:00")
	assert.NotNil(t, createdAt)
	assert.Equal(t, createdAtOrig, createdAt)

	updatedAt := trail.UpdatedAt().Format("2006-01-02T15:04:05 -07:00:00")
	assert.NotNil(t, updatedAt)
	assert.Equal(t, updatedAtOrig, updatedAt)
}
