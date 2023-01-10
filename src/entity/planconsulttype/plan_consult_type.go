package planconsulttype

import custerr "github.com/ha-zu/learn-clean-dev/src/entity/customerror"

type ConsultType uint64

const (
	CHAT ConsultType = iota + 1
	TELL
)

func NewPlanConsultType(consultType ConsultType) error {

	if consultType != CHAT && consultType != TELL {
		return custerr.ErrOutOfRange
	}

	return nil

}

func ChangePlanConsultType(consultType ConsultType) (ConsultType, error) {

	if consultType != CHAT && consultType != TELL {
		return consultType, custerr.ErrOutOfRange
	}

	if consultType == CHAT {
		return TELL, nil
	}

	return CHAT, nil

}
