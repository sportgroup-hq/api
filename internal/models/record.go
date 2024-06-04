package models

import (
	"encoding/json"

	"github.com/google/uuid"
)

const (
	AccessScopeCoach   AccessScope = "coach"
	AccessScopeStudent AccessScope = "student"
)

type AccessScope string
type AccessScopes []AccessScope

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

type EventRecords []EventRecord

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
