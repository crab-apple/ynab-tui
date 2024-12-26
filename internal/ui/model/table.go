package model

type Table struct {
	Columns []Column
	Rows    []Row
}

type Column struct {
	Key     string
	Display string
}

type Row map[string]string
