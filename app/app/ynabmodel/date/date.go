package date

import (
	"errors"
	"fmt"
	"log/slog"
	"time"
)

type Date struct {
	time time.Time
}

func (d Date) String() string {
	return d.time.String()[0:10]
}

func (d Date) Midnight() time.Time {
	return d.time
}

func (d Date) MinusDays(days int) (Date, error) {

	if days < 0 {
		return d, errors.New("days cannot be negative")
	}

	r, err := FromTime(d.time.AddDate(0, 0, -days))
	if err != nil {
		return d, err
	}
	return r, nil
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

func Today() Date {
	d, err := Parse(time.Now().Format("2006-01-02"))
	if err != nil {
		slog.Error(err.Error())
		panic(err)
	}
	return d
}
