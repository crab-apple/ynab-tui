package tea

import (
	"fmt"
	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/samber/lo"
	"log/slog"
	"ynabtui/internal/ui/tea/components/responsivetable"
	"ynabtui/internal/ynabapi"
	"ynabtui/internal/ynabmodel"
	"ynabtui/internal/ynabmodel/date"
)

type UI struct {
	api ynabapi.YnabApi
}

func NewUI(api ynabapi.YnabApi) UI {
	return UI{api: api}
}

type Model struct {
	uiModel UI

	transactions []ynabmodel.Transaction
	table        responsivetable.Model
}

type readTransactionsMsg struct {
	transactions []ynabmodel.Transaction
}

func InitialModel(api ynabapi.YnabApi) Model {

	uiModel := NewUI(api)

	t := responsivetable.New(
		table.WithFocused(true),
	)

	t.SetHeight(15)

	columns := []responsivetable.Column{
		{Title: "Date"},
		{Title: "Account"},
		{Title: "Category"},
		{Title: "Amount"},
		{Title: "Memo"},
	}
	t.SetColumns(columns)

	return Model{
		uiModel:      uiModel,
		transactions: nil,
		table:        t,
	}
}

func (m Model) readTransactions() tea.Msg {
	since, _ := date.Today().MinusDays(7)

	budgets, err := m.uiModel.api.ReadBudgets()
	if err != nil {
		// TODO handle
		panic(err)
	}

	budget := lo.MaxBy(budgets, func(a ynabmodel.Budget, b ynabmodel.Budget) bool {
		return a.LastModifiedOn.After(b.LastModifiedOn)
	})

	transactions, err := m.uiModel.api.ReadTransactions(budget.Id, since)
	if err != nil {
		// TODO handle
		panic(err)
	}

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
		m.table.SetRows(rows)

	case tea.WindowSizeMsg:
		m.table.SetWidth(msg.Width)
		m.table.SetHeight(msg.Height)

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
	return m.table.View()
}

func makeTransactionRow(t ynabmodel.Transaction) table.Row {
	return table.Row{t.Date.String(), t.AccountName, *t.CategoryName, t.Amount.Format(), t.Memo}
}
