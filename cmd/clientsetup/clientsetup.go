package clientsetup

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"ynabtui/app/driven_ports"
	"ynabtui/driven_adapters/for_communicating_with_ynab/ynabclient"
)

func SetupYnabClient() (driven_ports.ForCommunicatingWithYnab, error) {

	accessToken, err := readAccessToken()
	if err != nil {
		return nil, err
	}

	api, err := ynabclient.NewClient("https://api.ynab.com/v1", accessToken)
	if err != nil {
		return nil, err
	}

	return api, nil
}

func readAccessToken() (string, error) {
	homedir, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("unable to read access token: %w", err)
	}
	contents, err := os.ReadFile(filepath.Join(homedir, ".ynab", "access_token"))
	if err != nil {
		return "", fmt.Errorf("unable to read access token: %w", err)
	}
	return strings.TrimSpace(string(contents)), nil
}
