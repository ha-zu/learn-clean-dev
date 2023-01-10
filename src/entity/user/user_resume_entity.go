package user

import (
	"unicode/utf8"

	custerr "github.com/ha-zu/learn-clean-dev/src/entity/customerror"
)

type Resume struct {
	id          ResumeID
	userID      UserID
	description string
	from        uint64
	to          uint64
}

const (
	RESUME_DESCRIPTION_MAX_LEN = 1000
	RESUME_BASE_YEAR           = 1970
)

func NewResume(id ResumeID, userID UserID, desc string, from, to uint64) (*Resume, error) {

	r := &Resume{
		id:     id,
		userID: userID,
	}

	err := r.ValidateDesctiption(desc)
	if err != nil {
		return nil, err
	}

	err = r.ValidateFromTo(from, to)
	if err != nil {
		return nil, err
	}

	return r, nil

}

func (r *Resume) ValidateDesctiption(desc string) error {

	if utf8.RuneCountInString(desc) > RESUME_DESCRIPTION_MAX_LEN {
		return custerr.ErrValueIsTooLong
	}

	r.description = desc

	return nil

}

func (r *Resume) ValidateFromTo(from, to uint64) error {

	if from < RESUME_BASE_YEAR {
		return custerr.ErrOutOfRange
	}

	if to < RESUME_BASE_YEAR || r.from < to {
		return custerr.ErrOutOfRange
	}

	r.from = from
	r.to = to

	return nil

}
