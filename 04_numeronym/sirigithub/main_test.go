package main

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMainOutput(t *testing.T) {
	var buf bytes.Buffer
	out = &buf
	main()
	actual := buf.String()
	expected := "[a11y K8s abc]\n"
	assert.Equal(t, expected, actual)

}
func TestNumeronyms(t *testing.T) {
	tests := []struct {
		description string
		input,
		expected []string
	}{
		{description: "Empty string", input: []string{""}, expected: []string{""}},
		{description: "Short strings", input: []string{"sss", "ab", "a"}, expected: []string{"sss", "ab", "a"}},
		{description: "Alphanumeric string", input: []string{"sss1245sdfg"}, expected: []string{"s9g"}},
		{description: "Back slashes and spaces", input: []string{"sss\\", "A long string with spaces"},
			expected: []string{"s2\\", "A23s"}},
		{description: "Emoji strings", input: []string{"ππ·ππ·ππ", "πππ"},
			expected: []string{"π4π", "πππ"}},
		{description: "Unicode Strings", input: []string{"ζ₯ζ¬θͺζ₯ζ¬θͺ", "Γ€ΓΆΓβ¬βΓΌΓΌΓΌΓΆΓ€ΓΓ"},
			expected: []string{"ζ₯4θͺ", "Γ€10Γ"}},
		{description: "Strings with special characters",
			input:    []string{"a##67&a$", "**(]]"},
			expected: []string{"a6$", "*3]"}},
	}
	for _, test := range tests {
		test := test
		t.Run(test.description, func(t *testing.T) {
			actual := numeronyms(test.input...)
			expected := test.expected
			assert.Equal(t, actual, expected, "actual %v but expected %v", actual, expected)
		})
	}
}
