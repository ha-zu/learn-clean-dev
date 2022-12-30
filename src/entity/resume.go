package entity

import (
	"time"
	"unicode/utf8"

	"github.com/ha-zu/learn-clean-dev/src/entity/vo"
	"github.com/ha-zu/learn-clean-dev/src/entity/vo/resume"
	"github.com/ha-zu/learn-clean-dev/src/entity/vo/user"
)

type Resume struct {
	id          resume.ResumeID
	userID      user.UserID
	description string
	from        int
	to          int
}

const (
	DESCRIPTION_MAX_LEN = 1000
	BASE_YEAR           = 1970
)

func NewResume(id, description string, userID user.UserID, from, to int) (*Resume, error) {

	rID, err := resume.ResumeIDValid(id)
	if err != nil {
		return nil, err
	}

	if utf8.RuneCountInString(description) > DESCRIPTION_MAX_LEN {
		return nil, vo.ErrValueIsTooLong
	}

	if from < BASE_YEAR {
		return nil, vo.ErrOutOfRange
	}

	if to < BASE_YEAR || from < to {
		return nil, vo.ErrOutOfRange
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

	if utf8.RuneCountInString(desc) > DESCRIPTION_MAX_LEN {
		return vo.ErrValueIsTooLong
	}

	r.description = desc

	return nil
}

func (r *Resume) ChangeFrom(from int, ud time.Time) error {

	if from < BASE_YEAR {
		return vo.ErrOutOfRange
	}

	r.from = from

	return nil
}

func (r *Resume) ChangeTo(to int, ud time.Time) error {

	if to < BASE_YEAR || r.from < to {
		return vo.ErrOutOfRange
	}

	r.to = to

	return nil
}
