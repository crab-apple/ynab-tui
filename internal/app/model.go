package app

import (
	"fmt"
	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/samber/lo"
	"log/slog"
	"ynabtui/internal/ynabapi"
	"ynabtui/internal/ynabmodel"
	"ynabtui/internal/ynabmodel/date"
)

type Model struct {
	// TODO I'm not sure that it makes sense to have dependencies in the model. Should revisit this later.
	api ynabapi.YnabApi

	transactions      []ynabmodel.Transaction
	transactionsTable table.Model
}

type readTransactionsMsg struct {
	transactions []ynabmodel.Transaction
}

func InitialModel(api ynabapi.YnabApi) Model {

	t := table.New(
		table.WithFocused(true),
	)

	t.SetHeight(15)

	columns := []table.Column{
		{Title: "Date", Width: 15},
		{Title: "Account", Width: 15},
		{Title: "Category", Width: 30},
		{Title: "Amount", Width: 15},
		{Title: "Memo", Width: 30},
	}
	t.SetColumns(columns)

	return Model{
		api: api,

		transactions:      nil,
		transactionsTable: t,
	}
}

func (m Model) readTransactions() tea.Msg {
	since, _ := date.Today().MinusDays(7)

	budgets, err := m.api.ReadBudgets()
	if err != nil {
		// TODO handle
		panic(err)
	}

	budget := lo.MaxBy(budgets, func(a ynabmodel.Budget, b ynabmodel.Budget) bool {
		return a.LastModifiedOn.After(b.LastModifiedOn)
	})

	transactions, err := m.api.ReadTransactions(budget.Id, since)
	if err != nil {
		// TODO handle
		panic(err)
	}

	transactions = lo.Slice(transactions, 0, 10)

	return readTransactionsMsg{
		transactions: transactions,
	}
}

func (m Model) Init() tea.Cmd {
	return m.readTransactions
}
func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

	slog.Debug("Received message", "type", fmt.Sprintf("%T", msg), "value", msg)

	switch msg := msg.(type) {

	case readTransactionsMsg:
		m.transactions = msg.transactions
		rows := lo.Map(m.transactions, func(item ynabmodel.Transaction, i int) table.Row {
			return makeTransactionRow(item)
		})
		m.transactionsTable.SetRows(rows)

	// Is it a key press?
	case tea.KeyMsg:

		// Cool, what was the actual key pressed?
		switch msg.String() {

		// These keys should exit the program.
		case "ctrl+c", "q":
			return m, tea.Quit

		}
	}

	// Return the updated Model to the Bubble Tea runtime for processing.
	// Note that we're not returning a command.
	return m, nil
}
func (m Model) View() string {
	// The header
	s := "Recent transactions:\n\n"

	s += m.transactionsTable.View() + "\n"

	// The footer
	s += "\nPress q to quit.\n"

	// Send the UI for rendering
	return s
}

func makeTransactionRow(t ynabmodel.Transaction) table.Row {
	return table.Row{t.Date.String(), t.AccountName, *t.CategoryName, t.Amount.Format(), t.Memo}
}
