package category

import (
	"github.com/google/uuid"
	custerr "github.com/ha-zu/learn-clean-dev/src/entity/customerror"
)

type CategoryID string

func NewCategoryID() CategoryID {

	return CategoryID(uuid.New().String())

}

func NewCategoryIDByStr(id string) (CategoryID, error) {

	if id == "" {
		return CategoryID(""), custerr.ErrEmptyValue
	}

	return CategoryID(id), nil

}
