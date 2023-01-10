package tag

import custerr "github.com/ha-zu/learn-clean-dev/src/entity/customerror"

type TagName string

func NewTagName(name string) (TagName, error) {

	if name == "" {
		return TagName(""), custerr.ErrEmptyValue
	}

	return TagName(name), nil

}
