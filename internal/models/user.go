package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/sportgroup-hq/common-lib/api"
)

type Role string

const (
// RoleTeacher Role = "teacher"
// RoleStudent Role = "student"
)

type User struct {
	ID          uuid.UUID  `json:"id" bun:",pk,nullzero"`
	FirstName   string     `json:"firstName"`
	LastName    string     `json:"lastName"`
	Email       string     `json:"email"`
	PictureURL  string     `json:"pictureURL" bun:",nullzero"`
	Phone       string     `json:"phone" bun:",nullzero"`
	DateOfBirth *time.Time `json:"dateOfBirth,omitempty" bun:",nullzero"`
	CreatedAt   time.Time  `json:"createdAt" bun:",nullzero"`
	UpdatedAt   time.Time  `json:"updatedAt"  bun:",nullzero"`
	//Role      Role      `json:"role" bun:",nullzero"`
}

func UserToPB(user *User) *api.User {
	if user == nil {
		return nil
	}

	return &api.User{
		Id: user.ID.String(),
	}
}
