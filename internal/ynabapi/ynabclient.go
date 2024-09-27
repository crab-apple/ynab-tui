package ynabapi

import (
	"context"
	"fmt"
	"github.com/oapi-codegen/oapi-codegen/v2/pkg/securityprovider"
	openapitypes "github.com/oapi-codegen/runtime/types"
	"github.com/rickb777/date/v2"
	"net/http"
	"ynabtui/internal/ynabclientgen"
)

type YnabClient struct {
	clientGen *ynabclientgen.ClientWithResponses
}

func NewClient(apiUri string, token string) (YnabClient, error) {

	sp, err := securityprovider.NewSecurityProviderBearerToken(token)
	if err != nil {
		return YnabClient{}, err
	}

	gcr, err := ynabclientgen.NewClientWithResponses(apiUri, ynabclientgen.WithRequestEditorFn(sp.Intercept))
	if err != nil {
		return YnabClient{}, err
	}

	return YnabClient{clientGen: gcr}, nil
}

func (c YnabClient) ReadBudgets() ([]ynabclientgen.BudgetSummary, error) {

	res, err := c.clientGen.GetBudgetsWithResponse(context.TODO(), nil)

	if err != nil {
		return nil, err
	}
	if res.StatusCode() != http.StatusOK {
		return nil, fmt.Errorf("expected HTTP 200 but received %d", res.StatusCode())
	}

	return res.JSON200.Data.Budgets, nil
}

func (c YnabClient) ReadTransactions(budgetId string, since date.Date) ([]ynabclientgen.TransactionDetail, error) {

	res, err := c.clientGen.GetTransactionsWithResponse(context.TODO(), budgetId, &ynabclientgen.GetTransactionsParams{
		SinceDate: &openapitypes.Date{Time: since.Midnight()},
	})

	if err != nil {
		return nil, err
	}
	if res.StatusCode() != http.StatusOK {
		return nil, fmt.Errorf("expected HTTP 200 but received %d", res.StatusCode())
	}

	return res.JSON200.Data.Transactions, nil
}
