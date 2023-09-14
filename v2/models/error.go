package models

type ErrInvalidTo struct{}

func (e ErrInvalidTo) Error() string {
	return "parameter 'to' is required."
}

type ErrInvalidFrom struct{}

func (e ErrInvalidFrom) Error() string {
	return "parameter 'from' is required."
}

type ErrInvalidMessage struct{}

func (e ErrInvalidMessage) Error() string {
	return "parameter 'message' is required."
}

type ErrInvalidTitle struct{}

func (e ErrInvalidTitle) Error() string {
	return "parameter 'title' is required."
}

type ErrInvalidMessageHTML struct{}

func (e ErrInvalidMessageHTML) Error() string {
	return "parameter 'messageHtml' is required."
}
