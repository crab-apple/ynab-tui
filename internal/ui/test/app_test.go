package test

import (
	"github.com/charmbracelet/x/exp/teatest"
	"testing"
	"time"
	"ynabtui/app/driving_ports"
	"ynabtui/internal/ui/tea"
)

func TestQQuitsProgram(t *testing.T) {

	var forDisplayingTheScreen driving_ports.ForDisplayingTheScreen
	forDisplayingTheScreen = FakeScreenSupplier{}

	tm := teatest.NewTestModel(t, tea.InitialModel(forDisplayingTheScreen), teatest.WithInitialTermSize(120, 30))

	tm.Type("q")

	tm.WaitFinished(t, teatest.WithFinalTimeout(500*time.Millisecond))
}
