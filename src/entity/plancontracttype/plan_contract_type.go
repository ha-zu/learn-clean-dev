package plancontracttype

import custerr "github.com/ha-zu/learn-clean-dev/src/entity/customerror"

type ContractType uint64

const (
	ONE_TIME ContractType = iota + 1
	CONTINUOUS
)

func NewPlanContractType(contractType ContractType) error {

	if contractType != ONE_TIME && contractType != CONTINUOUS {
		return custerr.ErrOutOfRange
	}

	return nil

}

func ChangePlanContractType(contractType ContractType) (ContractType, error) {

	if contractType != ONE_TIME && contractType != CONTINUOUS {
		return contractType, custerr.ErrOutOfRange
	}

	if contractType == ONE_TIME {
		return CONTINUOUS, nil
	}

	return ONE_TIME, nil

}
