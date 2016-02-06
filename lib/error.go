package lib

import (
	"fmt"
)

const (
	DirError        = "1"
	FileError       = "2"
	RegexError      = "3"
	FileCreateError = "4"
	TokenError      = "5"
)

type Error struct {
	code        string
	source      string
	description string
}

func NewError(code, source, description string) *Error {
	return &Error{
		code:        code,
		source:      source,
		description: description,
	}
}

func EmptyError() *Error {
	return NewError("0", "", "")
}

func (e *Error) String() string {
	return fmt.Sprintf("code=%s,source=%s,description=%s", e.Code(), e.Source(), e.Description())
}

func ToLibError(err error, code, source string) *Error {
	return NewError(code, source, err.Error())
}

func (err *Error) Description() string {
	return err.description
}

func (err *Error) Code() string {
	return err.code
}

func (err *Error) Source() string {
	return err.source
}

func (err *Error) AppendDescription(description string) {
	err.description = fmt.Sprintf("%s; %s", err.description, description)
}
