package menteerecruitment

import (
	"time"
	"unicode/utf8"

	ctg "github.com/ha-zu/learn-clean-dev/src/entity/category"
	custerr "github.com/ha-zu/learn-clean-dev/src/entity/customerror"
	cs "github.com/ha-zu/learn-clean-dev/src/entity/planconsulttype"
	ct "github.com/ha-zu/learn-clean-dev/src/entity/plancontracttype"
	st "github.com/ha-zu/learn-clean-dev/src/entity/planstatus"
	"github.com/ha-zu/learn-clean-dev/src/entity/user"
)

type MenteeRecruitmentPlan struct {
	id                 MenteeRecruitmentPlanID
	userID             user.UserID
	title              string
	categoryID         ctg.CategoryID
	contractType       ct.ContractType
	consultType        cs.ConsultType
	status             st.PlanStatus
	description        string
	priceFrom          uint64
	priceTo            uint64
	termFrom           time.Time
	termTo             time.Time
	menteeContractList []MenteeRecruitentPlanContract
	menteeProposalList []MenteeRecruitentPlanProposal
}

const (
	MENTEE_PLAN_DESCRIPTION_MAX_LEN = 2000
	MENTEE_PLAN_TITLE_MAX_LEN       = 255
	MENTEE_PLAN_MAX_DATE            = 14
	MENTEE_PLAN_BASE_PRICE          = 1000
)

func NewMenteeRecruitmentPlan(id MenteeRecruitmentPlanID, usrID user.UserID, title string, ctgID ctg.CategoryID,
	contractType ct.ContractType, consultType cs.ConsultType, stat st.PlanStatus,
	desc string, priceFrom, priceTo uint64, termFrom, termTo time.Time,
	contractList []MenteeRecruitentPlanContract, proposalList []MenteeRecruitentPlanProposal) (*MenteeRecruitmentPlan, error) {

	plan := &MenteeRecruitmentPlan{
		id:                 id,
		userID:             usrID,
		categoryID:         ctgID,
		contractType:       contractType,
		consultType:        consultType,
		status:             stat,
		menteeContractList: contractList,
		menteeProposalList: proposalList,
	}

	err := plan.ValidatesMenteeRecruitmentPlanTitle(title)
	if err != nil {
		return nil, err
	}

	err = plan.ValidateMenteeRecruitmentPlanDescription(desc)
	if err != nil {
		return nil, err
	}

	err = plan.ValidateMenteeRecruitmentPlanPrice(priceFrom, priceTo)
	if err != nil {
		return nil, err
	}

	err = plan.ValidateMenteeRecruitmentPlanTerm(termFrom, termTo)
	if err != nil {
		return nil, err
	}

	return plan, nil

}

func (m *MenteeRecruitmentPlan) ValidatesMenteeRecruitmentPlanTitle(title string) error {

	if title == "" {
		return custerr.ErrEmptyValue
	}

	if l := utf8.RuneCountInString(title); l > MENTEE_PLAN_TITLE_MAX_LEN {
		return custerr.ErrValueIsTooLong
	}

	m.title = title

	return nil

}

func (m *MenteeRecruitmentPlan) ValidateMenteeRecruitmentPlanDescription(desc string) error {

	if desc == "" {
		return custerr.ErrEmptyValue
	}

	if l := utf8.RuneCountInString(desc); l > MENTEE_PLAN_DESCRIPTION_MAX_LEN {
		return custerr.ErrValueIsTooLong
	}

	m.description = desc

	return nil

}

func (m *MenteeRecruitmentPlan) ValidateMenteeRecruitmentPlanPrice(priceFrom, priceTo uint64) error {

	if priceFrom < MENTEE_PLAN_BASE_PRICE {
		return custerr.ErrOutOfRange
	}

	if priceFrom > priceTo {
		return custerr.ErrOutOfRange
	}

	m.priceFrom = priceFrom
	m.priceTo = priceTo

	return nil

}

func (m *MenteeRecruitmentPlan) ValidateMenteeRecruitmentPlanTerm(termFrom, termTo time.Time) error {

	if termFrom.IsZero() {
		return custerr.ErrNotMatchValue
	}

	if termTo.IsZero() {
		return custerr.ErrNotMatchValue
	}

	diff_days := termTo.Sub(termFrom).Hours() / 24
	if diff_days > MENTEE_PLAN_MAX_DATE {
		return custerr.ErrOutOfRange
	}

	m.termFrom = termFrom
	m.termTo = termTo

	return nil

}
