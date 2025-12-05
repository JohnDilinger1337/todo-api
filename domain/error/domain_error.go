package error

type DomainError struct {
	Code string
}

func (e *DomainError) Error() string {
	return e.Code
}

func New(code string) *DomainError {
	return &DomainError{
		Code: code,
	}
}
