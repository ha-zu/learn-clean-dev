package mentor

import (
	"unicode/utf8"

	"github.com/ha-zu/learn-clean-dev/src/entity"
	ctg "github.com/ha-zu/learn-clean-dev/src/entity/category"
	custerr "github.com/ha-zu/learn-clean-dev/src/entity/customerror"
	cs "github.com/ha-zu/learn-clean-dev/src/entity/planconsulttype"
	ct "github.com/ha-zu/learn-clean-dev/src/entity/plancontracttype"
	st "github.com/ha-zu/learn-clean-dev/src/entity/planstatus"
	"github.com/ha-zu/learn-clean-dev/src/entity/tag"
	"github.com/ha-zu/learn-clean-dev/src/entity/user"
)

type MentorPlan struct {
	id           MentorPlanID
	userID       user.UserID
	title        string
	description  string
	categoryID   ctg.CategoryID
	tag          []tag.Tag
	contractType ct.ContractType
	consultType  cs.ConsultType
	status       st.PlanStatus
	price        uint64
	proposal     []MentorProposal
	contract     []MentorContract
}

const (
	MENTOR_PLAN_TITLE       = 255
	MENTOR_PLAN_DESCRIPTION = 2000
	MENTOR_PLAN_BASE_PRICE  = 1000
)

func NewMentorPlan(id MentorPlanID, uID user.UserID, title, desc string, cID ctg.CategoryID, tag []tag.Tag,
	stat st.PlanStatus, contractType ct.ContractType, consultType cs.ConsultType, price uint64,
	proposal []MentorProposal, contract []MentorContract) (*MentorPlan, error) {

	err := NewMentorPlanID(id)
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

	err = ct.NewPlanContractType(contractType)
	if err != nil {
		return nil, err
	}

	err = cs.NewPlanConsultType(consultType)
	if err != nil {
		return nil, err
	}

	err = st.NewPlanStatus(stat)
	if err != nil {
		return nil, err
	}

	if price < MENTOR_PLAN_BASE_PRICE {
		return nil, custerr.ErrOutOfRange
	}

	return &MentorPlan{
		id:           id,
		userID:       uID,
		title:        title,
		description:  desc,
		categoryID:   cID,
		tag:          tag,
		contractType: contractType,
		consultType:  consultType,
		status:       stat,
		price:        price,
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

func (m *MentorPlan) ChangeMentorPlanPrice(price uint64) error {

	if price < entity.PLAN_PRICE_BASE {
		return custerr.ErrNotEnoughValue
	}

	m.price = price

	return nil

}
