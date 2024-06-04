package models

import (
	"encoding/json"
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
	AccessScopeCoach   AccessScope = "coach"
	AccessScopeStudent AccessScope = "student"
)

const (
	AssignTypeAll      AssignType = "all"
	AssignTypeSelected AssignType = "selected"
)

type AssignType string
type RecordType string

type AccessScope string
type AccessScopes []AccessScope

type UpdateEventRequest CreateEventRequest

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

type EventRecords []EventRecord

type EventRecord struct {
	ID      uuid.UUID `json:"id"`
	EventID uuid.UUID `json:"-"`

	Title string     `json:"title"`
	Type  RecordType `json:"type"`

	ReadAccessScopes  AccessScopes `json:"readAccessScopes"`
	WriteAccessScopes AccessScopes `json:"writeAccessScopes"`

	// Value set by user
	Value *json.RawMessage `json:"value"`
}

type EventRecordValue struct {
	ID      uuid.UUID `json:"-" bun:",pk,nullzero"`
	EventID uuid.UUID `json:"-"`

	RecordID uuid.UUID `json:"-"`
	UserID   uuid.UUID `json:"-"`

	// null or actual value
	Value *json.RawMessage `json:"value" bun:",nullzero"`
}

// GroupRecord - default record defined for group
type GroupRecord struct {
	Title             string       `json:"title"`
	Type              RecordType   `json:"type"`
	ReadAccessScopes  AccessScopes `json:"readAccessScopes" bun:",array"`
	WriteAccessScopes AccessScopes `json:"writeAccessScopes"  bun:",array"`
}

type CreateEventRequest struct {
	Title       string    `json:"title" binding:"required"`
	Description string    `json:"description"`
	Location    string    `json:"location"`
	StartAt     time.Time `json:"startAt" binding:"required"`
	EndAt       time.Time `json:"endAt" binding:"required"`

	AssignType      AssignType  `json:"assignType" binding:"required,oneof=all selected"`
	AssignedUserIDs []uuid.UUID `json:"assignedUserIDs" binding:"required_if=AssignType selected,dive,uuid4"`

	Records CreateRecordsRequest `json:"records" binding:"omitempty,required,unique=Title"`
}

type CreateRecordRequest struct {
	Title             string       `json:"title" binding:"required"`
	Type              RecordType   `json:"type" binding:"required,oneof=checkbox rating text number photo video file"`
	ReadAccessScopes  AccessScopes `json:"readAccessScopes" binding:"required,dive,oneof=coach student"`
	WriteAccessScopes AccessScopes `json:"writeAccessScopes" binding:"required,dive,oneof=coach student"`
}

type CreateRecordsRequest []CreateRecordRequest

func (s AccessScopes) AllowedForMemberType(memberType GroupMemberType) bool {
	for _, scope := range s {
		if scope == AccessScopeCoach && memberType == GroupMemberTypeCoach {
			return true
		}

		if scope == AccessScopeStudent && memberType == GroupMemberTypeStudent {
			return true
		}
	}

	return false
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
