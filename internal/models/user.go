package models

import (
	"github.com/google/uuid"
	"github.com/sportgroup-hq/common-lib/api"
)

type Role string

const (
// RoleTeacher Role = "teacher"
// RoleStudent Role = "student"
)

type User struct {
	ID        uuid.UUID `json:"id" bun:",pk,nullzero"`
	FirstName string    `json:"firstName"`
	LastName  string    `json:"lastName"`
	Email     string    `json:"email"`
	Picture   string    `json:"picture"`
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
