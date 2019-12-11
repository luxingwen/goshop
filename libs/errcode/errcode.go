package errcode

type ErrorCode struct {
	code int64
	err  string
}

func (e ErrorCode) Error() string {
	return e.err
}

func New(err string) ErrorCode {
	return ErrorCode{code: 999, err: err}
}

func NewErrcode(code int64, err string) ErrorCode {
	return ErrorCode{code: code, err: err}
}

func (e ErrorCode) Code() int64 {
	return e.code
}

func ErrCode(err error) int64 {
	switch err.(type) {
	case ErrorCode:
		return err.(ErrorCode).Code()
	default:
		return 999
	}
	return 999
}
