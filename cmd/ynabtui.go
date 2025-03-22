package main

import (
	"ynabtui/cmd/clientsetup"
	"ynabtui/cmd/logging"
	"ynabtui/internal/app"
)

func main() {

	defer logging.SetUpLogging()()

	forCommunicatingWithYnab, err := clientsetup.SetupYnabClient()
	if err != nil {
		panic(err)
	}

	app.RunApp(forCommunicatingWithYnab)
}
