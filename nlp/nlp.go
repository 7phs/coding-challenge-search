package nlp

var (
	Lem *Lemmer
)

func Init() {
	Lem = NewLemmer(DictStackExchange)
}
