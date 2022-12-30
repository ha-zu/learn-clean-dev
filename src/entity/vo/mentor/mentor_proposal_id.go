package mentor

import "github.com/ha-zu/learn-clean-dev/src/entity/vo"

type MentorProposalID string

func MentorProposalIDValidate(id string) (*MentorProposalID, error) {

	if id == "" {
		return nil, vo.ErrEmptyValue
	}

	mpID := MentorProposalID(id)

	return &mpID, nil

}
