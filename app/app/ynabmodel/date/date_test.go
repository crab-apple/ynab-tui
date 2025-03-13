package date

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"log"
	"testing"
	"time"
)

func TestCreateFromTime(t *testing.T) {
	t.Run("With midnight UTC time", func(t *testing.T) {
		times := []time.Time{
			makeTime("2024-01-02T00:00:00Z"),
			makeTime("2024-01-02T00:00:00+00:00"),
			makeTime("2024-01-02T00:00:00-00:00"),
		}
		for _, aTime := range times {
			t.Run(aTime.String(), func(t *testing.T) {
				_, err := FromTime(aTime)
				require.NoError(t, err)
			})
		}
	})

	t.Run("With non-midnight time", func(t *testing.T) {
		times := []time.Time{
			makeTime("2024-01-02T01:00:00Z"),
			makeTime("2024-01-02T00:01:00Z"),
			makeTime("2024-01-02T00:00:01Z"),
			makeTime("2024-01-02T00:00:00.000001Z"),
		}
		for _, aTime := range times {
			t.Run(aTime.String(), func(t *testing.T) {
				_, err := FromTime(aTime)
				require.Error(t, err)
			})
		}
	})

	t.Run("With non-UTC time", func(t *testing.T) {
		_, err := FromTime(makeTime("2024-01-02T00:00:00+01:00"))
		require.Error(t, err)
	})

	t.Run("With non-UTC time that just happens to equal midnight UTC", func(t *testing.T) {
		_, err := FromTime(makeTime("2024-01-02T00:01:00+01:00"))
		require.Error(t, err)
	})
}

func TestParse3ComponentString(t *testing.T) {
	date, err := Parse("2024-01-02")
	require.NoError(t, err)
	require.Equal(t, "2024-01-02", date.String())
}

func TestInvalidStrings(t *testing.T) {

	t.Run("Invalid month", func(t *testing.T) {
		_, err := Parse("2024-21-02")
		require.Error(t, err)
	})

	t.Run("Invalid day", func(t *testing.T) {
		_, err := Parse("2024-01-32")
		require.Error(t, err)
	})

	t.Run("Wrong format", func(t *testing.T) {
		_, err := Parse("2024-01-02T00:00:00Z")
		require.Error(t, err)
	})

}

func TestToday(t *testing.T) {
	d := Today()
	require.Equal(t, time.Now().Format("2006-01-02"), d.String())
}

func TestMidnight(t *testing.T) {
	d := makeDate("2023-04-05")
	require.Equal(t, makeTime("2023-04-05T00:00:00Z"), d.Midnight())
}

func TestMinusDays(t *testing.T) {

	t.Run("Subtraction works", func(t *testing.T) {

		type testCase struct {
			initial        Date
			daysSubtracted int
			expectedResult Date
		}

		describe := func(tc testCase) string {
			return fmt.Sprintf("%s minus %d days is %s", tc.initial.String(), tc.daysSubtracted, tc.expectedResult.String())
		}

		cases := []struct {
			initial        Date
			daysSubtracted int
			expectedResult Date
		}{
			{
				initial:        makeDate("2023-01-01"),
				daysSubtracted: 0,
				expectedResult: makeDate("2023-01-01"),
			},
			{
				initial:        makeDate("2023-01-30"),
				daysSubtracted: 1,
				expectedResult: makeDate("2023-01-29"),
			},
			{
				initial:        makeDate("2023-01-30"),
				daysSubtracted: 365,
				expectedResult: makeDate("2022-01-30"),
			},
			{
				initial:        makeDate("2025-01-30"),
				daysSubtracted: 365,
				expectedResult: makeDate("2024-01-31"),
			},
		}

		for _, c := range cases {
			t.Run(describe(c), func(t *testing.T) {
				d, err := c.initial.MinusDays(c.daysSubtracted)
				require.NoError(t, err)
				require.Equal(t, c.expectedResult, d)
			})
		}
	})

	t.Run("Negative input not allowed", func(t *testing.T) {
		_, err := makeDate("2020-01-01").MinusDays(-1)
		require.Error(t, err)
	})
}

func TestString(t *testing.T) {
	date, _ := FromTime(makeTime("2024-01-02T00:00:00Z"))
	require.Equal(t, "2024-01-02", date.String())

	date, _ = FromTime(makeTime("2024-01-03T00:00:00Z"))
	require.Equal(t, "2024-01-03", date.String())
}

func makeTime(s string) time.Time {
	t, err := time.Parse(time.RFC3339, s)
	if err != nil {
		log.Fatal(err)
	}
	return t
}

func makeDate(s string) Date {
	d, err := Parse(s)
	if err != nil {
		log.Fatal(err)
	}
	return d
}
