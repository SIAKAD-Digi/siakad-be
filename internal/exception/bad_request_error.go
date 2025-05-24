package exception

type BadRequestError struct {
	message string
}

func (n *BadRequestError) Error() string {
	return n.message
}

func NewBadRequestError(err string) *BadRequestError {
	return &BadRequestError{message: err}
}
