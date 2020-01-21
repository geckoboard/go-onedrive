package onedrive

import "errors"

// error types

var (
	// ErrFileTooLarge represents the error returned when the file is too large for upload
	ErrFileTooLarge = errors.New("file is too large for simple upload")
)

// The Error type defines the basic structure of errors that are returned from
// the OneDrive API.
// Error messages can be nested recursively, with inner errors containing more
// specific error codes
// See: https://docs.microsoft.com/en-gb/onedrive/developer/rest-api/concepts/errors?view=odsp-graph-online
type Error struct {
	Code       string `json:"code"`
	Message    string `json:"message"`
	InnerError *Error `json:"innererror"`
}

func (e Error) Error() string {
	return e.Message
}
