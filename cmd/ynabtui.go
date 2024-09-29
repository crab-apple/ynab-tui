package main

import (
	"os"
	"ynabtui/internal/app"
	"ynabtui/internal/files"
	"ynabtui/internal/settings"
	"ynabtui/internal/ynabapi"
)

func main() {

	accessToken, err := settings.ReadAccessToken()
	if err != nil {
		panic(err)
	}

	api, err := ynabapi.NewClient("https://api.ynab.com/v1", accessToken)
	if err != nil {
		panic(err)
	}

	app.RunApp(os.Stdin, os.Stdout, api, files.AppFilesImpl{})
}
