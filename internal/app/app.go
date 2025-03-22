package app

import (
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
	"io"
	"os"
	"ynabtui/app/app/ui"
	"ynabtui/app/driven_ports"
	tea2 "ynabtui/internal/ui/tea"
)

func RunApp(input io.Reader, output io.Writer, forCommunicatingWithYnab driven_ports.ForCommunicatingWithYnab) {

	p := tea.NewProgram(tea2.InitialModel(ui.NewUI(forCommunicatingWithYnab)), tea.WithInput(input), tea.WithOutput(output), tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}
