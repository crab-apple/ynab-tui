package ui

import (
	"github.com/samber/lo"
	"ynabtui/app/app/ynabmodel"
)

type TransactionsScreen struct {
	Transactions []ynabmodel.Transaction
}

func (ts TransactionsScreen) Table() Table {
	return Table{
		Columns: []Column{
			{Key: "date", Display: "Date"},
			{Key: "account", Display: "Account"},
			{Key: "category", Display: "Category"},
			{Key: "amount", Display: "Amount", CellAlign: AlignRight},
			{Key: "memo", Display: "Memo"},
		},
		Rows: lo.Map(ts.Transactions, func(t ynabmodel.Transaction, i int) Row {
			row := make(Row)
			row["account"] = t.AccountName
			row["category"] = t.CategoryName.Or("")
			row["date"] = t.Date.String()
			row["amount"] = t.Amount.Format()
			row["memo"] = t.Memo
			return row
		}),
	}
}

func NewTransactionsScreen(transactions []ynabmodel.Transaction) *TransactionsScreen {
	return &TransactionsScreen{Transactions: transactions}
}
