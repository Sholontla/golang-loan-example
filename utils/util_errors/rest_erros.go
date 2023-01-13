package utilerrors

import (
	"fmt"

	"github.com/valyala/fasthttp"
)

type RestErr interface {
	Message() string
	Status() int
	Error() string
	Causes() []interface{}
}

type restErrors struct {
	ErrMessage string        `json:"message"`
	ErrStatus  int           `json:"status"`
	ErrError   string        `json:"error"`
	ErrCauses  []interface{} `json:"causes"`
}

func (e restErrors) Error() string {
	return fmt.Sprintf("message: %s - status: %d - error: %s - causes: %v",
		e.ErrMessage, e.ErrStatus, e.ErrError, e.ErrCauses)
}

func (e restErrors) Message() string {
	return e.ErrMessage
}

func (e restErrors) Status() int {
	return e.ErrStatus
}

func (e restErrors) Causes() []interface{} {
	return e.ErrCauses
}

func NewInvalidPasswordError(message string) restErrors {
	return restErrors{
		ErrMessage: message,
		ErrStatus:  fasthttp.StatusUnauthorized,
		ErrError:   "unauthorized",
	}
}

func NewBadRequestError(message string) restErrors {
	return restErrors{
		ErrMessage: message,
		ErrStatus:  fasthttp.StatusBadRequest,
		ErrError:   "bad_request",
	}
}

func NewNotFoundObject(message string) restErrors {
	return restErrors{
		ErrMessage: message,
		ErrStatus:  fasthttp.StatusNotFound,
		ErrError:   "not_found_request",
	}
}

func UserDataAllReadyExists(message string) restErrors {
	return restErrors{
		ErrMessage: message,
		ErrStatus:  fasthttp.StatusConflict,
		ErrError:   "user_data_exists",
	}
}
func NewInternalServerError(message string, err error) restErrors {
	result := restErrors{
		ErrMessage: message,
		ErrStatus:  fasthttp.StatusInternalServerError,
		ErrError:   "internal_server_error",
	}
	if err != nil {
		result.ErrCauses = append(result.ErrCauses, err.Error())
	}
	return result
}

func NewValidationAccountError(message string) restErrors {
	return restErrors{
		ErrMessage: message,
		ErrStatus:  fasthttp.StatusUnprocessableEntity,
		ErrError:   "unprocessable_entity",
	}
}

func NewValidationAccountTypeError(message string) restErrors {
	return restErrors{
		ErrMessage: message,
		ErrStatus:  fasthttp.StatusUnprocessableEntity,
		ErrError:   "unprocessable_entity",
	}
}
