package menteerecruitment

import custerr "github.com/ha-zu/learn-clean-dev/src/entity/customerror"

type MenteeRecruitmentPlanID string

func MenteePlanIDValidate(id string) (*MenteeRecruitmentPlanID, error) {

	if id == "" {
		return nil, custerr.ErrEmptyValue
	}

	mpID := MenteeRecruitmentPlanID(id)

	return &mpID, nil
}
