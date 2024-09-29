package main

import (
	"github.com/stretchr/testify/require"
	"io"
	"sync"
	"testing"
	"time"
	"ynabtui/internal/app"
	"ynabtui/internal/ynabmodel"
	"ynabtui/test"
)

func TestQQuitsProgram(t *testing.T) {

	output := io.Discard

	inputReader, inputWriter := io.Pipe()

	wg := sync.WaitGroup{}
	wg.Add(1)

	go func() {
		app.RunApp(inputReader, output, test.NewFakeYnab().Api(), AppFilesFake{})
		wg.Done()
	}()

	_, err := inputWriter.Write([]byte("q"))
	require.NoError(t, err)

	require.False(t, waitTimeout(&wg, 100*time.Millisecond))
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

type AppFilesFake struct {
}

func (AppFilesFake) GetLogWriter() (io.Writer, func(), error) {
	return io.Discard, func() {}, nil
}

// waitTimeout waits for the waitgroup for the specified max timeout.
// Returns true if waiting timed out.
func waitTimeout(wg *sync.WaitGroup, timeout time.Duration) bool {
	c := make(chan struct{})
	go func() {
		defer close(c)
		wg.Wait()
	}()
	select {
	case <-c:
		return false // completed normally
	case <-time.After(timeout):
		return true // timed out
	}
}
