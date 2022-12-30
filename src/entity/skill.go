package entity

import (
	"github.com/ha-zu/learn-clean-dev/src/entity/vo"
	"github.com/ha-zu/learn-clean-dev/src/entity/vo/skill"
	"github.com/ha-zu/learn-clean-dev/src/entity/vo/tag"
	"github.com/ha-zu/learn-clean-dev/src/entity/vo/user"
)

type Skill struct {
	id       skill.SkillID
	userID   user.UserID
	tagID    tag.TagID
	evaluate int
	year     int
}

func NewSkill(id string, tagID tag.TagID, userID user.UserID, evaluate, year int) (*Skill, error) {

	sID, err := skill.SkillIDValid(id)
	if err != nil {
		return nil, err
	}

	if evaluate < 1 || evaluate > 5 {
		return nil, vo.ErrOutOfRange
	}

	if year < 0 || year > 5 {
		return nil, vo.ErrOutOfRange
	}

	return &Skill{
		id:       *sID,
		userID:   userID,
		tagID:    tagID,
		evaluate: evaluate,
		year:     year,
	}, nil
}

func (s *Skill) ChangeEvaluatee(evaluate int) error {

	if evaluate < 1 || evaluate > 5 {
		return vo.ErrOutOfRange
	}

	s.evaluate = evaluate

	return nil
}

func (s *Skill) ChangeYear(year int) error {

	if year < 0 || year > 5 {
		return vo.ErrOutOfRange
	}

	s.year = year

	return nil
}
