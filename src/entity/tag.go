package entity

import (
	"github.com/ha-zu/learn-clean-dev/src/entity/vo"
	"github.com/ha-zu/learn-clean-dev/src/entity/vo/tag"
)

type Tag struct {
	id      tag.TagID
	tagName tag.TagName
}

func NewTag(id, tag_name string) (*Tag, error) {

	tID, err := tag.TagIDValid(id)
	if err != nil {
		return nil, vo.ErrEmptyValue
	}

	tName, err := tag.TagNameValidate(tag_name)
	if err != nil {
		return nil, err
	}

	return &Tag{
		id:      *tID,
		tagName: *tName,
	}, nil
}
