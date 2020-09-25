package rule

type WordRule interface {
	Apply(word string) (string, error)
}
