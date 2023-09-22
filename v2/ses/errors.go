package ses

import "errors"

var (
	ErrNoProfileSet       = errors.New("no profile set")
	ErrProfileNotSearched = errors.New("profile no exists")
)
