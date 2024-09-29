package main

import (
	"github.com/stretchr/testify/require"
	"io"
	"sync"
	"testing"
	"time"
	"ynabtui/internal/ynabmodel"
	"ynabtui/test"
	"ynabtui/test/term"
)

type TestEnv struct {
	t     *testing.T
	ynab  *test.FakeYnab
	files AppFilesFake
	tterm *term.TestTerminal
	wg    *sync.WaitGroup
}

func NewTestEnv(t *testing.T) TestEnv {
	return TestEnv{
		t:     t,
		ynab:  test.NewFakeYnab(),
		files: AppFilesFake{},
		tterm: term.NewTestTerminal(),
		wg:    &sync.WaitGroup{},
	}
}

func (env TestEnv) Run() {

	// Run the program
	env.wg.Add(1)
	go func() {
		runApp(env.tterm.InputReader, env.tterm.OutputWriter, env.ynab.Api(), env.files)
		env.tterm.CleanUp()
		env.wg.Done()
	}()

	// Read the program output
	env.wg.Add(1)
	go func() {
		env.tterm.ProcessOutput()
		env.wg.Done()
	}()
}

func (env TestEnv) GetOutput() string {

	// Wait for the program to finish
	require.False(env.t, waitTimeout(env.wg, 100*time.Millisecond))

	// Check for errors
	select {
	case err := <-env.tterm.Errs:
		env.t.Error(err)
	default:
	}

	visible, err := env.tterm.GetOutput()
	require.NoError(env.t, err)

	return visible
}

func (env TestEnv) Type(r rune) {
	err := env.tterm.Type(r)
	require.NoError(env.t, err)
}

func TestQQuitsProgram(t *testing.T) {

	output := io.Discard

	inputReader, inputWriter := io.Pipe()

	wg := sync.WaitGroup{}
	wg.Add(1)

	go func() {
		runApp(inputReader, output, test.NewFakeYnab().Api(), AppFilesFake{})
		wg.Done()
	}()

	_, err := inputWriter.Write([]byte("q"))
	require.NoError(t, err)

	require.False(t, waitTimeout(&wg, 100*time.Millisecond))
}

func TestDisplaysTransactions(t *testing.T) {

	env := NewTestEnv(t)

	env.ynab.SetTransactions([]ynabmodel.Transaction{
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
