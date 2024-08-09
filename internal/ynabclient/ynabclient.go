package ynabclient

import (
	"fmt"
	"io"
	"net/http"
)

func ReadBudgets(token string) (string, error) {

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
func ReadTransactions(token string, budgetId string) (string, error) {

	client := http.Client{}

	req, err := http.NewRequest("GET", fmt.Sprintf("https://api.ynab.com/v1/budgets/%s/transactions", budgetId), nil)
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
