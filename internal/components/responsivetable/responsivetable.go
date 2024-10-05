package responsivetable

import (
	"github.com/charmbracelet/bubbles/table"
)

type Model struct {
	table table.Model
}

type Column struct {
	Title string
	Width int32
}

func (m *Model) SetColumns(c []table.Column) {
	m.table.SetColumns(c)
}

func (m *Model) SetRows(c []table.Row) {
	m.table.SetRows(c)
}

func (m *Model) SetWidth(w int) {
	m.table.SetWidth(w)
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
