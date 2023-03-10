package errno

type Errno struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

func (e Errno) Error() string {
	return e.Msg
}

func ParseErr(err error) (int, string) {
	if err == nil {
		return OK.Code, OK.Msg
	}

	switch e := err.(type) {
	case *Errno:
		return e.Code, e.Msg
	default:
		return InternalServerError.Code, err.Error()
	}
}
