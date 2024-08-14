package main

import (
	"github.com/stretchr/testify/require"
	"io"
	"sync"
	"testing"
	"time"
)

func TestQQuitsProgram(t *testing.T) {

	output := io.Discard

	inputReader, inputWriter := io.Pipe()

	wg := sync.WaitGroup{}
	wg.Add(1)

	go func() {
		runApp(inputReader, output, AppFilesFake{})
		wg.Done()
	}()

	_, err := inputWriter.Write([]byte("q"))
	require.NoError(t, err)

	require.False(t, waitTimeout(&wg, 100*time.Millisecond))
}

func TestDisplaysGroceries(t *testing.T) {

	outputReader, outputWriter := io.Pipe()
	inputReader, inputWriter := io.Pipe()
	errs := make(chan error, 8)

	wg := sync.WaitGroup{}

	// Run the program
	wg.Add(1)
	go func() {
		runApp(inputReader, outputWriter, AppFilesFake{})
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
	require.Contains(t, string(output), "Buy carrots")
	require.Contains(t, string(output), "Buy celery")
	require.Contains(t, string(output), "Buy kohlrabi")
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
