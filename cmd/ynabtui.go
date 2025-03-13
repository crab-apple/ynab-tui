package main

import (
	"os"
	"ynabtui/driven_adapters/for_communicating_with_ynab/ynabclient"
	"ynabtui/internal/app"
	"ynabtui/internal/files"
	"ynabtui/internal/settings"
)

func main() {

	accessToken, err := settings.ReadAccessToken()
	if err != nil {
		panic(err)
	}

	api, err := ynabclient.NewClient("https://api.ynab.com/v1", accessToken)
	if err != nil {
		panic(err)
	}

	app.RunApp(os.Stdin, os.Stdout, api, files.AppFilesImpl{})
}
