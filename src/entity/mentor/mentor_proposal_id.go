package mentor

import custerr "github.com/ha-zu/learn-clean-dev/src/entity/customerror"

type MentorProposalID string

func NewMentorProposalID(id MentorProposalID) error {

	if id == "" {
		return custerr.ErrEmptyValue
	}

	return nil

}
