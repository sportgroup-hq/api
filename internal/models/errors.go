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
	ErrForbidden     = err(1004, "forbidden", http.StatusForbidden)

	ErrGroupInviteInactive  = err(2001, "group invite is inactive", http.StatusForbidden)
	ErrAlreadyJoined        = err(2002, "already joined", http.StatusConflict)
	ErrNotJoined            = err(2003, "not joined", http.StatusForbidden)
	ErrCoachCannotLeave     = err(2004, "coach cannot leave", http.StatusForbidden)
	ErrAssigneesRequired    = err(2005, "assignees required", http.StatusBadRequest)
	ErrRecordTitleNotUnique = err(2006, "record title not unique", http.StatusBadRequest)
	ErrRecordNotFound       = err(2007, "record not found", http.StatusNotFound)
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
