package mentor

import custerr "github.com/ha-zu/learn-clean-dev/src/entity/customerror"

type MentorProposalID string

func MentorProposalIDValidate(id string) (*MentorProposalID, error) {

	if id == "" {
		return nil, custerr.ErrEmptyValue
	}

	mpID := MentorProposalID(id)

	return &mpID, nil

}
