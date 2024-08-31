package ynab

import (
	"github.com/google/uuid"
)

type Account struct {
	Id   uuid.UUID
	Name string
}
