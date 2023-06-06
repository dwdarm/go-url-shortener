package errors

type ErrDuplicate struct {
	s string
}

func (e *ErrDuplicate) Error() string {
	return e.s
}

func NewErrDuplicate(text string) error {
	return &ErrDuplicate{text}
}
