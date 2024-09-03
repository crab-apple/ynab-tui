package ynab

import (
	"github.com/google/uuid"
	"time"
)

type Transaction struct {
	Id           string
	Date         time.Time
	AccountId    uuid.UUID
	AccountName  string
	CategoryId   *uuid.UUID
	CategoryName *string
	Amount       Money
	Memo         string
}
