package tea

import (
	"fmt"
	"github.com/charmbracelet/bubbles/table"
	btea "github.com/charmbracelet/bubbletea"
	"github.com/samber/lo"
	"log/slog"
	tea "ynabtui/internal/ui/model"
	"ynabtui/internal/ui/tea/components/responsivetable"
	"ynabtui/internal/ynabapi"
	"ynabtui/internal/ynabmodel"
)

type Model struct {
	uiModel tea.UI

	transactions []ynabmodel.Transaction
	table        responsivetable.Model
}

type updateScreenMsg struct {
	screen any
}

func InitialModel(api ynabapi.YnabApi) Model {

	uiModel := tea.NewUI(api)

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

func (m Model) Init() btea.Cmd {
	return func() btea.Msg {
		screen := m.uiModel.FirstLoad()
		return updateScreenMsg{
			screen: screen,
		}
	}
}
func (m Model) Update(msg btea.Msg) (btea.Model, btea.Cmd) {

	slog.Debug("Received message", "type", fmt.Sprintf("%T", msg), "value", msg)

	switch msg := msg.(type) {

	case updateScreenMsg:
		switch screen := msg.screen.(type) {
		case tea.TransactionsScreen:
			m.transactions = screen.Transactions
			rows := lo.Map(m.transactions, func(item ynabmodel.Transaction, i int) table.Row {
				return makeTransactionRow(item)
			})
			m.table.SetRows(rows)
		}

	case btea.WindowSizeMsg:
		m.table.SetWidth(msg.Width)
		m.table.SetHeight(msg.Height)

	// Is it a key press?
	case btea.KeyMsg:

		// Cool, what was the actual key pressed?
		switch msg.String() {

		// These keys should exit the program.
		case "ctrl+c", "q":
			return m, btea.Quit
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
