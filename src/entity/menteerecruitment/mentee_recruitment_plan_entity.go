package menteerecruitment

import (
	"time"
	"unicode/utf8"

	"github.com/ha-zu/learn-clean-dev/src/entity"
	ctg "github.com/ha-zu/learn-clean-dev/src/entity/category"
	custerr "github.com/ha-zu/learn-clean-dev/src/entity/customerror"
	"github.com/ha-zu/learn-clean-dev/src/entity/user"
)

type MenteeRecruitmentPlan struct {
	id             MenteeRecruitmentPlanID
	userID         user.UserID
	title          string
	categoryID     ctg.CategoryID
	contractType   string
	consultType    string
	description    string
	price_from     int
	price_to       int
	term_from      time.Time
	term_to        time.Time
	status         string
	menteeContract []MenteeRecruitentPlanContract
	menteeProposal []MenteeRecruitentPlanProposal
}

const (
	MENTEE_PLAN_DESCRIPTION_MAX_LEN = 2000
	MENTEE_PLAN_TITLE_MAX_LEN       = 255
	MENTEE_PLAN_MAX_DATE            = 14
)

func NewMenteeRecruitmentPlan(id, title, contractType, consultType, desc, stat string,
	price_from, price_to int, term_from, term_to time.Time,
	usrID user.UserID, ctgID ctg.CategoryID, contract []MenteeRecruitentPlanContract, proposal []MenteeRecruitentPlanProposal) (*MenteeRecruitmentPlan, error) {

	mrpID, err := MenteePlanIDValidate(id)
	if err != nil {
		return nil, err
	}

	if title == "" {
		return nil, custerr.ErrEmptyValue
	}

	if l := utf8.RuneCountInString(title); l > MENTEE_PLAN_TITLE_MAX_LEN {
		return nil, custerr.ErrValueIsTooLong
	}

	if ctgID == "" {
		return nil, custerr.ErrEmptyValue
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

	if consultType != entity.PLAN_CONSULT_TEL && consultType != entity.PLAN_CONSULT_CHAT {
		return nil, custerr.ErrNotMatchValue
	}

	if desc == "" {
		return nil, custerr.ErrEmptyValue
	}

	if l := utf8.RuneCountInString(desc); l > MENTEE_PLAN_DESCRIPTION_MAX_LEN {
		return nil, custerr.ErrValueIsTooLong
	}

	if price_from < entity.PLAN_PRICE_BASE {
		return nil, custerr.ErrOutOfRange
	}

	if price_from > price_to {
		return nil, custerr.ErrOutOfRange
	}

	if term_from.IsZero() {
		return nil, custerr.ErrNotMatchValue
	}

	if term_to.IsZero() {
		return nil, custerr.ErrNotMatchValue
	}

	diff_days := term_to.Sub(term_from).Hours() / 24
	if diff_days > MENTEE_PLAN_MAX_DATE {
		return nil, custerr.ErrOutOfRange
	}

	if stat == "" {
		return nil, custerr.ErrEmptyValue
	}

	return &MenteeRecruitmentPlan{
		id:             *mrpID,
		userID:         usrID,
		title:          title,
		categoryID:     ctgID,
		contractType:   contractType,
		consultType:    consultType,
		description:    desc,
		price_from:     price_from,
		price_to:       price_to,
		term_from:      term_from,
		term_to:        term_to,
		status:         stat,
		menteeContract: contract,
		menteeProposal: proposal,
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
