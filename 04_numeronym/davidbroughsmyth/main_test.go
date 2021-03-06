package main

import (
	"bytes"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMainOutput(t *testing.T) {
	var buf bytes.Buffer
	out = &buf
	main()
	expected := strconv.Quote("[a11y K8s abc]\n")
	actual := strconv.Quote(buf.String())
	assert.Equal(t, expected, actual, "Unexpected output from main()")
}

var tests = map[string]struct {
	input, want []string
}{
	"empty":        {input: []string{}, want: []string{}},
	"SpecialChars": {input: []string{"@@!##%&&"}, want: []string{"@6&"}},
	"3unicode":     {input: []string{"ππΈπ"}, want: []string{"ππΈπ"}},
	"4unicode":     {input: []string{"πππε€§"}, want: []string{"π2ε€§"}},
	"MixCodes":     {input: []string{"π@#Xy ε€§Z2θ΅"}, want: []string{"π8θ΅"}},
	"Alphanumeric": {
		input: []string{"8abcABC5Z", "cAt", "b1Rd"},
		want:  []string{"87Z", "cAt", "b2d"},
	},
	"MaryPoppins": {
		input: []string{"It's", "supercalifragilisticexpialidocious"},
		want:  []string{"I2s", "s32s"},
	},
}

func TestNumeronyms(t *testing.T) {
	for name, test := range tests {
		test := test
		t.Run(name, func(t *testing.T) {
			got := numeronyms(test.input...)
			assert.Equal(t, test.want, got)
		})
	}
}
