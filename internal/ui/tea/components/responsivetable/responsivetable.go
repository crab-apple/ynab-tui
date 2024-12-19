package responsivetable

import (
	"github.com/charmbracelet/bubbles/table"
)

type Model struct {
	table   table.Model
	width   int
	columns []Column
}

type Column struct {
	Title string
}

func (m *Model) SetColumns(c []Column) {
	m.columns = c
	m.setColumnsInner()
}

func (m *Model) setColumnsInner() {

	styles := table.DefaultStyles() // TODO get a reference of the real styles if we want to change them.

	tcs := make([]table.Column, 0)
	for _, c := range m.columns {
		spaceAvailable := m.width / len(m.columns)
		width := spaceAvailable - styles.Header.GetHorizontalPadding()
		tcs = append(tcs, table.Column{
			Title: c.Title,
			Width: width,
		})
	}
	m.table.SetColumns(tcs)
}

func (m *Model) SetRows(c []table.Row) {
	m.table.SetRows(c)
}

func (m *Model) SetWidth(w int) {
	m.width = w
	m.table.SetWidth(w)
	m.setColumnsInner()
}

func (m *Model) SetHeight(h int) {
	m.table.SetHeight(h)
}

func New(opts ...table.Option) Model {
	t := table.New(opts...)
	return Model{table: t}
}

func (m *Model) View() string {
	return m.table.View()
}
