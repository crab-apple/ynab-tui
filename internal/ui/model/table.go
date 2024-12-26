package model

type Table struct {
	Columns []Column
	Rows    []Row
}

type Column struct {
	Key       string
	Display   string
	CellAlign Align
}

type Row map[string]string

type Align int

const (
	AlignLeft  Align = iota
	AlignRight Align = iota
)
