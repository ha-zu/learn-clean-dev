package tag

import custerr "github.com/ha-zu/learn-clean-dev/src/entity/customerror"

type TagName string

func NewTagName(name TagName) error {

	if name == "" {
		return custerr.ErrEmptyValue
	}

	return nil

}
