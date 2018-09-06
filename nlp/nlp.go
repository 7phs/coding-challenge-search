package nlp

var (
	Lem *JargonLemmer
)

type Lemmer interface {
	Parse(string) LemResult
}

type LemResult interface {
	Words() []string
	Lemmas() []string
}

func Init() {
	Lem = NewJargonLemmer(DictStackExchange)
}
