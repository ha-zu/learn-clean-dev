package user

import (
	custerr "github.com/ha-zu/learn-clean-dev/src/entity/customerror"
	"github.com/ha-zu/learn-clean-dev/src/entity/tag"
)

type Skill struct {
	id       SkillID
	userID   UserID
	tagID    tag.TagID
	evaluate uint64
	year     uint64
}

func NewSkill(id SkillID, userID UserID, tagID tag.TagID, evaluate, year uint64) (*Skill, error) {

	err := NewSkillID(id)
	if err != nil {
		return nil, err
	}

	if evaluate < 1 || evaluate > 5 {
		return nil, custerr.ErrOutOfRange
	}

	if year > 5 {
		return nil, custerr.ErrOutOfRange
	}

	return &Skill{
		id:       id,
		userID:   userID,
		tagID:    tagID,
		evaluate: evaluate,
		year:     year,
	}, nil
}

func (s *Skill) ChangeEvaluate(evaluate uint64) error {

	if evaluate < 1 || evaluate > 5 {
		return custerr.ErrOutOfRange
	}

	s.evaluate = evaluate

	return nil
}

func (s *Skill) ChangeYear(year uint64) error {

	if year > 5 {
		return custerr.ErrOutOfRange
	}

	s.year = year

	return nil
}
