package nlp

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
