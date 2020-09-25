package rule

type WordRule interface {
	Apply(word string) (string, error)
}

var rules = []WordRule{
	VowelPrefixWordRule,
	XrPrefixWordRule,
	ConsonantsPrefixWordRule,
	ConsonantsQuPrefixWordRule,
}
