package ynab

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCreateMoney(t *testing.T) {
	_, err := NewMoney(12340)
	require.NoError(t, err)
}

func TestNoThousandthsAccepted(t *testing.T) {
	_, err := NewMoney(12345)
	require.Error(t, err)
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
			m, _ := NewMoney(tc.input)
			require.Equal(t, tc.output, m.Format())
		})
	}
}
