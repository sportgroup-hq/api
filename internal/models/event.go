package models

import (
	"time"

	"github.com/google/uuid"
)

const (
	RecordTypeCheckbox RecordType = "checkbox" // true/false
	RecordTypeRating   RecordType = "rating"   // 1-5
	RecordTypeText     RecordType = "text"     // string
	RecordTypeNumber   RecordType = "number"   // number
	RecordTypePhoto    RecordType = "photo"
	RecordTypeVideo    RecordType = "video"
	RecordTypeFile     RecordType = "file"
)

const (
	AssignTypeAll      AssignType = "all"
	AssignTypeSelected AssignType = "selected"
)

type AssignType string
type RecordType string

type Event struct {
	ID        uuid.UUID `json:"id" bun:",pk,nullzero"`
	GroupID   uuid.UUID `json:"-" bun:",notnull"`
	CreatedBy uuid.UUID `json:"-"`

	Title       string `json:"title"`
	Description string `json:"description"`
	Location    string `json:"location"`

	StartAt time.Time `json:"startAt" bun:",nullzero"`
	EndAt   time.Time `json:"endAt" bun:",nullzero"`

	Records EventRecords `json:"records" bun:",type:jsonb"`

	// "all" or "selected"
	AssignType AssignType `json:"assignType"`
	//AssignedUserIDs []uuid.UUID `json:"assignedUserIDs" bun:"assigned_user_ids"`
	AssignedUsers []User `json:"assigned_users" bun:"m2m:event_assignees,join:Event=User"`

	CreatedAt time.Time `json:"createdAt" bun:",nullzero"`
	UpdatedAt time.Time `json:"-" bun:",nullzero"`
}

type Events []Event

type EventAssignee struct {
	ID        uuid.UUID `bun:",pk"`
	EventID   uuid.UUID `bun:",notnull"`
	Event     *Event    `bun:"rel:belongs-to"`
	UserID    uuid.UUID `bun:",notnull"`
	User      *User     `bun:"rel:belongs-to"`
	CreatedAt time.Time `bun:",nullzero"`
}

func (e *Event) FilterRecordsByAccess(memberType GroupMemberType) {
	records := make(EventRecords, 0, len(e.Records))

	for i := range e.Records {
		if e.Records[i].ReadAccessScopes.AllowedForMemberType(memberType) {
			records = append(records, e.Records[i])
		}
	}

	e.Records = records
}

func (e *Event) GetRecord(recordID uuid.UUID) *EventRecord {
	for _, record := range e.Records {
		if record.ID == recordID {
			return &record
		}
	}

	return nil
}

func (e *Event) AssignValues(values []EventRecordValue) {
	m := make(map[uuid.UUID]*EventRecord, len(e.Records))

	for i := range e.Records {
		m[e.Records[i].ID] = &e.Records[i]
	}

	for i := range values {
		if record, ok := m[values[i].RecordID]; ok {
			record.Value = values[i].Value
		}
	}
}

func (e Events) FilterRecordsByAccess(memberType GroupMemberType) Events {
	for i := range e {
		e[i].FilterRecordsByAccess(memberType)
	}

	return e
}

func (e Events) IDs() []uuid.UUID {
	ids := make([]uuid.UUID, 0, len(e))

	for _, event := range e {
		ids = append(ids, event.ID)
	}

	return ids
}

func (e Events) AssignValues(values []EventRecordValue) {
	m := make(map[uuid.UUID]*EventRecord, len(values))

	for i := range e {
		for j := range e[i].Records {
			m[e[i].Records[j].ID] = &(e[i].Records[j])
		}
	}

	for i := range values {
		if record, ok := m[values[i].RecordID]; ok {
			record.Value = values[i].Value
		}
	}
}

func (r EventRecords) ContainsNotUniqueTitle() bool {
	titles := make(map[string]struct{})

	for _, record := range r {
		if _, ok := titles[record.Title]; ok {
			return true
		}

		titles[record.Title] = struct{}{}
	}

	return false
}
