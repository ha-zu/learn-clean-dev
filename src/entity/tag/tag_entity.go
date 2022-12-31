package tag

import custerr "github.com/ha-zu/learn-clean-dev/src/entity/customerror"

type Tag struct {
	id      TagID
	tagName TagName
}

func NewTag(id, tag_name string) (*Tag, error) {

	tID, err := TagIDValid(id)
	if err != nil {
		return nil, custerr.ErrEmptyValue
	}

	tName, err := TagNameValidate(tag_name)
	if err != nil {
		return nil, err
	}

	return &Tag{
		id:      *tID,
		tagName: *tName,
	}, nil
}
