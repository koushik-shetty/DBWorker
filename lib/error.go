package lib

import (
	"fmt"
	"net/http"
	"strings"

	"mct-git.sb.mct.local/periscope/performanceanalytics/constants"
	pptMsg "mct-git.sb.mct.local/periscope/performanceanalytics/messages/ppt_service_messages"
)

type Error struct {
	Code        int
	Source      string
	Description string
}

func NewError(code int, source, description string) *Error {
	return &Error{
		Code:        code,
		Source:      source,
		Description: description,
	}
}

func (e error) ToError(code int, source string) *Error {
	return NewError(e.Error(), userCode)
}

func (err *Error) Description() string {
	return err.description
}

func (err *Error) Code() string {
	return err.userCode.code
}

func (err *Error) Source() string {
	return err.Source
}

func (err *Error) AppendDescription(description string) {
	err.description = fmt.Sprintf("%s; %s", err.description, description)
}
