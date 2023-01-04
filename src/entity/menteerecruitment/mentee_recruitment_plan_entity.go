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

	err := NewMenteePlanID(id)
	if err != nil {
		return nil, err
	}

	if title == "" {
		return nil, custerr.ErrEmptyValue
	}

	if l := utf8.RuneCountInString(title); l > MENTEE_PLAN_TITLE_MAX_LEN {
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

	if desc == "" {
		return nil, custerr.ErrEmptyValue
	}

	if l := utf8.RuneCountInString(desc); l > MENTEE_PLAN_DESCRIPTION_MAX_LEN {
		return nil, custerr.ErrValueIsTooLong
	}

	if priceFrom < MENTEE_PLAN_BASE_PRICE {
		return nil, custerr.ErrOutOfRange
	}

	if priceFrom > priceTo {
		return nil, custerr.ErrOutOfRange
	}

	if termFrom.IsZero() {
		return nil, custerr.ErrNotMatchValue
	}

	if termTo.IsZero() {
		return nil, custerr.ErrNotMatchValue
	}

	diff_days := termTo.Sub(termFrom).Hours() / 24
	if diff_days > MENTEE_PLAN_MAX_DATE {
		return nil, custerr.ErrOutOfRange
	}

	return &MenteeRecruitmentPlan{
		id:                 id,
		userID:             usrID,
		title:              title,
		categoryID:         ctgID,
		contractType:       contractType,
		consultType:        consultType,
		description:        desc,
		priceFrom:          priceFrom,
		priceTo:            priceTo,
		termFrom:           termFrom,
		termTo:             termTo,
		status:             stat,
		menteeContractList: contractList,
		menteeProposalList: proposalList,
	}, nil

}

func (m *MenteeRecruitmentPlan) ChangeMenteeRecruitmentPlanTitle(title string) error {

	if title == "" {
		return custerr.ErrEmptyValue
	}

	if l := utf8.RuneCountInString(title); l > MENTEE_PLAN_TITLE_MAX_LEN {
		return custerr.ErrValueIsTooLong
	}

	return nil

}

func (m *MenteeRecruitmentPlan) ChangeMenteeRecruitmentPlanDescription(desc string) error {

	if desc == "" {
		return custerr.ErrEmptyValue
	}

	if l := utf8.RuneCountInString(desc); l > MENTEE_PLAN_DESCRIPTION_MAX_LEN {
		return custerr.ErrValueIsTooLong
	}

	m.description = desc

	return nil

}
