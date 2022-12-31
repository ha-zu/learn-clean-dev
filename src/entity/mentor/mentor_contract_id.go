package mentor

import custerr "github.com/ha-zu/learn-clean-dev/src/entity/customerror"

type MentorContractID string

func MentorContractIDValidate(id string) (*MentorContractID, error) {

	if id == "" {
		return nil, custerr.ErrEmptyValue
	}

	mcID := MentorContractID(id)

	return &mcID, nil
}
