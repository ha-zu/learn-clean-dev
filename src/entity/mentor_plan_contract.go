package entity

import (
	"unicode/utf8"

	"github.com/ha-zu/learn-clean-dev/src/entity/vo"
	mr "github.com/ha-zu/learn-clean-dev/src/entity/vo/mentor"
	"github.com/ha-zu/learn-clean-dev/src/entity/vo/user"
)

type MentorContract struct {
	id           mr.MentorContractID
	mentorPlanID mr.MentorPlanID
	menteeID     user.UserID
	message      string
}

const MENTOR_CONTRACT_MESSAGE_MAX_LEN = 500

func NewMentorContract(id, message string, mentorPlanID mr.MentorPlanID, menteeID user.UserID) (*MentorContract, error) {

	mcID, err := mr.MentorContractIDValidate(id)
	if err != nil {
		return nil, err
	}

	if l := utf8.RuneCountInString(message); l > MENTOR_CONTRACT_MESSAGE_MAX_LEN {
		return nil, vo.ErrValueIsTooLong
	}

	return &MentorContract{
		id:           *mcID,
		mentorPlanID: mentorPlanID,
		menteeID:     menteeID,
		message:      message,
	}, nil
}

func (m *MentorContract) ChangeMentorPlanContractMessage(message string) error {

	if l := utf8.RuneCountInString(message); l > MENTOR_CONTRACT_MESSAGE_MAX_LEN {
		return vo.ErrValueIsTooLong
	}

	m.message = message

	return nil
}
