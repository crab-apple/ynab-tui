package ynabmodel

import (
	"github.com/google/uuid"
	"ynabtui/internal/ynabmodel/date"
)

type Transaction struct {
	Id           string
	Date         date.Date
	AccountId    uuid.UUID
	AccountName  string
	CategoryId   *uuid.UUID
	CategoryName *string
	Amount       Money
	Memo         string
}
