package vo

import "errors"

type Skill struct {
	id       string
	userID   string
	tagID    string
	evaluate int
	year     int
}

func SkillVerify(id, user_id, tag_id string, evaluate, year int) (*Skill, error) {

	if id == "" {
		return nil, errors.New("must not be empty")
	}

	if user_id == "" {
		return nil, errors.New("must not be empty")
	}

	if tag_id == "" {
		return nil, errors.New("must not be empty")
	}

	//ToDo other

	return &Skill{
		id:       id,
		userID:   user_id,
		tagID:    tag_id,
		evaluate: evaluate,
		year:     year,
	}, nil
}
