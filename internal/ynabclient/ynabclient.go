package ynabclient

import (
	"fmt"
	"github.com/rickb777/date/v2"
	"io"
	"net/http"
)

type YnabClient struct {
	token string
}

func NewClient(token string) YnabClient {
	return YnabClient{token: token}
}

func (c YnabClient) ReadBudgets() (string, error) {

	req, err := http.NewRequest("GET", "https://api.ynab.com/v1/budgets", nil)
	if err != nil {
		return "", fmt.Errorf("error while reading budgets,  %w", err)
	}
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.token))

	return c.doRequest(req)
}
func (c YnabClient) ReadTransactions(budgetId string, since date.Date) (string, error) {

	req, err := http.NewRequest("GET", fmt.Sprintf("https://api.ynab.com/v1/budgets/%s/transactions", budgetId), nil)
	if err != nil {
		return "", fmt.Errorf("error while reading transactions,  %w", err)
	}
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.token))
	q := req.URL.Query()
	q.Add("since_date", since.String())
	req.URL.RawQuery = q.Encode()

	return c.doRequest(req)
}

func (c YnabClient) doRequest(req *http.Request) (string, error) {

	client := http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	if resp.StatusCode != 200 {
		return "", fmt.Errorf("status code %d: %s", resp.StatusCode, body)
	}

	return string(body), nil
}
