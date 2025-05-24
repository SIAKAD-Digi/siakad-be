package exception

type UnauthorizeError struct {
	message string
}

func (n *UnauthorizeError) Error() string {
	return n.message
}

func NewUnauthorizeError(err string) *UnauthorizeError {
	return &UnauthorizeError{message: err}
}
