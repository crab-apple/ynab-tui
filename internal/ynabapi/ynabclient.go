package ynabapi

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/oapi-codegen/oapi-codegen/v2/pkg/securityprovider"
	openapitypes "github.com/oapi-codegen/runtime/types"
	"github.com/samber/lo"
	"net/http"
	"ynabtui/internal/ynabclientgen"
	"ynabtui/internal/ynabmodel"
	"ynabtui/internal/ynabmodel/date"
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

func (c YnabClient) ReadBudgets() ([]ynabmodel.Budget, error) {

	res, err := c.clientGen.GetBudgetsWithResponse(context.TODO(), nil)

	if err != nil {
		return nil, err
	}
	if res.StatusCode() != http.StatusOK {
		return nil, fmt.Errorf("expected HTTP 200 but received %d", res.StatusCode())
	}

	return lo.Map(res.JSON200.Data.Budgets, func(item ynabclientgen.BudgetSummary, index int) ynabmodel.Budget {
		return ynabmodel.Budget{
			Id:             item.Id,
			LastModifiedOn: *item.LastModifiedOn,
		}
	}), nil
}

func (c YnabClient) ReadTransactions(budgetId uuid.UUID, since date.Date) ([]ynabmodel.Transaction, error) {

	res, err := c.clientGen.GetTransactionsWithResponse(context.TODO(), budgetId.String(), &ynabclientgen.GetTransactionsParams{
		SinceDate: &openapitypes.Date{Time: since.Midnight()},
	})

	if err != nil {
		return nil, err
	}
	if res.StatusCode() != http.StatusOK {
		return nil, fmt.Errorf("expected HTTP 200 but received %d, %s", res.StatusCode(), res.Body)
	}

	result := make([]ynabmodel.Transaction, 0)

	for _, t := range res.JSON200.Data.Transactions {
		mapped, err := mapTransaction(t)
		if err != nil {
			return nil, err
		}
		result = append(result, mapped)
	}
	return result, nil
}

func mapTransaction(t ynabclientgen.TransactionDetail) (ynabmodel.Transaction, error) {

	d, err := date.FromTime(t.Date.Time)
	if err != nil {
		return ynabmodel.Transaction{}, err
	}

	amount, err := ynabmodel.NewMoney(t.Amount)
	if err != nil {
		return ynabmodel.Transaction{}, err
	}

	memo := ""
	if t.Memo != nil {
		memo = *t.Memo
	}

	return ynabmodel.Transaction{
		Id:           t.Id,
		Date:         d,
		AccountId:    t.AccountId,
		AccountName:  t.AccountName,
		CategoryId:   t.CategoryId,
		CategoryName: t.CategoryName,
		Amount:       amount,
		Memo:         memo,
	}, nil
}
