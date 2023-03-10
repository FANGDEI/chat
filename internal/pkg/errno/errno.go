package errno

// 错误码
// 第 1 位 : 服务级错误码; 比如 1 为系统级错误; 2 为普通错误, 通常是由用户非法操作引起
// 第2 3位 : 模块级错误码, 比如 01 为用户模块; 02 为订单模块
// 第4 5位 : 具体的错误码, 比如 01 为手机号不合法; 02 为验证码输入错误

var (
	OK = &Errno{Code: 0, Msg: "OK"}

	InternalServerError = &Errno{Code: 10001, Msg: "Internal server error"}

	// token
	ErrTokenGenerate = &Errno{Code: 10002, Msg: "Token generate error"}

	// MySQL
	ErrDatabase = &Errno{Code: 10003, Msg: "Database error"}

	// Redis
	ErrRedis = &Errno{Code: 10101, Msg: "Redis error"}

	ErrBadRequest = &Errno{Code: 20001, Msg: "Request error"}

	// Email
	ErrEmailFormat = &Errno{Code: 20101, Msg: "Email format error"}
	ErrEmailBan    = &Errno{Code: 20102, Msg: "Email can't send to same account in one minute"}
	ErrEmailAuth   = &Errno{Code: 20103, Msg: "Email code verify failed"}

	// User
	ErrUserNameLength     = &Errno{Code: 20201, Msg: "User length of the Name is illegal"}
	ErrUserPasswordLength = &Errno{Code: 20202, Msg: "User length of the Password is illegal"}
	ErrUserExists         = &Errno{Code: 20203, Msg: "User already exists"}
	ErrUserNotFound       = &Errno{Code: 20204, Msg: "User not found"}
	ErrUserPassword       = &Errno{Code: 20205, Msg: "User name or password is wrong"}
	ErrUserOldPassword    = &Errno{Code: 20206, Msg: "User old password is wrong"}
)
