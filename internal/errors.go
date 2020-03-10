package model

import (
	"github.com/pkg/errors"
)

// AppError ID Numbering Format
// ID numbering format: 1XXXYYYY
// The first part: 1XXX represent the domain
// The second part: YYYY represent the error sequence number of the domain

// AppError custom error
type AppError struct {
	ID      int64
	Message string
}

// Error implements error interface
func (a AppError) Error() string {
	return a.Message
}

// ErrGeneric is used for testing purposes and for errors handled later in callstack
var ErrGeneric = errors.New("generic error")

// ErrCastingFailure is used to represent error when casting
var ErrCastingFailure = errors.New("failed to cast")

var (
	// ErrStartDateAfterCurrentEndDate error
	ErrStartDateAfterCurrentEndDate = AppError{
		ID:      10010001,
		Message: "Start Date must be smaller than current End Date",
	}

	// ErrEndDateBeforeCurrentStartDate error
	ErrEndDateBeforeCurrentStartDate = AppError{
		ID:      10010002,
		Message: "End Date must be larger than current Start Date",
	}

	// ErrEndDateBeforePresentTime error
	ErrEndDateBeforePresentTime = AppError{
		ID:      10010003,
		Message: "End Date must be set after current time",
	}
)
