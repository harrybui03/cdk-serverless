package app

import "net/http"

type Error struct {
	Code       int    `json:"code"`
	Message    string `json:"message"`
	DebugError error  `json:"-"`
	ErrorCode  string `json:"error_code,omitempty" yaml:"error_code,omitempty"`
}

type Response struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

type SimpleResponse struct {
	Message string `json:"message"`
}

func (e Error) Error() string {
	return e.Message
}

func NewInternalError(err error, msg ...string) *Error {
	if len(msg) > 0 {
		return &Error{
			Code:       500,
			Message:    msg[0],
			DebugError: err,
		}
	}

	return &Error{
		Code:       500,
		Message:    "Internal Server Error",
		DebugError: err,
	}
}

func NewBadRequestError(err error, msg ...string) *Error {
	if len(msg) > 0 {
		return &Error{
			Code:       400,
			Message:    msg[0],
			DebugError: err,
		}
	}

	return &Error{
		Code:       400,
		Message:    "Bad Request",
		DebugError: err,
	}
}

func NewNotFoundError(err error, msg ...string) *Error {
	if len(msg) > 0 {
		return &Error{
			Code:       http.StatusNotFound,
			Message:    msg[0],
			DebugError: err,
		}
	}

	return &Error{
		Code:       http.StatusNotFound,
		Message:    "Not Found",
		DebugError: err,
	}
}
