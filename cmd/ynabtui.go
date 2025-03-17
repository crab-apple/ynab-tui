package main

import (
	"fmt"
	"os"
	"strings"
	"ynabtui/cmd/files"
	"ynabtui/cmd/logging"
	"ynabtui/driven_adapters/for_communicating_with_ynab/ynabclient"
	"ynabtui/internal/app"
)

func main() {

	defer logging.SetUpLogging()()

	accessToken, err := readAccessToken()
	if err != nil {
		panic(err)
	}

	api, err := ynabclient.NewClient("https://api.ynab.com/v1", accessToken)
	if err != nil {
		panic(err)
	}

	app.RunApp(os.Stdin, os.Stdout, api)
}

func readAccessToken() (string, error) {
	c, err := files.ReadYnabConfigFile("access_token")
	if err != nil {
		return "", fmt.Errorf("unable to read access token: %w", err)
	}
	return strings.TrimSpace(c), nil
}
