package menteerecruitment

import custerr "github.com/ha-zu/learn-clean-dev/src/entity/customerror"

type MenteeRecruitmentPlanProposalID string

func MenteeRecruitmentPlanProposalIDValidate(id string) (*MenteeRecruitmentPlanProposalID, error) {

	if id == "" {
		return nil, custerr.ErrEmptyValue
	}

	mrpID := MenteeRecruitmentPlanProposalID(id)

	return &mrpID, nil
}
