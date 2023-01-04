package menteerecruitment

import custerr "github.com/ha-zu/learn-clean-dev/src/entity/customerror"

type MenteeRecruitmentPlanProposalID string

func NewMenteeRecruitmentPlanProposalID(id MenteeRecruitmentPlanProposalID) error {

	if id == "" {
		return custerr.ErrEmptyValue
	}

	return nil
}
