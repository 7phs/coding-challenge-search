package nlp

import (
	"strings"
)

type MockLemmer struct{}

func NewMockLemmer() *MockLemmer {
	return &MockLemmer{}
}

func (o MockLemmer) Parse(line string) LemResult {
	words := strings.Split(line, " ")

	return &LemmerResult{
		words:  words,
		lemmas: words,
	}
}
