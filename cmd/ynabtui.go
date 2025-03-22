package main

import (
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
	"os"
	"ynabtui/app/app"
	"ynabtui/cmd/clientsetup"
	"ynabtui/cmd/logging"
	tea2 "ynabtui/internal/ui/tea"
)

func main() {

	defer logging.SetUpLogging()()

	forCommunicatingWithYnab, err := clientsetup.SetupYnabClient()
	if err != nil {
		panic(err)
	}

	var application = app.NewApp(forCommunicatingWithYnab)

	p := tea.NewProgram(tea2.InitialModel(application.ForDisplayingTheScreen()), tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}
