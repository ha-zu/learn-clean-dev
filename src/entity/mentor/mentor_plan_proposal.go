package mentor

import (
	"unicode/utf8"

	custerr "github.com/ha-zu/learn-clean-dev/src/entity/customerror"
	"github.com/ha-zu/learn-clean-dev/src/entity/user"
)

type MentorProposal struct {
	id           MentorProposalID
	mentorPlanID MentorPlanID
	menteeID     user.UserID
	proposal     string
}

const MENTOR_PLAN_PROPOSAL_MAX_LEN = 2000

func NewMentorProposal(id, proposal string, mentorPlanID MentorPlanID, menteeID user.UserID) (*MentorProposal, error) {

	mpID, err := MentorProposalIDValidate(id)
	if err != nil {
		return nil, err
	}

	if proposal == "" {
		return nil, custerr.ErrEmptyValue
	}

	if l := utf8.RuneCountInString(proposal); l > MENTOR_PLAN_PROPOSAL_MAX_LEN {
		return nil, custerr.ErrValueIsTooLong
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
		return custerr.ErrEmptyValue
	}

	if l := utf8.RuneCountInString(proposal); l > MENTOR_PLAN_PROPOSAL_MAX_LEN {
		return custerr.ErrValueIsTooLong
	}

	mp.proposal = proposal

	return nil
}
