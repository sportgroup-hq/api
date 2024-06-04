package models

import (
	"time"

	"github.com/google/uuid"
)

type CreateEventRequest struct {
	Title       string    `json:"title" binding:"required,min=1,max=255"`
	Description string    `json:"description" binding:"omitempty,max=1000"`
	Location    string    `json:"location" binding:"omitempty,max=255"`
	StartAt     time.Time `json:"startAt" binding:"required,ltefield=EndAt"`
	EndAt       time.Time `json:"endAt" binding:"required,gtefield=StartAt"`

	AssignType      AssignType  `json:"assignType" binding:"required,oneof=all selected"`
	AssignedUserIDs []uuid.UUID `json:"assignedUserIDs" binding:"required_if=AssignType selected,dive,uuid4"`

	Records CreateRecordsRequest `json:"records" binding:"omitempty,required,unique=Title"`
}

type CreateRecordRequest struct {
	Title             string       `json:"title" binding:"required,min=1,max=255"`
	Type              RecordType   `json:"type" binding:"required,oneof=checkbox rating text number photo video file"`
	ReadAccessScopes  AccessScopes `json:"readAccessScopes" binding:"required,dive,oneof=coach student"`
	WriteAccessScopes AccessScopes `json:"writeAccessScopes" binding:"required,dive,oneof=coach student"`
}

type CreateRecordsRequest []CreateRecordRequest

func (r CreateEventRequest) ToEvent() *Event {
	return &Event{
		Title:       r.Title,
		Description: r.Description,
		Location:    r.Location,
		StartAt:     r.StartAt,
		EndAt:       r.EndAt,

		AssignType: r.AssignType,
	}
}

func (r CreateRecordsRequest) ToEventRecords(eventID uuid.UUID) []EventRecord {
	var records []EventRecord

	for _, record := range r {
		records = append(records, EventRecord{
			ID:      uuid.New(),
			EventID: eventID,

			Title: record.Title,
			Type:  record.Type,

			ReadAccessScopes:  record.ReadAccessScopes,
			WriteAccessScopes: record.WriteAccessScopes,
		})
	}

	return records
}
