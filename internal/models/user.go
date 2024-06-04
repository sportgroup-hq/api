package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/sportgroup-hq/common-lib/api"
)

type Role string

const (
// RoleTeacher Role = "coach"
// RoleStudent Role = "student"
)

type User struct {
	ID          uuid.UUID  `json:"id" bun:",pk,nullzero"`
	FirstName   string     `json:"firstName"`
	LastName    string     `json:"lastName"`
	Email       string     `json:"email"`
	PictureURL  string     `json:"pictureURL" bun:",nullzero"`
	Phone       string     `json:"phone" bun:",nullzero"`
	DateOfBirth *time.Time `json:"dateOfBirth" bun:",nullzero"`
	Sex         string     `json:"sex" bun:",nullzero"`
	Address     string     `json:"address" bun:",nullzero"`
	CreatedAt   time.Time  `json:"-" bun:",nullzero"`
	UpdatedAt   time.Time  `json:"-"  bun:",nullzero"`
	//Role      Role      `json:"role" bun:",nullzero"`
}

type UpdateUserRequest struct {
	ID          uuid.UUID  `json:"-"`
	FirstName   *string    `json:"firstName"`
	LastName    *string    `json:"lastName"`
	DateOfBirth *time.Time `json:"dateOfBirth"`
	Sex         *string    `json:"sex"`
	Phone       *string    `json:"phone"`
	Address     *string    `json:"address"`
}

func (r UpdateUserRequest) Apply(user *User) *User {
	if r.FirstName != nil {
		user.FirstName = *r.FirstName
	}

	if r.LastName != nil {
		user.LastName = *r.LastName
	}

	if r.DateOfBirth != nil {
		user.DateOfBirth = r.DateOfBirth
	}

	if r.Sex != nil {
		user.Sex = *r.Sex
	}

	if r.Phone != nil {
		user.Phone = *r.Phone
	}

	if r.Address != nil {
		user.Address = *r.Address
	}

	return user
}

func UserToPB(user *User) *api.User {
	if user == nil {
		return nil
	}

	return &api.User{
		Id: user.ID.String(),
	}
}
