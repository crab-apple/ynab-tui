package ynabmodel

import (
	"github.com/google/uuid"
	"ynabtui/internal/lang/optional"
	"ynabtui/internal/ynabmodel/date"
)

type Transaction struct {
	Id           string
	Date         date.Date
	AccountId    uuid.UUID
	AccountName  string
	CategoryId   optional.Optional[uuid.UUID]
	CategoryName optional.Optional[string]
	Amount       Money
	Memo         string
}
