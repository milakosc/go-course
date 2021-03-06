package main

import (
	"bytes"
	"fmt"
	"reflect"
	"strconv"
	"testing"
)

func TestMainOutput(t *testing.T) {
	var buf bytes.Buffer
	out = &buf

	main()

	expected := strconv.Quote(`a:2
b:1
`)
	actual := strconv.Quote(buf.String())

	if expected != actual {
		t.Errorf("Unexpected output in main()\nexpected: %q\nactual: %q", expected, actual)
	}
}

func TestLetters(t *testing.T) {
	tests := map[string]struct {
		input    string
		expected map[rune]int
	}{
		"Empty string": {
			input:    "",
			expected: map[rune]int{}},
		"ASCII string": {
			input: "  ``~8I,dzye[uY6<mCh<n 9Otefp0fX0-@2<C)z)}.go-Hq{n]LX 8uKnRxj (92@08b9P ",
			expected: map[rune]int{
				'@': 2, 'C': 2, 'H': 1, 'I': 1, 'K': 1, 'L': 1, 'O': 1, 'P': 1, 'R': 1, 'X': 2, 'Y': 1, '[': 1,
				' ': 6, '(': 1, ')': 2, ',': 1, '-': 2, '.': 1, '0': 3, '2': 2, '6': 1, '8': 3, '9': 3, '<': 3,
				'p': 1, 'q': 1, 't': 1, 'u': 2, 'x': 1, 'y': 1, 'z': 2, '{': 1, '}': 1, '~': 1, 'o': 1,
				']': 1, '`': 2, 'b': 1, 'd': 1, 'e': 2, 'f': 2, 'g': 1, 'h': 1, 'j': 1, 'm': 1, 'n': 3,
			}},
		"Rune string": {
			input: "θε«πβ²ΰ§²β―β·$β¬β·ββ²β¬β β¦πΆππππβ±­αΈΈαΉββ€β‘μ±π€π§μ·‘γΌͺΞ½γ°πλΎ¬β",
			expected: map[rune]int{
				'γ°': 1, 'γΌͺ': 1, 'ε«': 1, 'θ': 1, 'λΎ¬': 1, 'μ±': 1, 'μ·‘': 1, 'π': 1, 'πΆ': 1, 'π': 1,
				'β¦': 1, 'β¬': 1, 'β²': 1, 'β·': 1, 'β': 1, 'β': 1, 'β': 1, 'β€': 1, 'β‘': 1, 'β±­': 1,
				'$': 1, 'Ξ½': 1, 'ΰ§²': 1, 'αΈΈ': 1, 'αΉ': 1, 'β¬': 1, 'β―': 1, 'β²': 1, 'β·': 1, 'β ': 1,
				'π': 1, 'π': 1, 'π§': 1, 'π': 1, 'π': 1, 'π€': 1,
			}},
		"String with esc chars": {
			input: "\"Kia ora\"",
			expected: map[rune]int{
				'i': 1, 'o': 1, 'r': 1, ' ': 1, '"': 2, 'K': 1, 'a': 2,
			}},
	}

	for name, test := range tests {
		testData := test
		t.Run(name, func(t *testing.T) {
			actual := letters(testData.input)
			if !reflect.DeepEqual(testData.expected, actual) {
				fmt.Println(actual)
				t.Errorf("Unexpected output in main()\nexpected: %q\nactual: %v", testData.expected, actual)
			}
		})
	}
}

func TestSortLetters(t *testing.T) {
	tests := map[string]struct {
		input    map[rune]int
		expected []string
	}{
		"Empty map": {
			input:    map[rune]int{},
			expected: []string{}},
		"ASCII map": {
			input: map[rune]int{
				'@': 2, 'C': 2, 'H': 1, 'I': 1, 'K': 1, 'L': 1, 'O': 1, 'P': 1, 'R': 1, 'X': 2, 'Y': 1, '[': 1,
				' ': 6, '(': 1, ')': 2, ',': 1, '-': 2, '.': 1, '0': 3, '2': 2, '6': 1, '8': 3, '9': 3, '<': 3,
				'p': 1, 'q': 1, 't': 1, 'u': 2, 'x': 1, 'y': 1, 'z': 2, '{': 1, '}': 1, '~': 1, 'o': 1,
				']': 1, '`': 2, 'b': 1, 'd': 1, 'e': 2, 'f': 2, 'g': 1, 'h': 1, 'j': 1, 'm': 1, 'n': 3,
			},
			expected: []string{
				" :6", "(:1", "):2", ",:1", "-:2", ".:1", "0:3", "2:2", "6:1", "8:3", "9:3", "<:3", "@:2", "C:2",
				"H:1", "I:1", "K:1", "L:1", "O:1", "P:1", "R:1", "X:2", "Y:1", "[:1", "]:1", "`:2", "b:1", "d:1",
				"e:2", "f:2", "g:1", "h:1", "j:1", "m:1", "n:3", "o:1", "p:1", "q:1", "t:1", "u:2", "x:1", "y:1",
				"z:2", "{:1", "}:1", "~:1",
			}},
		"Rune map": {
			input: map[rune]int{
				'γ°': 1, 'γΌͺ': 1, 'ε«': 1, 'θ': 1, 'λΎ¬': 1, 'μ±': 1, 'μ·‘': 1, 'π': 1, 'πΆ': 1, 'π': 1,
				'β¦': 1, 'β¬': 1, 'β²': 1, 'β·': 1, 'β': 1, 'β': 1, 'β': 1, 'β€': 1, 'β‘': 1, 'β±­': 1,
				'$': 1, 'π': 1, 'π§': 1, 'π': 1, 'π': 1, 'π€': 1, 'β―': 1, 'β²': 1, 'β·': 1, 'β ': 1,
				'π': 1, 'Ξ½': 1, 'ΰ§²': 1, 'αΈΈ': 1, 'αΉ': 1, 'β¬': 1,
			},
			expected: []string{
				"$:1", "Ξ½:1", "ΰ§²:1", "αΈΈ:1", "αΉ:1", "β¬:1", "β―:1", "β²:1", "β·:1", "β :1", "β¦:1", "β¬:1",
				"β²:1", "β·:1", "β:1", "β:1", "β:1", "β€:1", "β‘:1", "β±­:1", "γ°:1", "γΌͺ:1", "ε«:1", "θ:1", "λΎ¬:1", "μ±:1",
				"μ·‘:1", "π:1", "πΆ:1", "π:1", "π:1", "π:1", "π§:1", "π:1", "π:1", "π€:1",
			}},
		"Map with esc chars": {
			input: map[rune]int{
				'i': 1, 'o': 1, 'r': 1, ' ': 1, '"': 2, 'K': 1, 'a': 2,
			},
			expected: []string{
				" :1", "\":2", "K:1", "a:2", "i:1", "o:1", "r:1",
			}},
	}

	for name, test := range tests {
		testData := test
		t.Run(name, func(t *testing.T) {
			actual := sortLetters(testData.input)
			if !reflect.DeepEqual(testData.expected, actual) {
				t.Errorf("Unexpected output in main()\nexpected: %s\nactual: %s", testData.expected, actual)
			}
		})
	}
}
