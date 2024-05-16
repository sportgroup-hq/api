package models

import (
	"time"

	"github.com/google/uuid"
)

type GroupMemberType string

const (
	GroupMemberTypeStudent GroupMemberType = "student"
	GroupMemberTypeAdmin   GroupMemberType = "admin"
	GroupMemberTypeOwner   GroupMemberType = "owner"
)

type Group struct {
	ID        uuid.UUID `json:"id" bun:",pk,nullzero"`
	Name      string    `json:"name"`
	Sport     string    `json:"sport"`
	CreatedAt time.Time `json:"createdAt" bun:",nullzero"`
	UpdatedAt time.Time `json:"updatedAt" bun:",nullzero"`
}

type GroupMember struct {
	ID        uuid.UUID       `json:"id" bun:",pk,nullzero"`
	GroupID   uuid.UUID       `json:"groupId"`
	UserID    uuid.UUID       `json:"userId"`
	Type      GroupMemberType `json:"type"`
	CreatedAt time.Time       `json:"createdAt" bun:",nullzero"`
	UpdatedAt time.Time       `json:"updatedAt" bun:",nullzero"`
}

type GroupInvite struct {
	ID      uuid.UUID `json:"id" bun:",pk,nullzero"`
	GroupID uuid.UUID `json:"groupId"`
	Code    string    `json:"code"`
	Active  string    `json:"action"`

	CreatedAt time.Time `json:"createdAt" bun:",nullzero"`
	UpdatedAt time.Time `json:"updatedAt" bun:",nullzero"`
}
