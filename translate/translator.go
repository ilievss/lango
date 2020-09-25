package translate

type Translator interface {
	Translate(string) (string, error)
}
