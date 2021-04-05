package error

import "strconv"

type DomainError int

func (e DomainError) Error() string {
	switch e {
	case NotFoundError:
		return "NotFoundError"
	case UnAuthError:
		return "UnAuthError"
	case BadRequestError:
		return "BadRequestError"
	case ForbiddenError:
		return "ForbiddenError"
	case UnknownError:
		return "UnknownError"
	default:
		return strconv.Itoa(int(e))
	}
}

const (
	NotFoundError   DomainError = 10
	UnAuthError     DomainError = 11
	BadRequestError DomainError = 12
	ForbiddenError  DomainError = 13
	UnknownError    DomainError = 99
)
