package ynabclient

import (
	"context"
	"fmt"
	"github.com/oapi-codegen/oapi-codegen/v2/pkg/securityprovider"
	openapi_types "github.com/oapi-codegen/runtime/types"
	"github.com/rickb777/date/v2"
	"io"
	"net/http"
	"ynabtui/internal/ynabclientgen"
)

type YnabClient struct {
	clientGen *ynabclientgen.Client
}

func NewClient(apiUri string, token string) (YnabClient, error) {

	sp, err := securityprovider.NewSecurityProviderBearerToken(token)

	gc, err := ynabclientgen.NewClient(apiUri, ynabclientgen.WithRequestEditorFn(sp.Intercept))
	if err != nil {
		return YnabClient{}, err
	}

	return YnabClient{clientGen: gc}, nil
}

func (c YnabClient) ReadBudgets() (string, error) {
	return c.doGet(func(client *ynabclientgen.Client) (*http.Response, error) {
		return client.GetBudgets(context.TODO(), nil)
	})
}
func (c YnabClient) ReadTransactions(budgetId string, since date.Date) (string, error) {
	return c.doGet(func(client *ynabclientgen.Client) (*http.Response, error) {
		return client.GetTransactions(context.TODO(), budgetId, &ynabclientgen.GetTransactionsParams{
			SinceDate: &openapi_types.Date{Time: since.Midnight()},
		})
	})
}

func (c YnabClient) doGet(query func(client *ynabclientgen.Client) (*http.Response, error)) (string, error) {

	resp, err := query(c.clientGen)

	if err != nil {
		return "", err
	}
	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("expected HTTP 200 but received %d", resp.StatusCode)
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}
