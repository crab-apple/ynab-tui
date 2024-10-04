package test

import (
	"github.com/stretchr/testify/require"
	"io"
	"sync"
	"testing"
	"time"
	"ynabtui/internal/app"
	"ynabtui/test/term"
)

type TestEnv struct {
	t     *testing.T
	Ynab  *FakeYnab
	files AppFilesFake
	tterm *term.TestTerminal
	wg    *sync.WaitGroup
}

func NewTestEnv(t *testing.T) TestEnv {
	return TestEnv{
		t:     t,
		Ynab:  NewFakeYnab(),
		files: AppFilesFake{},
		tterm: term.NewTestTerminal(),
		wg:    &sync.WaitGroup{},
	}
}

func (env TestEnv) Run() {

	// Run the program
	env.wg.Add(1)
	go func() {
		app.RunApp(env.tterm.InputReader, env.tterm.OutputWriter, env.Ynab.Api(), env.files)
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

func (env TestEnv) WaitUntil(condition func(output string) bool, timeout time.Duration) string {
	start := time.Now()
	for time.Since(start) < timeout {
		output, err := env.tterm.GetOutput()
		require.NoError(env.t, err)
		if condition(output) {
			return output
		}
		time.Sleep(1 * time.Millisecond)
	}
	env.t.Error("Did not reach expected state in time")
	return ""
}

func (env TestEnv) WaitFinish() {

	// Wait for the program to finish
	require.False(env.t, waitTimeout(env.wg, 100*time.Millisecond))

	// Check for errors
	select {
	case err := <-env.tterm.Errs:
		env.t.Error(err)
	default:
	}
}

func (env TestEnv) Type(r rune) {
	err := env.tterm.Type(r)
	require.NoError(env.t, err)
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