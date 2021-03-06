package main

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMain(t *testing.T) {
	var buf bytes.Buffer
	out = &buf
	main()
	assert.Equal(t, "[a11y K8s abc]\n", buf.String())
}

func TestNumeronym(t *testing.T) {
	assert.Equal(t, "i18n", numeronym("internationalization"))
	assert.Equal(t, "c8y", numeronym("capability"))
	assert.Equal(t, "d7n", numeronym("dystopian"))
	assert.Equal(t, "car", numeronym("car"))
	assert.Equal(t, "at", numeronym("at"))
	assert.Equal(t, "i", numeronym("i"))
	assert.Empty(t, numeronym(""))
}

func TestNumeronymUnicode(t *testing.T) {
	assert.Equal(t, "c4n", numeronym("cπtion"))
	assert.Equal(t, "π6y", numeronym("πcπughty"))
	assert.Equal(t, "ππ", numeronym("ππ"))
	assert.Equal(t, "πππ", numeronym("πππ"))
}

func TestNumeronyms(t *testing.T) {
	assert.Equal(t, []string{"a11y", "K8s", "abc"}, numeronyms("accessibility", "Kubernetes", "abc"))
	assert.Equal(t, []string{"n4r", "b3y", "d2r"}, numeronyms("nectar", "bunny", "deer"))
	assert.Equal(t, []string{""}, numeronyms(""))
}

func TestNumeronymsUnicode(t *testing.T) {
	assert.Equal(t, []string{"a12y", "K9π", "πbc"}, numeronyms("aπ£ccessibility", "Kubernetesπ", "πbc"))
	assert.Equal(t, []string{"π4π", "π3π"}, numeronyms("ππππππ", "ππππ€π"))
	assert.Equal(t, []string{"π", "ππ", "πππ"}, numeronyms("π", "ππ", "πππ"))
}
