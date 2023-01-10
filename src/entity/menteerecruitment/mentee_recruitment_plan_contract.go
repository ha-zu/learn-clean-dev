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

func NewMenteeRecruitentPlanContract(id MenteeRecruitmentPlanContractID, mtrpID MenteeRecruitmentPlanID,
	mentorID user.UserID, message string) (*MenteeRecruitentPlanContract, error) {

	pc := &MenteeRecruitentPlanContract{
		id:                     id,
		menteeRecruitentPlanID: mtrpID,
		mentorID:               mentorID,
	}

	err := pc.ValidateMenteeRecruitentPlanContractMessage(message)
	if err != nil {
		return nil, err
	}

	return pc, nil

}

func (m *MenteeRecruitentPlanContract) ValidateMenteeRecruitentPlanContractMessage(message string) error {

	if l := utf8.RuneCountInString(message); l > MENTEE_PLAN_CONTRACT_MESSAGE_MAX_LEN {
		return custerr.ErrValueIsTooLong
	}

	m.message = message

	return nil

}
