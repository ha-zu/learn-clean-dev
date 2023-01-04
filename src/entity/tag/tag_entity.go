package tag

import custerr "github.com/ha-zu/learn-clean-dev/src/entity/customerror"

type Tag struct {
	id   TagID
	name TagName
}

func NewTag(id TagID, name TagName) (*Tag, error) {

	err := NewTagID(id)
	if err != nil {
		return nil, custerr.ErrEmptyValue
	}

	err = NewTagName(name)
	if err != nil {
		return nil, err
	}

	return &Tag{
		id:   id,
		name: name,
	}, nil
}
