package tea

import (
	"github.com/samber/lo"
	uimodel "ynabtui/internal/ui/model"
	"ynabtui/internal/ynabapi"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/evertras/bubble-table/table"
)

const (
	fixedVerticalMargin = 2
)

type updateScreenMsg struct {
	screen any
}

type Model struct {
	uiModel   uimodel.UI
	flexTable table.Model
}

func InitialModel(api ynabapi.YnabApi) Model {

	uiModel := uimodel.NewUI(api)

	return Model{
		uiModel: uiModel,
		flexTable: table.New([]table.Column{
			table.NewFlexColumn("a", "Pending", 1),
			table.NewFlexColumn("b", "Pending", 1),
			table.NewFlexColumn("c", "Pending", 1),
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

			screen := msg.screen.(uimodel.TransactionsScreen)

			m.flexTable = m.flexTable.WithColumns(lo.Map(screen.Table().Columns, func(column uimodel.Column, _ int) table.Column {
				return table.NewFlexColumn(column.Key, column.Display, 1)
			}))

			m.flexTable = m.flexTable.WithRows(lo.Map(screen.Table().Rows, func(row uimodel.Row, _ int) table.Row {
				return table.NewRow(lo.MapValues(row, func(value string, key string) interface{} {
					return value
				}))
			}))
		}

	case tea.KeyMsg:
		switch msg.String() {
		case "q":
			cmds = append(cmds, tea.Quit)
		}

	case tea.WindowSizeMsg:
		m.flexTable = m.flexTable.
			WithTargetWidth(msg.Width).
			WithMinimumHeight(msg.Height - fixedVerticalMargin)
	}

	return m, tea.Batch(cmds...)
}

func (m Model) View() string {
	strs := []string{
		m.flexTable.View(),
		"Press q to quit",
	}

	return lipgloss.JoinVertical(lipgloss.Left, strs...) + "\n"
}
