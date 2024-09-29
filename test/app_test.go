package test

import (
	"github.com/stretchr/testify/require"
	"strings"
	"testing"
	"time"
	"ynabtui/internal/ynabmodel"
)

func TestQQuitsProgram(t *testing.T) {

	env := NewTestEnv(t)

	env.Run()

	env.Type('q')

	env.WaitFinish()
}

func TestDisplaysTransactions(t *testing.T) {

	env := NewTestEnv(t)

	env.Ynab.SetTransactions([]ynabmodel.Transaction{
		MakeTransaction(&AccChecking, &CatGroceries, "2020-01-01", 12340, "Last minute groceries"),
		MakeTransaction(&AccCash, &CatGroceries, "2020-01-02", 3500, "Chewing gum"),
		MakeTransaction(&AccChecking, &CatRent, "2020-01-02", 1000000, ""),
	})

	env.Run()

	output := env.WaitUntil(func(output string) bool {
		return strings.Contains(output, "gum")
	}, 1*time.Second)

	require.Contains(t, output, "Last minute groceries")
	require.Contains(t, output, "Chewing gum")

	env.Type('q')

	env.WaitFinish()
}
