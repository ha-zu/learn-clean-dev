package planstatus

import custerr "github.com/ha-zu/learn-clean-dev/src/entity/customerror"

type PlanStatus uint64

const (
	PUBLISH PlanStatus = iota + 1 //open
	CLOSE                         //close(cancel)
)

func NewPlanStatus(stat PlanStatus) error {

	if stat != PUBLISH && stat != CLOSE {
		return custerr.ErrOutOfRange
	}

	return nil

}

func ChangePlanStatus(stat PlanStatus) (PlanStatus, error) {

	if stat != PUBLISH && stat != CLOSE {
		return stat, custerr.ErrOutOfRange
	}

	if stat == PUBLISH {
		return CLOSE, nil
	}

	return PUBLISH, nil

}
