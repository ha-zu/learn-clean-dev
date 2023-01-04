package menteerecruitment

import (
	"unicode/utf8"

	custerr "github.com/ha-zu/learn-clean-dev/src/entity/customerror"
	"github.com/ha-zu/learn-clean-dev/src/entity/user"
)

type MenteeRecruitentPlanProposal struct {
	id       MenteeRecruitmentPlanProposalID
	mrpID    MenteeRecruitmentPlanID
	mentorID user.UserID
	proposal string
}

const MENTEE_PLAN_RECRUITMENT_PROPOSAL_MAX_LEN = 2000

func NewMenteeRecruitentPlanProposal(id MenteeRecruitmentPlanProposalID, mrpID MenteeRecruitmentPlanID,
	mentorID user.UserID, proposal string) (*MenteeRecruitentPlanProposal, error) {

	err := NewMenteeRecruitmentPlanProposalID(id)
	if err != nil {
		return nil, err
	}

	if proposal == "" {
		return nil, custerr.ErrEmptyValue
	}

	if l := utf8.RuneCountInString(proposal); l > MENTEE_PLAN_RECRUITMENT_PROPOSAL_MAX_LEN {
		return nil, custerr.ErrValueIsTooLong
	}

	return &MenteeRecruitentPlanProposal{
		id:       id,
		mrpID:    mrpID,
		mentorID: mentorID,
		proposal: proposal,
	}, nil
}

func (m *MenteeRecruitentPlanProposal) ChangeMenteeRecruitentPlanProposal(proposal string) error {

	if proposal == "" {
		return custerr.ErrEmptyValue
	}

	if l := utf8.RuneCountInString(proposal); l > MENTEE_PLAN_RECRUITMENT_PROPOSAL_MAX_LEN {
		return custerr.ErrValueIsTooLong
	}

	m.proposal = proposal

	return nil

}
