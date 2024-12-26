package model

type Table struct {
	Columns []Column
	Rows    []Row
}

type Column struct {
	key     string
	display string
}

type Row map[string]string
