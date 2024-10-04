package app

import (
	"fmt"
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

	transactions []ynabmodel.Transaction
	cursor       int
	selected     map[int]struct{}
}

type readTransactionsMsg struct {
	transactions []ynabmodel.Transaction
}

func InitialModel(api ynabapi.YnabApi) Model {

	return Model{
		api: api,

		transactions: nil,

		// A map which indicates which choices are selected. We're using
		// the  map like a mathematical set. The keys refer to the indexes
		// of the `choices` slice, above.
		selected: make(map[int]struct{}),
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

	// Is it a key press?
	case tea.KeyMsg:

		// Cool, what was the actual key pressed?
		switch msg.String() {

		// These keys should exit the program.
		case "ctrl+c", "q":
			return m, tea.Quit

		// The "up" and "k" keys move the cursor up
		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}

		// The "down" and "j" keys move the cursor down
		case "down", "j":
			if m.cursor < len(m.transactions)-1 {
				m.cursor++
			}

		// The "enter" key and the spacebar (a literal space) toggle
		// the selected state for the item that the cursor is pointing at.
		case "enter", " ":
			_, ok := m.selected[m.cursor]
			if ok {
				delete(m.selected, m.cursor)
			} else {
				m.selected[m.cursor] = struct{}{}
			}
		}
	}

	// Return the updated Model to the Bubble Tea runtime for processing.
	// Note that we're not returning a command.
	return m, nil
}
func (m Model) View() string {
	// The header
	s := "Recent transactions:\n\n"

	// Iterate over our choices
	for i, transaction := range m.transactions {

		// Is the cursor pointing at this choice?
		cursor := " " // no cursor
		if m.cursor == i {
			cursor = ">" // cursor!
		}

		// Is this choice selected?
		checked := " " // not selected
		if _, ok := m.selected[i]; ok {
			checked = "x" // selected!
		}

		// Render the row
		s += fmt.Sprintf("%s [%s] %s\n", cursor, checked, renderTransaction(transaction))
	}

	// The footer
	s += "\nPress q to quit.\n"

	// Send the UI for rendering
	return s
}

func renderTransaction(t ynabmodel.Transaction) string {
	return fmt.Sprintf("%-15s%-20s%-10s%10s  %-20s", t.Date.String(), t.AccountName, *t.CategoryName, t.Amount.Format(), t.Memo)
}
