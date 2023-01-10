package mentor

import (
	"github.com/google/uuid"
	custerr "github.com/ha-zu/learn-clean-dev/src/entity/customerror"
)

type MentorContractID string

func NewMentorContractID() MentorContractID {
	return MentorContractID(uuid.New().String())
}

func NewMentorContractIDStr(id string) (MentorContractID, error) {

	if id == "" {
		return MentorContractID(""), custerr.ErrEmptyValue
	}

	return MentorContractID(id), nil

}
