package utilerrors

import (
	"fmt"

	"github.com/valyala/fasthttp"
)

type MiddleWareErr interface {
	Message() string
	Status() int
	Error() string
	Causes() []interface{}
}

type middleWareErrors struct {
	ErrMessage string        `json:"message"`
	ErrStatus  int           `json:"status"`
	ErrError   string        `json:"error"`
	ErrCauses  []interface{} `json:"causes"`
}

func (e middleWareErrors) Error() string {
	return fmt.Sprintf("message: %s - status: %d - error: %s - causes: %v",
		e.ErrMessage, e.ErrStatus, e.ErrError, e.ErrCauses)
}

func (e middleWareErrors) Message() string {
	return e.ErrMessage
}

func (e middleWareErrors) Status() int {
	return e.ErrStatus
}

func (e middleWareErrors) Causes() []interface{} {
	return e.ErrCauses
}

func NewMiddleWareError(message string) middleWareErrors {
	return middleWareErrors{
		ErrMessage: message,
		ErrStatus:  fasthttp.StatusUnauthorized,
		ErrError:   "unauthorized",
	}
}
