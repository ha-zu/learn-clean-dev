package entity

import (
	"time"
	"unicode/utf8"

	"github.com/ha-zu/learn-clean-dev/src/entity/vo"
	ctg "github.com/ha-zu/learn-clean-dev/src/entity/vo/category"
	mr "github.com/ha-zu/learn-clean-dev/src/entity/vo/mentee_recruitment"
	"github.com/ha-zu/learn-clean-dev/src/entity/vo/user"
)

type MenteeRecruitmentPlan struct {
	id           mr.MenteeRecruitmentPlanID
	userID       user.UserID
	title        string
	categoryID   ctg.CategoryID
	contractType string
	consultType  string
	description  string
	price_from   int
	price_to     int
	term_from    time.Time
	term_to      time.Time
	status       string
}

const (
	MENTEE_DESCRIPTION_MAX_LEN = 2000
	MENTEE_TITLE_MAX_LEN       = 255
	MENTEE_PLAN_MAX_DATE       = 14
)

func NewMenteeRecruitmentPlan(id, title, contractType, consultType, desc, stat string, price_from, price_to int,
	term_from, term_to time.Time, usrID user.UserID, ctgID ctg.CategoryID) (*MenteeRecruitmentPlan, error) {

	mrpID, err := mr.MenteePlanIDValidate(id)
	if err != nil {
		return nil, err
	}

	if title == "" {
		return nil, vo.ErrEmptyValue
	}

	if l := utf8.RuneCountInString(title); l > MENTEE_TITLE_MAX_LEN {
		return nil, vo.ErrValueIsTooLong
	}

	if ctgID == "" {
		return nil, vo.ErrEmptyValue
	}

	if contractType == "" {
		return nil, vo.ErrEmptyValue
	}

	if contractType != vo.PLAN_CONTRACT_ONECE && contractType != vo.PLAN_CONTRACT_CONTINUTION {
		return nil, vo.ErrNotMatchValue
	}

	if consultType == "" {
		return nil, vo.ErrEmptyValue
	}

	if consultType != vo.PLAN_CONSULT_TEL && consultType != vo.PLAN_CONSULT_CHAT {
		return nil, vo.ErrNotMatchValue
	}

	if desc == "" {
		return nil, vo.ErrEmptyValue
	}

	if l := utf8.RuneCountInString(desc); l > MENTEE_DESCRIPTION_MAX_LEN {
		return nil, vo.ErrValueIsTooLong
	}

	if price_from < vo.PLAN_PRICE_BASE {
		return nil, vo.ErrOutOfRange
	}

	if price_from > price_to {
		return nil, vo.ErrOutOfRange
	}

	if term_from.IsZero() {
		return nil, vo.ErrNotMatchValue
	}

	if term_to.IsZero() {
		return nil, vo.ErrNotMatchValue
	}

	if stat == "" {
		return nil, vo.ErrEmptyValue
	}

	return &MenteeRecruitmentPlan{
		id:           *mrpID,
		userID:       usrID,
		title:        title,
		categoryID:   ctgID,
		contractType: contractType,
		consultType:  consultType,
		description:  desc,
		price_from:   price_from,
		price_to:     price_to,
		term_from:    term_from,
		term_to:      term_to,
		status:       stat,
	}, nil
}

func (m *MenteeRecruitmentPlan) ChangeMenteeRecruitmentPlanDescription(desc string) error {

	m.description = desc

	return nil
}
