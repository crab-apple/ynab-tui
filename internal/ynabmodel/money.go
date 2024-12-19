package ynabmodel

import (
	"errors"
	"golang.org/x/text/message"
	"regexp"
	"strconv"
	"strings"
)

type Money struct {
	cents int64
}

func NewMoneyFromThousandths(thousandths int64) (Money, error) {
	if thousandths%10 != 0 {
		return Money{}, errors.New("only two decimal places supported")
	}
	return Money{cents: thousandths / 10}, nil
}

func NewMoneyFromString(str string) (Money, error) {

	matches, err := regexp.MatchString("^\\d+\\.\\d\\d$", str)
	if err != nil {
		return Money{}, err
	}

	if !matches {
		return Money{}, errors.New("Money strings must contain at least one digit to the left of the point and exactly two digits to the right")
	}

	centsStr := strings.ReplaceAll(str, ".", "")
	val, err := strconv.Atoi(centsStr)
	if err != nil {
		return Money{}, err
	}

	return Money{cents: int64(val)}, nil
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
