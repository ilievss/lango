package gopher

type FailedToTranslateError struct {
	Message string
}

func (e FailedToTranslateError) Error() string {
	return e.Message
}
