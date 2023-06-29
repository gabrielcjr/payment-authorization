package entity

import "time"

var (
	location *time.Location
)

type TrailDate struct {
	createdAt time.Time
	updatedAt time.Time
	deletedAt time.Time
}

func NewTrailDate() *TrailDate {
	return &TrailDate{
		createdAt: getNow(),
		updatedAt: getNow(),
		deletedAt: time.Time{},
	}
}

func (t *TrailDate) SetCreationToToday() {
	t.createdAt = getNow()
	t.updatedAt = getNow()
	t.deletedAt = time.Time{}
}

func (t *TrailDate) SetCreationToDate(date time.Time) {
	t.createdAt = date
	t.updatedAt = date
	t.deletedAt = time.Time{}
}

func (t *TrailDate) SetAlterationToToday() {
	t.updatedAt = getNow()
	t.deletedAt = time.Time{}
}

func (t *TrailDate) SetAlterationToDate(date time.Time) {
	t.updatedAt = date
	t.deletedAt = time.Time{}
}

func (t *TrailDate) SetDeletionToToday() {
	t.deletedAt = getNow()
}

func (t *TrailDate) SetDeletionToDate(date time.Time) {
	t.deletedAt = date
}

func (t *TrailDate) CreatedAt() time.Time {
	return t.createdAt
}

func (t *TrailDate) DeletedAt() time.Time {
	return t.deletedAt
}

func (t *TrailDate) UpdatedAt() time.Time {
	return t.updatedAt
}

func getNow() time.Time {
	now := time.Now()
	location := getLocation()
	date := time.Date(now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second(), 0, location)
	return date
}

func getLocation() *time.Location {
	location, _ := time.LoadLocation("America/Sao_Paulo")
	return location
}
