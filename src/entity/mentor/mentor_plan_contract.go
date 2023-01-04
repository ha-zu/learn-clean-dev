package mentor

import (
	"unicode/utf8"

	custerr "github.com/ha-zu/learn-clean-dev/src/entity/customerror"
	"github.com/ha-zu/learn-clean-dev/src/entity/user"
)

type MentorContract struct {
	id           MentorContractID
	mentorPlanID MentorPlanID
	menteeID     user.UserID
	message      string
}

const MENTOR_CONTRACT_MESSAGE_MAX_LEN = 500

func NewMentorContract(id MentorContractID, mentorPlanID MentorPlanID, menteeID user.UserID, message string) (*MentorContract, error) {

	err := NewMentorContractID(id)
	if err != nil {
		return nil, err
	}

	if l := utf8.RuneCountInString(message); l > MENTOR_CONTRACT_MESSAGE_MAX_LEN {
		return nil, custerr.ErrValueIsTooLong
	}

	return &MentorContract{
		id:           id,
		mentorPlanID: mentorPlanID,
		menteeID:     menteeID,
		message:      message,
	}, nil

}

func (m *MentorContract) ChangeMentorPlanContractMessage(message string) error {

	if l := utf8.RuneCountInString(message); l > MENTOR_CONTRACT_MESSAGE_MAX_LEN {
		return custerr.ErrValueIsTooLong
	}

	m.message = message

	return nil

}
