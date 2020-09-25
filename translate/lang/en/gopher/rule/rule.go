package rule

const empty = ""
const apostrophe = "'"

type WordRule interface {
	Apply(word string) (string, error)
}

var rules = []WordRule{
	VowelPrefixWordRule,
	XrPrefixWordRule,
	ConsonantsPrefixWordRule,
	ConsonantsQuPrefixWordRule,
}
