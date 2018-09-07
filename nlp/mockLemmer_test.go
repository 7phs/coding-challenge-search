package nlp

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewMockLemmer(t *testing.T) {
	var lemmer Lemmer

	lemmer = NewMockLemmer()

	exist := lemmer.Parse("")
	assert.Len(t, exist.Words(), 0)
	assert.Len(t, exist.Lemmas(), 0)

	exist = lemmer.Parse("hello world")
	expected := []string{"hello", "world"}
	assert.Equal(t, expected, exist.Words(), "words")
	assert.Equal(t, expected, exist.Lemmas(), "lemmas")
}
