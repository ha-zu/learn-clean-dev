package user

import (
	"time"
	"unicode/utf8"

	custerr "github.com/ha-zu/learn-clean-dev/src/entity/customerror"
)

type Resume struct {
	id          ResumeID
	userID      UserID
	description string
	from        int
	to          int
}

const (
	RESUME_DESCRIPTION_MAX_LEN = 1000
	RESUME_BASE_YEAR           = 1970
)

func NewResume(id, description string, userID UserID, from, to int) (*Resume, error) {

	rID, err := ResumeIDValid(id)
	if err != nil {
		return nil, err
	}

	if utf8.RuneCountInString(description) > RESUME_DESCRIPTION_MAX_LEN {
		return nil, custerr.ErrValueIsTooLong
	}

	if from < RESUME_BASE_YEAR {
		return nil, custerr.ErrOutOfRange
	}

	if to < RESUME_BASE_YEAR || from < to {
		return nil, custerr.ErrOutOfRange
	}

	return &Resume{
		id:          *rID,
		userID:      userID,
		description: description,
		from:        from,
		to:          to,
	}, nil
}

func (r *Resume) ChangeDesctiption(desc string, ud time.Time) error {

	if utf8.RuneCountInString(desc) > RESUME_DESCRIPTION_MAX_LEN {
		return custerr.ErrValueIsTooLong
	}

	r.description = desc

	return nil
}

func (r *Resume) ChangeFrom(from int, ud time.Time) error {

	if from < RESUME_BASE_YEAR {
		return custerr.ErrOutOfRange
	}

	r.from = from

	return nil
}

func (r *Resume) ChangeTo(to int, ud time.Time) error {

	if to < RESUME_BASE_YEAR || r.from < to {
		return custerr.ErrOutOfRange
	}

	r.to = to

	return nil
}
