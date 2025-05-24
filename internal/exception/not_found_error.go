package exception

type NotFoundError struct {
	message string
}

func (n *NotFoundError) Error() string {
	return n.message
}

func NewNotFoundError(err string) *NotFoundError {
	return &NotFoundError{message: err}
}
