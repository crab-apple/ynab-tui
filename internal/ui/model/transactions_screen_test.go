package model

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"ynabtui/internal/lang/optional"
	"ynabtui/internal/ynabmodel"
	"ynabtui/test"
)

var transactions = []ynabmodel.Transaction{
	test.MakeTransaction(test.AccCash, optional.Of(test.CatGroceries), "2020-10-01", "35.00", ""),
	test.MakeTransaction(test.AccChecking, optional.Of(test.CatRent), "2020-10-02", "1000.00", "The rent"),
}

func TestShouldHaveATableWithTheTransactions(t *testing.T) {

	// Given
	screen := NewTransactionsScreen(transactions)

	// When
	table := screen.Table()

	// Then
	assert.Len(t, table.Rows, 2)
}

func TestShouldContainColumnInformation(t *testing.T) {
	// Given
	screen := NewTransactionsScreen(make([]ynabmodel.Transaction, 0))

	// When
	table := screen.Table()

	// Then
	assert.Equal(t, []Column{
		{key: "date", display: "Date"},
		{key: "account", display: "Account"},
		{key: "category", display: "Category"},
		{key: "amount", display: "Amount"},
		{key: "memo", display: "Memo"},
	}, table.Columns)
}

func TestShouldDisplayAccount(t *testing.T) {
	// Given
	screen := NewTransactionsScreen(transactions)

	// When
	table := screen.Table()

	// Then
	assert.Equal(t, test.AccCash.Name, table.Rows[0]["account"])
}

func TestShouldDisplayCategory(t *testing.T) {
	// Given
	screen := NewTransactionsScreen(transactions)

	// When
	table := screen.Table()

	// Then
	assert.Equal(t, test.CatGroceries.Name, table.Rows[0]["category"])
}

func TestShouldDisplayBlankCategoryWhenNoCategory(t *testing.T) {
	screen := NewTransactionsScreen([]ynabmodel.Transaction{
		test.MakeTransaction(test.AccCash, optional.Empty[ynabmodel.Category](), "2020-10-01", "35.00", ""),
	})

	// When
	table := screen.Table()

	// Then
	assert.Equal(t, "", table.Rows[0]["category"])
}

func TestShouldDisplayDate(t *testing.T) {
	// Given
	screen := NewTransactionsScreen(transactions)

	// When
	table := screen.Table()

	// Then
	assert.Equal(t, "2020-10-01", table.Rows[0]["date"])
}

func TestShouldDisplayAmount(t *testing.T) {
	// Given
	screen := NewTransactionsScreen(transactions)

	// When
	table := screen.Table()

	// Then
	assert.Equal(t, "35.00", table.Rows[0]["amount"])
}

func TestShouldDisplayEmptyMemo(t *testing.T) {
	// Given
	screen := NewTransactionsScreen(transactions)

	// When
	table := screen.Table()

	// Then
	assert.Equal(t, "", table.Rows[0]["memo"])
}

func TestShouldDisplayNonEmptyMemo(t *testing.T) {
	// Given
	screen := NewTransactionsScreen(transactions)

	// When
	table := screen.Table()

	// Then
	assert.Equal(t, "The rent", table.Rows[1]["memo"])
}
