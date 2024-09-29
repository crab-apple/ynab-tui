package main

import (
	"github.com/stretchr/testify/require"
	"testing"
	"ynabtui/internal/ynabmodel"
	"ynabtui/test"
)

func TestQQuitsProgram(t *testing.T) {

	env := test.NewTestEnv(t)

	env.Run()

	env.Type('q')

	env.GetOutput()
}

func TestDisplaysTransactions(t *testing.T) {

	env := test.NewTestEnv(t)

	env.Ynab.SetTransactions([]ynabmodel.Transaction{
		test.MakeTransaction(&test.AccChecking, &test.CatGroceries, "2020-01-01", 12340, "Last minute groceries"),
		test.MakeTransaction(&test.AccCash, &test.CatGroceries, "2020-01-02", 3500, "Chewing gum"),
		test.MakeTransaction(&test.AccChecking, &test.CatRent, "2020-01-02", 1000000, ""),
	})

	env.Run()

	env.Type('q')

	visible := env.GetOutput()

	require.Contains(t, visible, "Last minute groceries")
	require.Contains(t, visible, "Chewing gum")
}
