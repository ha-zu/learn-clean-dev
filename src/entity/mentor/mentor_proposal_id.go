package mentor

import (
	"github.com/google/uuid"
	custerr "github.com/ha-zu/learn-clean-dev/src/entity/customerror"
)

type MentorProposalID string

func NewMentorProposalID() MentorProposalID {
	return MentorProposalID(uuid.New().String())
}

func NewMentorProposalIDStr(id string) (MentorProposalID, error) {

	if id == "" {
		return MentorProposalID(""), custerr.ErrEmptyValue
	}

	return MentorProposalID(id), nil

}
