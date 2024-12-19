package test

import (
	"github.com/charmbracelet/x/exp/teatest"
	"testing"
	"time"
	"ynabtui/internal/ui/tea"
)

func TestQQuitsProgram(t *testing.T) {

	tm := teatest.NewTestModel(t, tea.InitialModel(NewFakeYnab().Api()), teatest.WithInitialTermSize(120, 30))

	tm.Type("q")

	tm.WaitFinished(t, teatest.WithFinalTimeout(500*time.Millisecond))
}
