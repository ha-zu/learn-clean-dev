package user

import (
	custerr "github.com/ha-zu/learn-clean-dev/src/entity/customerror"
	"github.com/ha-zu/learn-clean-dev/src/entity/tag"
)

type Skill struct {
	id       SkillID
	userID   UserID
	tagID    tag.TagID
	evaluate int
	year     int
}

func NewSkill(id string, tagID tag.TagID, userID UserID, evaluate, year int) (*Skill, error) {

	sID, err := SkillIDValid(id)
	if err != nil {
		return nil, err
	}

	if evaluate < 1 || evaluate > 5 {
		return nil, custerr.ErrOutOfRange
	}

	if year < 0 || year > 5 {
		return nil, custerr.ErrOutOfRange
	}

	return &Skill{
		id:       *sID,
		userID:   userID,
		tagID:    tagID,
		evaluate: evaluate,
		year:     year,
	}, nil
}

func (s *Skill) ChangeEvaluate(evaluate int) error {

	if evaluate < 1 || evaluate > 5 {
		return custerr.ErrOutOfRange
	}

	s.evaluate = evaluate

	return nil
}

func (s *Skill) ChangeYear(year int) error {

	if year < 0 || year > 5 {
		return custerr.ErrOutOfRange
	}

	s.year = year

	return nil
}
