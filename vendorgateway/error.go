package vendorgateway

import "errors"

var (
	ErrEmptyMessage   = errors.New("inavalid argument: message is empthy")
	ErrEmptySessionId = errors.New("invalid argument: session id is empty")
	ErrEmptyUserInput = errors.New("invalid argument: user input is empty")
)
