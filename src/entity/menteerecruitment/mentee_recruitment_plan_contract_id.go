package menteerecruitment

import custerr "github.com/ha-zu/learn-clean-dev/src/entity/customerror"

type MenteeRecruitmentPlanContractID string

func NewMenteeRecruitmentPlanContractID(id MenteeRecruitmentPlanContractID) error {

	if id == "" {
		return custerr.ErrEmptyValue
	}

	return nil
}
