package nlp

import (
	"strings"

	"github.com/7phs/coding-challenge-search/model"
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

type Lemmer struct {
	lemmer *jargon.Lemmatizer
}

func NewLemmer(dict DictType) *Lemmer {
	return &Lemmer{
		lemmer: jargon.NewLemmatizer(dict.Dictionary(), MaxGramLen),
	}
}

func (o Lemmer) Parse(line string) model.LemResult {
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

type LemmerResult struct {
	words  []string
	lemmas []string
}

func (o *LemmerResult) AddWord(word string) {
	o.words = append(o.words, word)
}

func (o *LemmerResult) AddLemma(lemma string) {
	o.lemmas = append(o.lemmas, lemma)
}

func (o LemmerResult) Words() []string {
	return o.words
}

func (o LemmerResult) Lemmas() []string {
	return o.lemmas
}
