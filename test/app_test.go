package test

import (
	"github.com/stretchr/testify/require"
	"testing"
	"ynabtui/internal/ynabmodel"
)

func TestQQuitsProgram(t *testing.T) {

	env := NewTestEnv(t)

	env.Run()

	env.Type('q')

	env.GetOutput()
}

func TestDisplaysTransactions(t *testing.T) {

	env := NewTestEnv(t)

	env.Ynab.SetTransactions([]ynabmodel.Transaction{
		MakeTransaction(&AccChecking, &CatGroceries, "2020-01-01", 12340, "Last minute groceries"),
		MakeTransaction(&AccCash, &CatGroceries, "2020-01-02", 3500, "Chewing gum"),
		MakeTransaction(&AccChecking, &CatRent, "2020-01-02", 1000000, ""),
	})

	env.Run()

	env.Type('q')

	visible := env.GetOutput()

	require.Contains(t, visible, "Last minute groceries")
	require.Contains(t, visible, "Chewing gum")
}
