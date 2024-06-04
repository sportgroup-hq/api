package models

import "github.com/google/uuid"

type EventComment struct {
	ID      uuid.UUID `json:"id" bun:",pk,nullzero"`
	EventID uuid.UUID `json:"-"`

	UserID uuid.UUID `json:"-"`
	User   *User     `json:"user" bun:"rel:belongs-to"`

	Text string `json:"text"`
}

type CreateCommentRequest struct {
	Text string `json:"text" binding:"required,min=1,max=1000"`
}
