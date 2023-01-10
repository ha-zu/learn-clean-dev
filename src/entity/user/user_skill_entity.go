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

	s := &Skill{
		id:     id,
		userID: userID,
		tagID:  tagID,
	}

	err := s.ValidateEvaluate(evaluate)
	if err != nil {
		return nil, err
	}

	err = s.ValidateYear(year)
	if err != nil {
		return nil, err
	}

	return s, nil
}

func (s *Skill) ValidateEvaluate(evaluate uint64) error {

	if evaluate < 1 || evaluate > 5 {
		return custerr.ErrOutOfRange
	}

	s.evaluate = evaluate

	return nil
}

func (s *Skill) ValidateYear(year uint64) error {

	if year > 5 {
		return custerr.ErrOutOfRange
	}

	s.year = year

	return nil
}
