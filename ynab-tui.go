package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

func main() {

	token, err := readAccessToken()
	if err != nil {
		panic(err)
	}

	budgets, err := readBudgets(token)
	if err != nil {
		panic(err)
	}

	println(budgets)
}

func readAccessToken() (string, error) {
	homedir, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("unable to read access token: %w", err)
	}
	f := filepath.Join(homedir, ".ynab", "access_token")
	contents, err := os.ReadFile(f)
	if err != nil {
		return "", fmt.Errorf("unable to read access token: %w", err)
	}
	return strings.TrimSpace(string(contents)), nil
}

func readBudgets(token string) (string, error) {

	client := http.Client{}

	req, err := http.NewRequest("GET", "https://api.ynab.com/v1/budgets", nil)
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))

	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("error while reading budgets,  %w", err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("error while reading budgets,  %w", err)
	}
	if resp.StatusCode != 200 {
		return "", fmt.Errorf("error while reading budgets, status code %d: %s", resp.StatusCode, body)
	}

	return string(body), nil
}
