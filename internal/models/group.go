package models

import (
	"time"

	"github.com/google/uuid"
)

type Group struct {
	ID        uuid.UUID `json:"id" bun:",pk"`
	Name      string    `json:"name"`
	Sport     string    `json:"sport"`
	OwnerID   uuid.UUID `json:"ownerId"`
	Owner     *User     `json:"owner" bun:"rel:belongs-to"`
	CreatedAt time.Time `json:"createdAt" bun:",nullzero"`
	UpdatedAt time.Time `json:"updatedAt" bun:",nullzero"`
}

type GroupInvite struct {
	ID      uuid.UUID `json:"id" bun:",pk"`
	GroupID uuid.UUID `json:"groupId"`
	Code    string    `json:"code"`
	Active  string    `json:"action"`

	CreatedAt time.Time `json:"createdAt" bun:",nullzero"`
	UpdatedAt time.Time `json:"updatedAt" bun:",nullzero"`
}
