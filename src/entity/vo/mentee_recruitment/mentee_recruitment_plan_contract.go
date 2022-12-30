package menteerecruitment

import "github.com/ha-zu/learn-clean-dev/src/entity/vo"

type MenteeRecruitmentPlanContractID string

func MenteeRecruitmentPlanContractIDValidate(id string) (*MenteeRecruitmentPlanContractID, error) {

	if id == "" {
		return nil, vo.ErrEmptyValue
	}

	mrpcID := MenteeRecruitmentPlanContractID(id)

	return &mrpcID, nil
}
