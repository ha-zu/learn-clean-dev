package entity

import (
	"unicode/utf8"

	"github.com/ha-zu/learn-clean-dev/src/entity/vo"
	ctg "github.com/ha-zu/learn-clean-dev/src/entity/vo/category"
	mr "github.com/ha-zu/learn-clean-dev/src/entity/vo/mentor"
	"github.com/ha-zu/learn-clean-dev/src/entity/vo/user"
)

type MentorPlan struct {
	id           mr.MentorPlanID
	userID       user.UserID
	title        string
	categoryID   ctg.CategoryID
	tag          []Tag
	description  string
	status       string
	contractType string
	price        int
	consultType  string
	proposal     []MentorProposal
	contract     []MentorContract
}

const (
	MENTOR_PLAN_TITLE       = 255
	MENTOR_PLAN_DESCRIPTION = 2000
)

func NewMentorPlan(id, title, desc, stat, contractType, consultType string,
	uID user.UserID, cID ctg.CategoryID, tag []Tag, proposal []MentorProposal, contract []MentorContract) (*MentorPlan, error) {

	mID, err := mr.MentorPlanIDValidate(id)
	if err != nil {
		return nil, err
	}

	if title == "" {
		return nil, vo.ErrEmptyValue
	}

	if l := utf8.RuneCountInString(title); l > MENTOR_PLAN_TITLE {
		return nil, vo.ErrValueIsTooLong
	}

	if l := utf8.RuneCountInString(desc); l > MENTOR_PLAN_DESCRIPTION {
		return nil, vo.ErrValueIsTooLong
	}

	if stat == "" {
		return nil, vo.ErrEmptyValue
	}

	if stat == vo.PLAN_DEFAULT_STAT {
		return nil, vo.ErrNotDefaultValue
	}

	if contractType == "" {
		return nil, vo.ErrEmptyValue
	}

	if contractType != vo.PLAN_CONTRACT_ONECE && consultType != vo.PLAN_CONTRACT_CONTINUTION {
		return nil, vo.ErrNotMatchValue
	}

	if consultType == "" {
		return nil, vo.ErrEmptyValue
	}

	if consultType != vo.PLAN_CONTRACT_ONECE && consultType != vo.PLAN_CONTRACT_CONTINUTION {
		return nil, vo.ErrNotMatchValue
	}

	return &MentorPlan{
		id:           *mID,
		userID:       uID,
		title:        title,
		categoryID:   cID,
		tag:          tag,
		description:  desc,
		contractType: contractType,
		consultType:  consultType,
		proposal:     proposal,
		contract:     contract,
	}, nil

}

func (m *MentorPlan) ChangeMentorPlanTitle(title string) error {

	if title == "" {
		return vo.ErrEmptyValue
	}

	if l := utf8.RuneCountInString(title); l > MENTOR_PLAN_TITLE {
		return vo.ErrValueIsTooLong
	}

	m.title = title

	return nil
}

func (m *MentorPlan) ChangeMentorPlanPrice(price int) error {

	if price < vo.PLAN_PRICE_BASE {
		return vo.ErrNotEnoughValue
	}

	m.price = price

	return nil
}

func (m *MentorPlan) ChangeStatus(stat string) error {

	if stat == "" {
		return vo.ErrEmptyValue
	}

	if stat == vo.PLAN_DEFAULT_STAT || stat == vo.PLAN_STOP_STAT {
		return vo.ErrNotCorrectFormat
	}

	if stat == vo.PLAN_DEFAULT_STAT {
		m.status = vo.PLAN_STOP_STAT
		return nil
	}

	m.status = vo.PLAN_DEFAULT_STAT

	return nil

}
