package errmsg

type CustomError struct {
	Code   int
	Errors map[string][]string
	Msg    string
}

func (e *CustomError) Error() string {
	return e.Msg
}

func NewCustomErrors(errCode int, opts ...Option) *CustomError {
	err := &CustomError{
		Code:   errCode,
		Errors: make(map[string][]string),
		Msg:    "Permintaan Anda gagal diproses",
	}

	for _, opt := range opts {
		opt(err)
	}

	return err
}

func (e *CustomError) SetCode(code int) *CustomError {
	e.Code = code
	return e
}

func (e *CustomError) SetMessage(msg string) *CustomError {
	e.Msg = msg
	return e
}

func (e *CustomError) Add(field, msg string) *CustomError {
	e.Errors[field] = append(e.Errors[field], msg)
	return e
}

func (e *CustomError) HasErrors() bool {
	return len(e.Errors) > 0
}

type Option func(*CustomError)

func WithMessage(msg string) Option {
	return func(err *CustomError) {
		err.Msg = msg
	}
}

func errorCustomHandler(err *CustomError) (int, *CustomError) {
	return err.Code, err
}
