package mentor

import (
	"unicode/utf8"

	"github.com/ha-zu/learn-clean-dev/src/entity"
	ctg "github.com/ha-zu/learn-clean-dev/src/entity/category"
	custerr "github.com/ha-zu/learn-clean-dev/src/entity/customerror"
	"github.com/ha-zu/learn-clean-dev/src/entity/tag"
	"github.com/ha-zu/learn-clean-dev/src/entity/user"
)

type MentorPlan struct {
	id           MentorPlanID
	userID       user.UserID
	title        string
	categoryID   ctg.CategoryID
	tag          []tag.Tag
	description  string
	status       string
	price        int
	contractType string
	consultType  string
	proposal     []MentorProposal
	contract     []MentorContract
}

const (
	MENTOR_PLAN_TITLE       = 255
	MENTOR_PLAN_DESCRIPTION = 2000
)

func NewMentorPlan(id, title, desc, stat, contractType, consultType string, price int,
	uID user.UserID, cID ctg.CategoryID, tag []tag.Tag, proposal []MentorProposal, contract []MentorContract) (*MentorPlan, error) {

	mID, err := MentorPlanIDValidate(id)
	if err != nil {
		return nil, err
	}

	if title == "" {
		return nil, custerr.ErrEmptyValue
	}

	if l := utf8.RuneCountInString(title); l > MENTOR_PLAN_TITLE {
		return nil, custerr.ErrValueIsTooLong
	}

	if l := utf8.RuneCountInString(desc); l > MENTOR_PLAN_DESCRIPTION {
		return nil, custerr.ErrValueIsTooLong
	}

	if stat == "" {
		return nil, custerr.ErrEmptyValue
	}

	if stat == entity.PLAN_DEFAULT_STAT {
		return nil, custerr.ErrNotDefaultValue
	}

	if contractType == "" {
		return nil, custerr.ErrEmptyValue
	}

	if contractType != entity.PLAN_CONTRACT_ONECE && contractType != entity.PLAN_CONTRACT_CONTINUTION {
		return nil, custerr.ErrNotMatchValue
	}

	if consultType == "" {
		return nil, custerr.ErrEmptyValue
	}

	if consultType != entity.PLAN_CONTRACT_ONECE && consultType != entity.PLAN_CONTRACT_CONTINUTION {
		return nil, custerr.ErrNotMatchValue
	}

	return &MentorPlan{
		id:           *mID,
		userID:       uID,
		title:        title,
		categoryID:   cID,
		tag:          tag,
		description:  desc,
		status:       stat,
		price:        price,
		contractType: contractType,
		consultType:  consultType,
		proposal:     proposal,
		contract:     contract,
	}, nil

}

func (m *MentorPlan) ChangeMentorPlanTitle(title string) error {

	if title == "" {
		return custerr.ErrEmptyValue
	}

	if l := utf8.RuneCountInString(title); l > MENTOR_PLAN_TITLE {
		return custerr.ErrValueIsTooLong
	}

	m.title = title

	return nil
}

func (m *MentorPlan) ChangeMentorPlanDescription(desc string) error {

	if l := utf8.RuneCountInString(desc); l > MENTOR_PLAN_DESCRIPTION {
		return custerr.ErrValueIsTooLong
	}

	m.description = desc

	return nil
}

func (m *MentorPlan) ChangeMentorPlanStatus(stat string) error {

	if stat == "" {
		return custerr.ErrEmptyValue
	}

	if stat == entity.PLAN_DEFAULT_STAT || stat == entity.PLAN_STOP_STAT {
		return custerr.ErrNotCorrectFormat
	}

	if stat == entity.PLAN_DEFAULT_STAT {
		m.status = entity.PLAN_STOP_STAT
		return nil
	}

	m.status = entity.PLAN_DEFAULT_STAT

	return nil

}

func (m *MentorPlan) ChangeMentorPlanPrice(price int) error {

	if price < entity.PLAN_PRICE_BASE {
		return custerr.ErrNotEnoughValue
	}

	m.price = price

	return nil
}

func (m *MentorPlan) ChangeMentorPlanContractType(contractType string) error {

	if contractType == "" {
		return custerr.ErrEmptyValue
	}

	if contractType != entity.PLAN_CONTRACT_ONECE && contractType != entity.PLAN_CONTRACT_CONTINUTION {
		return custerr.ErrNotMatchValue
	}

	m.contractType = contractType

	return nil

}

func (m *MentorPlan) ChangeMentorPlanConsultType(consultType string) error {

	if consultType == "" {
		return custerr.ErrEmptyValue
	}

	if consultType != entity.PLAN_CONTRACT_ONECE && consultType != entity.PLAN_CONTRACT_CONTINUTION {
		return custerr.ErrNotMatchValue
	}

	m.consultType = consultType

	return nil

}
