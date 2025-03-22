package test

import (
	"github.com/charmbracelet/x/exp/teatest"
	"testing"
	"time"
	"ynabtui/app/app/ui"
	"ynabtui/internal/ui/tea"
)

func TestQQuitsProgram(t *testing.T) {

	tm := teatest.NewTestModel(t, tea.InitialModel(ui.NewUI(NewFakeYnab().Api())), teatest.WithInitialTermSize(120, 30))

	tm.Type("q")

	tm.WaitFinished(t, teatest.WithFinalTimeout(500*time.Millisecond))
}
