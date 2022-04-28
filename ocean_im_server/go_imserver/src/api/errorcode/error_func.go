package errorcode

import "fmt"

type ErrorCode struct {
	Code    int
	Message string
}

func (e *ErrorCode) Error() string {
	return fmt.Sprintf("ErrorCode - code: %d, message: %s", e.Code, e.Message)
}

type ErrorSourceReason struct {
	Code        int
	Message     string
	ErrorReason error
}

func (e *ErrorSourceReason) Error() string {
	return fmt.Sprintf("ErrorSourceReason - code: %d, message: %s, error: %s", e.Code, e.Message, e.ErrorReason)
}

func NewErr(code ErrorCode, err error) *ErrorSourceReason {
	return &ErrorSourceReason{
		Code:        code.Code,
		Message:     code.Message,
		ErrorReason: err,
	}
}

func ParserErr(err error) (int, string) {
	if err == nil {
		return Ok.Code, Ok.Message
	}
	switch typed := err.(type) {
	case *ErrorCode:
		return typed.Code, typed.Message
	case *ErrorSourceReason:
		return typed.Code, typed.Message
	default:

	}
	return InternalServerError.Code, InternalServerError.Message
}
