package config

type LangTypeError struct {
	msg string
}

func NewLangTypeError(msg string) *LangTypeError {
	return &LangTypeError{msg: msg}
}

func (e *LangTypeError) Error() string {
	return e.msg
}

type ReadDirError struct {
	msg string
}

func NewReadDirError(msg string) *ReadDirError {
	return &ReadDirError{msg: msg}
}

func (e *ReadDirError) Error() string {
	return e.msg
}
