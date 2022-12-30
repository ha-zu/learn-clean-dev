package mentor

import "github.com/ha-zu/learn-clean-dev/src/entity/vo"

type MentorContractID string

func MentorContractIDValidate(id string) (*MentorContractID, error) {

	if id == "" {
		return nil, vo.ErrEmptyValue
	}

	mcID := MentorContractID(id)

	return &mcID, nil
}
