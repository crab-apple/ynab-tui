package ynabmodel

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCreateMoney(t *testing.T) {
	_, err := NewMoneyFromThousandths(12340)
	require.NoError(t, err)
}

func TestNoThousandthsAccepted(t *testing.T) {
	_, err := NewMoneyFromThousandths(12345)
	require.Error(t, err)
}

func TestShouldParseStringWithCents(t *testing.T) {

	// When
	m, err := NewMoneyFromString("12.34")

	// Then
	require.NoError(t, err)
	assert.Equal(t, int64(1234), m.cents)

	// When
	m, err = NewMoneyFromString("0.34")

	// Then
	require.NoError(t, err)
	assert.Equal(t, int64(34), m.cents)
}

func TestShouldNotAcceptInvalidStrings(t *testing.T) {

	expectedMsg := "Money strings must contain at least one digit to the left of the point and exactly two digits to the right"

	var invalidStrings = []string{
		"1111",
		"1111.",
		"1111.1",
		"1111.123",
		".12",
		"1.12.34",
	}

	for _, str := range invalidStrings {
		t.Run(str, func(t *testing.T) {
			// When
			_, err := NewMoneyFromString(str)

			// Then
			require.Error(t, err)
			assert.Equal(t, expectedMsg, err.Error())
		})
	}
}

func TestFormat(t *testing.T) {

	tests := map[string]struct {
		input  int64
		output string
	}{
		"cents":                               {input: 12340, output: "12.34"},
		"tenths":                              {input: 12300, output: "12.30"},
		"units":                               {input: 12000, output: "12.00"},
		"tens":                                {input: 10000, output: "10.00"},
		"only cents":                          {input: 20, output: "0.02"},
		"only tenths":                         {input: 200, output: "0.20"},
		"thousandths separator":               {input: 111222333444550, output: "111,222,333,444.55"},
		"negative":                            {input: -12340, output: "-12.34"},
		"negative with thousandths separator": {input: -111222330, output: "-111,222.33"},
		"negative with zeroes":                {input: -30, output: "-0.03"},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			m, _ := NewMoneyFromThousandths(tc.input)
			require.Equal(t, tc.output, m.Format())
		})
	}
}
