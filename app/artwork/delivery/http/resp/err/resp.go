package err

import (
	"net/http"
	error2 "pixstall-artwork/domain/error"
)

type Response struct {
	Message string `json:"message"`
}

func NewErrorResponse(err error) (int, interface{}) {
	switch err {
	case error2.NotFoundError:
		return http.StatusNotFound, Response{
			Message:err.Error(),
		}
	case error2.UnAuthError:
		return http.StatusUnauthorized, Response{
			Message:err.Error(),
		}
	case error2.UnknownError:
		return http.StatusInternalServerError, Response{
			Message:err.Error(),
		}
	}
	return http.StatusInternalServerError, Response{
		Message:err.Error(),
	}
}