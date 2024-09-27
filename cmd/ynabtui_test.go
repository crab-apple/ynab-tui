package main

import (
	"github.com/stretchr/testify/require"
	"io"
	"sync"
	"testing"
	"time"
	"ynabtui/internal/ynabmodel"
	"ynabtui/test"
)

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

	ynab := test.NewFakeYnab()

	ynab.SetTransactions([]ynabmodel.Transaction{
		test.MakeTransaction(&test.AccChecking, &test.CatGroceries, "2020-01-01", 12340, "Last minute groceries"),
		test.MakeTransaction(&test.AccCash, &test.CatGroceries, "2020-01-02", 3500, "Chewing gum"),
		test.MakeTransaction(&test.AccChecking, &test.CatRent, "2020-01-02", 1000000, ""),
	})

	outputReader, outputWriter := io.Pipe()
	inputReader, inputWriter := io.Pipe()
	errs := make(chan error, 8)

	wg := sync.WaitGroup{}

	// Run the program
	wg.Add(1)
	go func() {
		runApp(inputReader, outputWriter, ynab.Api(), AppFilesFake{})
		if err := outputWriter.Close(); err != nil {
			errs <- err
		}
		wg.Done()
	}()

	// Read the program output
	output := make([]byte, 0)
	wg.Add(1)
	go func() {
		for {
			b := make([]byte, 8)
			n, err := outputReader.Read(b)
			output = append(output, b[:n]...)
			if err == io.EOF {
				break
			}
			if err != nil {
				errs <- err
				break
			}
		}
		wg.Done()
	}()

	var err error

	_, err = inputWriter.Write([]byte("q"))
	require.NoError(t, err)

	// Wait for the program to finish
	require.False(t, waitTimeout(&wg, 100*time.Millisecond))

	// Check for errors
	select {
	case err = <-errs:
		t.Error(err)
	default:
	}

	// Assert output
	visible, err := test.ParseTerminalOutput(output)
	require.NoError(t, err)

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
