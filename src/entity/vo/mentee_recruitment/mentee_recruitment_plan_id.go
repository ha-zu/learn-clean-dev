package menteerecruitment

import "github.com/ha-zu/learn-clean-dev/src/entity/vo"

type MenteeRecruitmentPlanID string

func MenteePlanIDValidate(id string) (*MenteeRecruitmentPlanID, error) {

	if id == "" {
		return nil, vo.ErrEmptyValue
	}

	mpID := MenteeRecruitmentPlanID(id)

	return &mpID, nil
}
