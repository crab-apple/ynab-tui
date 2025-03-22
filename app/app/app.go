package app

import (
	"ynabtui/app/app/ui"
	"ynabtui/app/driven_ports"
	"ynabtui/app/driving_ports"
)

type App struct {
	forCommunicatingWithYnab driven_ports.ForCommunicatingWithYnab
}

func NewApp(forCommunicatingWithYnab driven_ports.ForCommunicatingWithYnab) App {
	return App{forCommunicatingWithYnab}
}

func (a App) ForDisplayingTheScreen() driving_ports.ForDisplayingTheScreen {
	return ui.NewUI(a.forCommunicatingWithYnab)
}
