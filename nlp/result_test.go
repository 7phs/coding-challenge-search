package nlp

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLemmerResult(t *testing.T) {
	lemmer := &LemmerResult{}

	lemmer.AddWord("word1")
	lemmer.AddWord("word2")

	lemmer.AddLemma("lemma1")
	lemmer.AddLemma("lemma2")

	var exist LemResult

	exist = lemmer

	assert.Equal(t, []string{"word1", "word2"}, exist.Words())
	assert.Equal(t, []string{"lemma1", "lemma2"}, exist.Lemmas())
}
