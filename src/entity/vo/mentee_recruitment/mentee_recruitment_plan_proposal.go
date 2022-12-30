package menteerecruitment

import "github.com/ha-zu/learn-clean-dev/src/entity/vo"

type MenteeRecruitmentPlanProposalID string

func MenteeRecruitmentPlanProposalIDValidate(id string) (*MenteeRecruitmentPlanProposalID, error) {

	if id == "" {
		return nil, vo.ErrEmptyValue
	}

	mrpID := MenteeRecruitmentPlanProposalID(id)

	return &mrpID, nil
}
