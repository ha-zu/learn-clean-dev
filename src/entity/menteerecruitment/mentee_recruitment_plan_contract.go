package menteerecruitment

import (
	"unicode/utf8"

	custerr "github.com/ha-zu/learn-clean-dev/src/entity/customerror"
	"github.com/ha-zu/learn-clean-dev/src/entity/user"
)

type MenteeRecruitentPlanContract struct {
	id                     MenteeRecruitmentPlanContractID
	menteeRecruitentPlanID MenteeRecruitmentPlanID
	mentorID               user.UserID
	message                string
}

const MENTEE_PLAN_CONTRACT_MESSAGE_MAX_LEN = 500

func NewMenteeRecruitentPlanContract(id MenteeRecruitmentPlanContractID, message string, mtrpID MenteeRecruitmentPlanID, mentorID user.UserID) (*MenteeRecruitentPlanContract, error) {

	err := NewMenteeRecruitmentPlanContractID(id)
	if err != nil {
		return nil, err
	}

	if l := utf8.RuneCountInString(message); l > MENTEE_PLAN_CONTRACT_MESSAGE_MAX_LEN {
		return nil, custerr.ErrValueIsTooLong
	}

	return &MenteeRecruitentPlanContract{
		id:                     id,
		menteeRecruitentPlanID: mtrpID,
		mentorID:               mentorID,
		message:                message,
	}, nil
}

func (m *MenteeRecruitentPlanContract) ChangeMenteeRecruitentPlanContractMessage(message string) error {

	if l := utf8.RuneCountInString(message); l > MENTEE_PLAN_CONTRACT_MESSAGE_MAX_LEN {
		return custerr.ErrValueIsTooLong
	}

	m.message = message

	return nil
}
