package menteerecruitment

import (
	"github.com/google/uuid"
	custerr "github.com/ha-zu/learn-clean-dev/src/entity/customerror"
)

type MenteeRecruitmentPlanProposalID string

func NewMenteeRecruitmentPlanProposalID() MenteeRecruitmentPlanProposalID {
	return MenteeRecruitmentPlanProposalID(uuid.New().String())
}

func NewMenteeRecruitmentPlanProposalIDStr(id string) (MenteeRecruitmentPlanProposalID, error) {

	if id == "" {
		return MenteeRecruitmentPlanProposalID(""), custerr.ErrEmptyValue
	}

	return MenteeRecruitmentPlanProposalID(id), nil

}
