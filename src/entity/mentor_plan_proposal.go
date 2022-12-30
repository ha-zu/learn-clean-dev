package entity

import (
	"unicode/utf8"

	"github.com/ha-zu/learn-clean-dev/src/entity/vo"
	mr "github.com/ha-zu/learn-clean-dev/src/entity/vo/mentor"
	"github.com/ha-zu/learn-clean-dev/src/entity/vo/user"
)

type MentorProposal struct {
	id           mr.MentorProposalID
	mentorPlanID mr.MentorPlanID
	menteeID     user.UserID
	proposal     string
}

const MENTOR_PLAN_PROPOSAL_MAX_LEN = 2000

func NewMentorProposal(id, proposal string, mentorPlanID mr.MentorPlanID, menteeID user.UserID) (*MentorProposal, error) {

	mpID, err := mr.MentorProposalIDValidate(id)
	if err != nil {
		return nil, err
	}

	if proposal == "" {
		return nil, vo.ErrEmptyValue
	}

	if l := utf8.RuneCountInString(proposal); l > MENTOR_PLAN_PROPOSAL_MAX_LEN {
		return nil, vo.ErrValueIsTooLong
	}

	return &MentorProposal{
		id:           *mpID,
		mentorPlanID: mentorPlanID,
		menteeID:     menteeID,
		proposal:     proposal,
	}, nil
}

func (mp *MentorProposal) ChangeMentorProposal(proposal string) error {

	if proposal == "" {
		return vo.ErrEmptyValue
	}

	if l := utf8.RuneCountInString(proposal); l > MENTOR_PLAN_PROPOSAL_MAX_LEN {
		return vo.ErrValueIsTooLong
	}

	mp.proposal = proposal

	return nil
}
