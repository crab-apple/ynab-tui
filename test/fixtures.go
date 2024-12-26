package test

import (
	"fmt"
	"github.com/google/uuid"
	"math/rand"
	"ynabtui/internal/lang/optional"
	"ynabtui/internal/ynabmodel"
	"ynabtui/internal/ynabmodel/date"
)

var (
	AccCash = ynabmodel.Account{
		Id:   uuid.UUID{},
		Name: "Cash",
	}
	AccChecking = ynabmodel.Account{
		Id:   uuid.UUID{},
		Name: "Checking account",
	}
)

var (
	CatGroceries = ynabmodel.Category{
		Id:   uuid.UUID{},
		Name: "Groceries",
	}
	CatRent = ynabmodel.Category{
		Id:   uuid.UUID{},
		Name: "Rent",
	}
)

func MakeTransaction(account *ynabmodel.Account, category optional.Optional[ynabmodel.Category], dateStr string, amount string, memo string) ynabmodel.Transaction {

	d, err := date.Parse(dateStr)
	if err != nil {
		panic(err)
	}

	amountMoney, err := ynabmodel.NewMoneyFromString(amount)
	if err != nil {
		panic(err)
	}

	return ynabmodel.Transaction{
		Id:           fmt.Sprintf("%d", rand.Uint32()),
		Date:         d,
		AccountId:    account.Id,
		AccountName:  account.Name,
		CategoryId:   optional.Map(category, func(c ynabmodel.Category) uuid.UUID { return c.Id }),
		CategoryName: optional.Map(category, func(c ynabmodel.Category) string { return c.Name }),
		Amount:       amountMoney,
		Memo:         memo,
	}
}
