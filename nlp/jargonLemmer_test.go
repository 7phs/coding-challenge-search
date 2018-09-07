package nlp

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewJargonLemmer(t *testing.T) {
	var lemmer Lemmer

	for _, dict := range []DictType{DictStackExchange, 0} {
		lemmer = NewJargonLemmer(dict)

		exist := lemmer.Parse("")
		assert.Len(t, exist.Words(), 0)
		assert.Len(t, exist.Lemmas(), 0)

		exist = lemmer.Parse("hello world")
		expected := []string{"hello", "world"}
		assert.Equal(t, expected, exist.Words(), "words")
		assert.Equal(t, expected, exist.Lemmas(), "lemmas")
	}
}
