package menteerecruitment

import (
	"github.com/google/uuid"
	custerr "github.com/ha-zu/learn-clean-dev/src/entity/customerror"
)

type MenteeRecruitmentPlanContractID string

func NewMenteeRecruitmentPlanContractID() MenteeRecruitmentPlanContractID {

	return MenteeRecruitmentPlanContractID(uuid.New().String())

}

func NewMenteeRecruitmentPlanContractIDByStr(id string) (MenteeRecruitmentPlanContractID, error) {

	if id == "" {
		return MenteeRecruitmentPlanContractID(""), custerr.ErrEmptyValue
	}

	return MenteeRecruitmentPlanContractID(id), nil
}
