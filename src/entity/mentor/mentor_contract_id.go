package mentor

import custerr "github.com/ha-zu/learn-clean-dev/src/entity/customerror"

type MentorContractID string

func NewMentorContractID(id MentorContractID) error {

	if id == "" {
		return custerr.ErrEmptyValue
	}

	return nil

}
