package vo

import "errors"

type Resume struct {
	id          string
	userID      string
	description string
	from        int
	to          int
}

func ResumeVerify(id, user_id, description string, from, to int) (*Resume, error) {

	if id == "" {
		return nil, errors.New("must not be empty")
	}

	if user_id == "" {
		return nil, errors.New("must not be empty")
	}

	//ToDo other

	return &Resume{
		id:          id,
		userID:      user_id,
		description: description,
		from:        from,
		to:          to,
	}, nil
}
