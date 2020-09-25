package rule

type RuleNotApplicableError struct {
	Message string
}

func (e RuleNotApplicableError) Error() string {
	return e.Message
}
