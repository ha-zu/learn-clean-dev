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

	pp := &MenteeRecruitentPlanProposal{
		id:       id,
		mrpID:    mrpID,
		mentorID: mentorID,
	}

	err := pp.ValidateMenteeRecruitentPlanProposal(proposal)
	if err != nil {
		return nil, err
	}

	return pp, nil

}

func (m *MenteeRecruitentPlanProposal) ValidateMenteeRecruitentPlanProposal(proposal string) error {

	if proposal == "" {
		return custerr.ErrEmptyValue
	}

	if l := utf8.RuneCountInString(proposal); l > MENTEE_PLAN_RECRUITMENT_PROPOSAL_MAX_LEN {
		return custerr.ErrValueIsTooLong
	}

	m.proposal = proposal

	return nil

}
