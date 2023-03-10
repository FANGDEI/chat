package errno

import (
	"chat/internal/app/service"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Errno struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

func (e Errno) Error() string {
	return e.Msg
}
func ServerErr(errno *Errno, cause string) error {
	err, _ := status.New(codes.Aborted, cause).WithDetails(
		&service.Error{
			Code: int32(errno.Code),
			Msg:  errno.Msg,
		},
	)
	return err.Err()
}

func ParseErr(err error) (int, string) {
	if err == nil {
		return OK.Code, OK.Msg
	}

	if rpcErr, ok := status.FromError(err); ok {
		rpcErr.Message()
		for _, v := range rpcErr.Details() {
			detail := v.(*service.Error)
			return int(detail.Code), detail.Msg
		}
	}
	return InternalServerError.Code, err.Error()
}
