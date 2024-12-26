package tea

import (
	"fmt"
	"github.com/charmbracelet/bubbles/table"
	btea "github.com/charmbracelet/bubbletea"
	"github.com/samber/lo"
	"log/slog"
	uimodel "ynabtui/internal/ui/model"
	"ynabtui/internal/ui/tea/components/responsivetable"
	"ynabtui/internal/ynabapi"
)

type Model struct {
	uiModel uimodel.UI

	table responsivetable.Model
}

type updateScreenMsg struct {
	screen any
}

func InitialModel(api ynabapi.YnabApi) Model {

	uiModel := uimodel.NewUI(api)

	t := responsivetable.New(
		table.WithFocused(true),
	)

	t.SetHeight(15)

	return Model{
		uiModel: uiModel,
		table:   t,
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
		case uimodel.TransactionsScreen:

			m.table.SetColumns(lo.Map(screen.Table().Columns, func(col uimodel.Column, _ int) responsivetable.Column {
				return responsivetable.Column{Title: col.Display}
			}))

			m.table.SetRows(
				lo.Map(screen.Table().Rows, func(row uimodel.Row, _ int) table.Row {
					return lo.Map(screen.Table().Columns, func(col uimodel.Column, _ int) string {
						return row[col.Key]
					})
				}))
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
