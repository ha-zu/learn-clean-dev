package tag

import (
	"github.com/ha-zu/learn-clean-dev/src/entity/vo"
)

type TagName string

func TagNameValidate(tag_name string) (*TagName, error) {

	if tag_name == "" {
		return nil, vo.ErrEmptyValue
	}

	//ToDo: What to do about tag name length

	tName := TagName(tag_name)

	return &tName, nil
}
