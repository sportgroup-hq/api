package models

import (
	"errors"
	"net/http"
)

type InternalError struct {
	error
	InternalCode   int
	HTTPStatusCode int
}

var (
	ErrUnknown       = err(1000, "unknown error", http.StatusInternalServerError)
	ErrNotFound      = err(1001, "not found", http.StatusNotFound)
	ErrDuplicate     = err(1002, "duplicate", http.StatusConflict)
	ErrPathMalformed = err(1003, "path malformed", http.StatusBadRequest)

	ErrGroupInviteInactive = err(2001, "group invite is inactive", http.StatusForbidden)
	ErrAlreadyJoined       = err(2002, "already joined", http.StatusConflict)
	ErrNotJoined           = err(2003, "not joined", http.StatusForbidden)
	ErrOwnerCannotLeave    = err(2004, "owner cannot leave", http.StatusForbidden)
	ErrNotOwner            = err(2005, "not owner", http.StatusForbidden)
)

func err(internalCode int, text string, httpCode int) InternalError {
	return InternalError{
		InternalCode:   internalCode,
		error:          errors.New(text),
		HTTPStatusCode: httpCode,
	}
}

func Error(err error) InternalError {
	var e InternalError
	ok := errors.As(err, &e)
	if !ok {
		return ErrUnknown
	}

	return e
}
