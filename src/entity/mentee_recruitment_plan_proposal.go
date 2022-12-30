package entity

import (
	"unicode/utf8"

	"github.com/ha-zu/learn-clean-dev/src/entity/vo"
	mr "github.com/ha-zu/learn-clean-dev/src/entity/vo/mentee_recruitment"
	"github.com/ha-zu/learn-clean-dev/src/entity/vo/user"
)

type MenteeRecruitentPlanProposal struct {
	id       mr.MenteeRecruitmentPlanProposalID
	mrpID    mr.MenteeRecruitmentPlanID
	mentorID user.UserID
	proposal string
}

const MENTEE_PLAN_RECRUITMENT_PROPOSAL_MAX_LEN = 2000

func NewMenteeRecruitentPlanProposal(id, proposal string, mrpID mr.MenteeRecruitmentPlanID, mentorID user.UserID) (*MenteeRecruitentPlanProposal, error) {

	// ToDo tag count Logic

	mrppID, err := mr.MenteeRecruitmentPlanProposalIDValidate(id)
	if err != nil {
		return nil, err
	}

	if proposal == "" {
		return nil, vo.ErrEmptyValue
	}

	if l := utf8.RuneCountInString(proposal); l > MENTEE_PLAN_RECRUITMENT_PROPOSAL_MAX_LEN {
		return nil, vo.ErrValueIsTooLong
	}

	return &MenteeRecruitentPlanProposal{
		id:       *mrppID,
		mrpID:    mrpID,
		mentorID: mentorID,
		proposal: proposal,
	}, nil
}

func (m *MenteeRecruitentPlanProposal) ChangeMenteeRecruitentPlanProposal(proposal string) error {

	if proposal == "" {
		return vo.ErrEmptyValue
	}

	if l := utf8.RuneCountInString(proposal); l > MENTEE_PLAN_RECRUITMENT_PROPOSAL_MAX_LEN {
		return vo.ErrValueIsTooLong
	}

	m.proposal = proposal

	return nil

}
