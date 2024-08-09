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

	r, err := c.doRequest("https://api.ynab.com/v1/budgets", nil)
	if err != nil {
		return "", fmt.Errorf("error while reading budgets,  %w", err)
	}
	return r, nil
}
func (c YnabClient) ReadTransactions(budgetId string, since date.Date) (string, error) {

	r, err := c.doRequest(fmt.Sprintf("https://api.ynab.com/v1/budgets/%s/transactions", budgetId), map[string]string{
		"since_date": since.String(),
	})
	if err != nil {
		return "", fmt.Errorf("error while reading transactions,  %w", err)
	}
	return r, nil
}

func (c YnabClient) doRequest(path string, query map[string]string) (string, error) {

	req, err := http.NewRequest("GET", path, nil)
	if err != nil {
		return "", err
	}

	if query != nil {
		q := req.URL.Query()
		for k, v := range query {
			q.Add(k, v)
		}
		req.URL.RawQuery = q.Encode()
	}

	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.token))

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
