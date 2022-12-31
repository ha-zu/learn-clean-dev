package tag

import custerr "github.com/ha-zu/learn-clean-dev/src/entity/customerror"

type TagName string

func TagNameValidate(tag_name string) (*TagName, error) {

	if tag_name == "" {
		return nil, custerr.ErrEmptyValue
	}

	tName := TagName(tag_name)

	return &tName, nil
}
