package entity

import (
	"unicode/utf8"

	"github.com/ha-zu/learn-clean-dev/src/entity/vo"
	mr "github.com/ha-zu/learn-clean-dev/src/entity/vo/mentee_recruitment"
	"github.com/ha-zu/learn-clean-dev/src/entity/vo/user"
)

type MenteeRecruitentPlanContract struct {
	id                     mr.MenteeRecruitmentPlanContractID
	menteeRecruitentPlanID mr.MenteeRecruitmentPlanID
	mentorID               user.UserID
	message                string
}

const MENTEE_PLAN_CONTRACT_MESSAGE_MAX_LEN = 500

func NewMenteeRecruitentPlanContract(id, message string, mtrpID mr.MenteeRecruitmentPlanID, mentorID user.UserID) (*MenteeRecruitentPlanContract, error) {

	mrpcID, err := mr.MenteeRecruitmentPlanContractIDValidate(id)
	if err != nil {
		return nil, err
	}

	if l := utf8.RuneCountInString(message); l > MENTEE_PLAN_CONTRACT_MESSAGE_MAX_LEN {
		return nil, vo.ErrValueIsTooLong
	}

	return &MenteeRecruitentPlanContract{
		id:                     *mrpcID,
		menteeRecruitentPlanID: mtrpID,
		mentorID:               mentorID,
		message:                message,
	}, nil
}

func (m *MenteeRecruitentPlanContract) ChangeMenteeRecruitentPlanContractMessage(message string) error {

	if l := utf8.RuneCountInString(message); l > MENTEE_PLAN_CONTRACT_MESSAGE_MAX_LEN {
		return vo.ErrValueIsTooLong
	}

	m.message = message

	return nil
}
