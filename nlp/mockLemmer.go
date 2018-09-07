package nlp

import (
	"strings"
)

type MockLemmer struct{}

func NewMockLemmer() *MockLemmer {
	return &MockLemmer{}
}

func (o MockLemmer) Parse(line string) LemResult {
	if len(line) == 0 {
		return &LemmerResult{}
	}

	words := strings.Split(line, " ")

	return &LemmerResult{
		words:  words,
		lemmas: words,
	}
}
