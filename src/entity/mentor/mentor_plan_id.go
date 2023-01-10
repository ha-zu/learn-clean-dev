package mentor

import (
	"github.com/google/uuid"
	custerr "github.com/ha-zu/learn-clean-dev/src/entity/customerror"
)

type MentorPlanID string

func NewMentorPlanID() MentorPlanID {
	return MentorPlanID(uuid.New().String())
}

func NewMentorPlanIDStr(id string) (MentorPlanID, error) {

	if id == "" {
		return MentorPlanID(""), custerr.ErrEmptyValue
	}

	return MentorPlanID(id), nil

}
