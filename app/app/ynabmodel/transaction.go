package ynabmodel

import (
	"github.com/google/uuid"
	"ynabtui/app/app/ynabmodel/date"
	"ynabtui/internal/lang/optional"
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
