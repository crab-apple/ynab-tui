package ynabmodel

import (
	"errors"
	"golang.org/x/text/message"
)

type Money struct {
	cents int64
}

func NewMoney(thousandths int64) (Money, error) {
	if thousandths%10 != 0 {
		return Money{}, errors.New("only two decimal places supported")
	}
	return Money{cents: thousandths / 10}, nil
}

func (m Money) Format() string {
	return formatCents(m.cents)
}

func formatCents(c int64) string {
	if c < 0 {
		return "-" + formatCents(-c)
	}
	p := message.NewPrinter(message.MatchLanguage("en"))
	return p.Sprint(c/100) + "." + p.Sprintf("%02d", c%100)
}
