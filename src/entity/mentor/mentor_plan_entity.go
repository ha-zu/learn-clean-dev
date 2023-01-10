package mentor

import (
	"unicode/utf8"

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

	mp := &MentorPlan{
		id:           id,
		userID:       uID,
		categoryID:   cID,
		tag:          tag,
		contractType: contractType,
		consultType:  consultType,
		status:       stat,
		proposal:     proposal,
		contract:     contract,
	}

	err := mp.ValidateMentorPlanTitle(title)
	if err != nil {
		return nil, err
	}

	err = mp.ValidateMentorPlanDescription(desc)
	if err != nil {
		return nil, err
	}

	err = mp.ValidateMentorPlanPrice(price)
	if err != nil {
		return nil, err
	}

	return mp, nil

}

func (m *MentorPlan) ValidateMentorPlanTitle(title string) error {

	if title == "" {
		return custerr.ErrEmptyValue
	}

	if l := utf8.RuneCountInString(title); l > MENTOR_PLAN_TITLE {
		return custerr.ErrValueIsTooLong
	}

	m.title = title

	return nil

}

func (m *MentorPlan) ValidateMentorPlanDescription(desc string) error {

	if l := utf8.RuneCountInString(desc); l > MENTOR_PLAN_DESCRIPTION {
		return custerr.ErrValueIsTooLong
	}

	m.description = desc

	return nil

}

func (m *MentorPlan) ValidateMentorPlanPrice(price uint64) error {

	if price < MENTOR_PLAN_BASE_PRICE {
		return custerr.ErrNotEnoughValue
	}

	m.price = price

	return nil

}
