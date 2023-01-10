package menteerecruitment

import (
	"github.com/google/uuid"
	custerr "github.com/ha-zu/learn-clean-dev/src/entity/customerror"
)

type MenteeRecruitmentPlanID string

func NewMenteePlanID() MenteeRecruitmentPlanID {
	return MenteeRecruitmentPlanID(uuid.New().String())
}

func NewMenteePlanIDStr(id string) (MenteeRecruitmentPlanID, error) {

	if id == "" {
		return MenteeRecruitmentPlanID(""), custerr.ErrEmptyValue
	}

	return MenteeRecruitmentPlanID(id), nil

}
