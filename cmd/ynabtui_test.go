package main

import (
	"bytes"
	"github.com/stretchr/testify/require"
	"sync"
	"testing"
	"time"
)

func TestQQuitsProgram(t *testing.T) {

	output := bytes.Buffer{}
	input := bytes.Buffer{}

	wg := sync.WaitGroup{}
	wg.Add(1)

	go func() {
		runApp(&input, &output)
		wg.Done()
	}()

	input.WriteRune('q')

	require.False(t, waitTimeout(&wg, 100*time.Millisecond))
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
