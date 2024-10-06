package test

import (
	"bytes"
	"github.com/charmbracelet/x/exp/teatest"
	"testing"
	"time"
	"ynabtui/internal/ynabmodel"
)

func TestQQuitsProgram(t *testing.T) {

	env := NewTestEnv(t)

	tm := env.Start()

	tm.Type("q")

	tm.WaitFinished(t, teatest.WithFinalTimeout(500*time.Millisecond))
}

func TestDisplaysTransactions(t *testing.T) {

	env := NewTestEnv(t)

	env.Ynab.SetTransactions([]ynabmodel.Transaction{
		MakeTransaction(&AccChecking, &CatGroceries, "2020-01-01", 12340, "Last minute groceries"),
		MakeTransaction(&AccCash, &CatGroceries, "2020-01-02", 3500, "Chewing gum"),
		MakeTransaction(&AccChecking, &CatRent, "2020-01-02", 1000000, ""),
	})

	env.Start()

	env.AssertOutput(func(bts []byte) bool {
		return bytes.Contains(bts, []byte("Last minute groceries")) && bytes.Contains(bts, []byte("Chewing gum"))
	})

	env.Quit()
}
