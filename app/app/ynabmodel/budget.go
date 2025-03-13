package ynabmodel

import (
	"github.com/google/uuid"
	"time"
)

type Budget struct {
	Id             uuid.UUID
	LastModifiedOn time.Time
}
