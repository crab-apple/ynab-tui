package date

import (
	"fmt"
	"time"
)

type Date struct {
	time time.Time
}

func (d Date) String() string {
	return d.time.String()[0:10]
}

func FromTime(t time.Time) (Date, error) {
	canonical := time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, time.UTC)
	if t.UTC() != canonical {
		return Date{}, fmt.Errorf("time must be at midnight UTC, received %s", t)
	}
	return Date{t}, nil
}

func Parse(s string) (Date, error) {
	t, err := time.Parse(time.DateOnly, s)
	if err != nil {
		return Date{}, err
	}
	return FromTime(t)
}
