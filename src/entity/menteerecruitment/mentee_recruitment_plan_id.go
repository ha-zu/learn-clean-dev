package menteerecruitment

import custerr "github.com/ha-zu/learn-clean-dev/src/entity/customerror"

type MenteeRecruitmentPlanID string

func NewMenteePlanID(id MenteeRecruitmentPlanID) error {

	if id == "" {
		return custerr.ErrEmptyValue
	}

	return nil
}
