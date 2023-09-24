package sesv2

import (
	"errors"
	"fmt"
)

func fErr(param string) error {
	return fmt.Errorf(`parameter "%v" is required.`, param)
}

var (
	ErrNoProfileSet       = errors.New("no profile set")
	ErrProfileNotSearched = errors.New("profile no exists")
	ErrInvalidTo          = fErr("to")
	ErrInvalidFrom        = fErr("from")
	ErrInvalidMessage     = fErr("message")
	ErrInvalidTitle       = fErr("to")
	ErrInvalidMessageHTML = fErr("messageHtml")
)
