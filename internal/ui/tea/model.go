package tea

import (
	"github.com/samber/lo"
	uimodel "ynabtui/internal/ui/model"
	"ynabtui/internal/ynabapi"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	btable "github.com/evertras/bubble-table/table"
)

const (
	fixedVerticalMargin = 2
)

type updateScreenMsg struct {
	screen any
}

type Model struct {
	uiModel   uimodel.UI
	flexTable btable.Model
}

func InitialModel(api ynabapi.YnabApi) Model {

	uiModel := uimodel.NewUI(api)

	return Model{
		uiModel: uiModel,
		flexTable: btable.New([]btable.Column{
			btable.NewFlexColumn("a", "Pending", 1),
			btable.NewFlexColumn("b", "Pending", 1),
			btable.NewFlexColumn("c", "Pending", 1),
		}).WithStaticFooter("A footer!"),
	}
}

func (m Model) Init() tea.Cmd {
	return func() tea.Msg {
		screen := m.uiModel.FirstLoad()
		return updateScreenMsg{
			screen: screen,
		}
	}
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

	var (
		cmd  tea.Cmd
		cmds []tea.Cmd
	)

	m.flexTable, cmd = m.flexTable.Update(msg)
	cmds = append(cmds, cmd)

	switch msg := msg.(type) {

	case updateScreenMsg:
		switch msg.screen.(type) {

		case uimodel.TransactionsScreen:
			m.flexTable = updateDisplayTable(m.flexTable, msg.screen.(uimodel.TransactionsScreen).Table())
		}

	case tea.KeyMsg:
		switch msg.String() {
		case "q":
			cmds = append(cmds, tea.Quit)
		}

	case tea.WindowSizeMsg:
		m.flexTable = resizeDisplayTable(m.flexTable, msg.Width, msg.Height)
	}

	return m, tea.Batch(cmds...)
}

func updateDisplayTable(prev btable.Model, table uimodel.Table) btable.Model {
	result := prev.
		HeaderStyle(lipgloss.NewStyle().Bold(true).AlignHorizontal(lipgloss.Left)).
		WithColumns(lo.Map(table.Columns, func(column uimodel.Column, _ int) btable.Column {
			displayColumn := btable.NewFlexColumn(column.Key, column.Display, 1).
				WithStyle(lipgloss.NewStyle().AlignHorizontal(lipgloss.Left))

			if column.CellAlign == uimodel.AlignRight {
				displayColumn = displayColumn.WithStyle(displayColumn.Style().AlignHorizontal(lipgloss.Right))
			}
			return displayColumn
		}))

	result = result.WithRows(lo.Map(table.Rows, func(row uimodel.Row, _ int) btable.Row {
		return btable.NewRow(lo.MapValues(row, func(value string, key string) interface{} {
			return value
		}))
	}))
	return result
}

func resizeDisplayTable(table btable.Model, width int, h int) btable.Model {
	return table.
		WithTargetWidth(width).
		WithMinimumHeight(h - fixedVerticalMargin)
}

func (m Model) View() string {
	strs := []string{
		m.flexTable.View(),
		"Press q to quit",
	}

	return lipgloss.JoinVertical(lipgloss.Left, strs...) + "\n"
}
