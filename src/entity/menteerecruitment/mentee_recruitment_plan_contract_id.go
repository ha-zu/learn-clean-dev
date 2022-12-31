package menteerecruitment

import custerr "github.com/ha-zu/learn-clean-dev/src/entity/customerror"

type MenteeRecruitmentPlanContractID string

func MenteeRecruitmentPlanContractIDValidate(id string) (*MenteeRecruitmentPlanContractID, error) {

	if id == "" {
		return nil, custerr.ErrEmptyValue
	}

	mrpcID := MenteeRecruitmentPlanContractID(id)

	return &mrpcID, nil
}
