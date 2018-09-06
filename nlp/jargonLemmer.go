package nlp

import (
	"strings"

	"github.com/clipperhouse/jargon"
	"github.com/clipperhouse/jargon/stackexchange"
)

const (
	MaxGramLen = 64
)

const (
	DictStackExchange DictType = iota + 1
)

type DictType int

func (o DictType) Dictionary() jargon.Dictionary {
	switch o {
	case DictStackExchange:
		return stackexchange.Dictionary
	default:
		return stackexchange.Dictionary
	}
}

type JargonLemmer struct {
	lemmer *jargon.Lemmatizer
}

func NewJargonLemmer(dict DictType) *JargonLemmer {
	return &JargonLemmer{
		lemmer: jargon.NewLemmatizer(dict.Dictionary(), MaxGramLen),
	}
}

func (o JargonLemmer) Parse(line string) LemResult {
	line = strings.ToLower(line)

	result := LemmerResult{}
	tokens := jargon.Tokenize(
		strings.NewReader(line))

	for {
		word := tokens.Next()
		if word == nil {
			break
		}
		if word.IsSpace() || word.IsPunct() {
			continue
		}

		result.AddWord(word.String())
	}

	lemmas := o.lemmer.Lemmatize(
		jargon.Tokenize(
			strings.NewReader(line)))
	for {
		lemma := lemmas.Next()
		if lemma == nil {
			break
		}
		if lemma.IsSpace() || lemma.IsPunct() {
			continue
		}

		result.AddLemma(lemma.String())
	}

	return result
}
