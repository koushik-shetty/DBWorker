package lib_test

import (
	"testing"

	"github.com/golang/protobuf/proto"
	"github.com/stretchr/testify/assert"
	"mct-git.sb.mct.local/periscope/performanceanalytics/constants"
	"mct-git.sb.mct.local/periscope/performanceanalytics/lib"
	msg "mct-git.sb.mct.local/periscope/performanceanalytics/messages"
	pptMsg "mct-git.sb.mct.local/periscope/performanceanalytics/messages/ppt_service_messages"
)

func TestError(t *testing.T) {
	err := lib.NewError("description", lib.ServerError)
	assert.Equal(t, "description", err.Description())
	assert.Equal(t, constants.ServerError, err.UserCode())
	assert.Equal(t, "", err.UserMessage())
	assert.Equal(t, 500, err.ErrorCode())
}

func TestWithUserMessage(t *testing.T) {
	err := lib.NewError("description", lib.ServerError).WithUserMsg("reason for error")
	assert.Equal(t, constants.ServerError, err.UserCode())
	assert.Equal(t, "reason for error", err.UserMessage())
}

func TestErrorAppendDescription(t *testing.T) {
	err := lib.NewError("description one", lib.ServerError)
	err.AppendDescription("description two")
	assert.Equal(t, "description one; description two", err.Description())
}

func TestCombineProtobufErrors(t *testing.T) {
	errors := []*msg.Error{
		&msg.Error{Code: proto.String("1"), Message: proto.String("Error One")},
		&msg.Error{Code: proto.String("2"), Message: proto.String("Error Two")},
	}
	assert.Equal(t, "Error One (code=1), Error Two (code=2)", lib.CombineProtobufErrors(errors))
}

func TestCombinePPTProtobufErrors(t *testing.T) {
	errors := []*pptMsg.Error{
		&pptMsg.Error{Code: proto.String("1"), Message: proto.String("Error One")},
		&pptMsg.Error{Code: proto.String("2"), Message: proto.String("Error Two")},
	}
	assert.Equal(t, "Error One (code=1), Error Two (code=2)", lib.CombinePPTProtobufErrors(errors))
}

func TestCombineProtobufErrorsIgnoresMissingCode(t *testing.T) {
	errors := []*msg.Error{
		&msg.Error{Message: proto.String("Error")},
	}
	assert.Equal(t, "Error", lib.CombineProtobufErrors(errors))
}

func TestCombineProtobufErrorsIgnoresMissingMessage(t *testing.T) {
	errors := []*msg.Error{
		&msg.Error{Code: proto.String("1")},
	}
	assert.Equal(t, "(code=1)", lib.CombineProtobufErrors(errors))
}
