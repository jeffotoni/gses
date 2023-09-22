package models

import "errors"

var (
	ErrInvalidTo          = errors.New("parameter 'to' is required.")
	ErrInvalidFrom        = errors.New("parameter 'from' is required.")
	ErrInvalidMessage     = errors.New("parameter 'message' is required.")
	ErrInvalidTitle       = errors.New("parameter 'to' title required.")
	ErrInvalidMessageHTML = errors.New("parameter 'messageHtml' is required.")
)
