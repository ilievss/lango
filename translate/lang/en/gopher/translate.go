package gopher

type WordTranslator interface {
	TranslateWord(string) (string, error)
}

type SentenceTranslator interface {
	TranslateSentence(string) (string, error)
}

type EnTranslator interface {
	Translate(string) (string, error)
}
