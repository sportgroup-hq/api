package models

import (
	"time"

	"github.com/google/uuid"
)

type GroupMemberType string

const (
	GroupMemberTypeStudent GroupMemberType = "student"
	GroupMemberTypeCoach   GroupMemberType = "coach"
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
	Active  bool      `json:"action"`

	CreatedAt time.Time `json:"createdAt" bun:",nullzero"`
	UpdatedAt time.Time `json:"updatedAt" bun:",nullzero"`
}

type JoinGroupRequest struct {
	Code string `json:"code" binding:"required,len=6"`
}

type CreateGroupRequest struct {
	Name  string `json:"name" binding:"required,min=1,max=255"`
	Sport string `json:"sport" binding:"required,min=1,max=255"`
}

type UpdateGroupRequest struct {
	ID    uuid.UUID `json:"-"`
	Name  *string   `json:"name" binding:"omitempty,min=1,max=255"`
	Sport *string   `json:"sport" binding:"omitempty,min=1,max=255"`
}

func (m GroupMember) CanEditGroup() bool {
	return m.Type == GroupMemberTypeCoach
}

func (m GroupMember) CanDeleteGroup() bool {
	return m.Type == GroupMemberTypeCoach
}

func (m GroupMember) CanDeleteEvent() bool {
	return m.Type == GroupMemberTypeCoach
}

func (m GroupMember) CanCreateEvent() bool {
	return m.Type == GroupMemberTypeCoach
}

func (m GroupMember) CanAccessGroupRecords() bool {
	return m.Type == GroupMemberTypeCoach
}

func (r CreateGroupRequest) ToGroup() *Group {
	return &Group{
		Name:  r.Name,
		Sport: r.Sport,
	}
}

func (r UpdateGroupRequest) Apply(group *Group) *Group {
	if r.Name != nil {
		group.Name = *r.Name
	}

	if r.Sport != nil {
		group.Sport = *r.Sport
	}

	return group
}
