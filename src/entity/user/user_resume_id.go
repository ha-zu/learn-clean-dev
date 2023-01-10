package user

import (
	"github.com/google/uuid"
	custerr "github.com/ha-zu/learn-clean-dev/src/entity/customerror"
)

type ResumeID string

func NewResumeID() ResumeID {
	return ResumeID(uuid.New().String())
}

func NewResumeIDStr(id ResumeID) (ResumeID, error) {

	if id == "" {
		return ResumeID(""), custerr.ErrEmptyValue
	}

	return ResumeID(id), nil

}
