package errno

// 错误码
// 第 1 位 : 服务级错误码; 比如 1 为系统级错误; 2 为普通错误, 通常是由用户非法操作引起
// 第2 3位 : 模块级错误码, 比如 01 为用户模块; 02 为订单模块
// 第4 5位 : 具体的错误码, 比如 01 为手机号不合法; 02 为验证码输入错误

var (
	OK = &Errno{Code: 0, Msg: "OK"}

	InternalServerError = &Errno{Code: 10001, Msg: "Internal server error"}

	ErrTokenGenerate = &Errno{Code: 10002, Msg: "Token generate error"}
	ErrDatabase      = &Errno{Code: 10003, Msg: "Database error"}
	ErrRedis         = &Errno{Code: 10101, Msg: "Redis error"}

	// Email
	ErrEmailFormat = &Errno{Code: 20101, Msg: "邮箱格式错误"}
	ErrEmailBan    = &Errno{Code: 20102, Msg: "未到邮件可发送时间"}
	ErrEmailAuth   = &Errno{Code: 20103, Msg: "验证码错误"}

	// User
	ErrUserNameLength     = &Errno{Code: 20201, Msg: "用户名非法"}
	ErrUserPasswordLength = &Errno{Code: 20202, Msg: "密码非法"}
	ErrUserExists         = &Errno{Code: 20203, Msg: "用户已存在"}
	ErrUserNotFound       = &Errno{Code: 20204, Msg: "用户不存在"}
	ErrUserPassword       = &Errno{Code: 20205, Msg: "用户名或密码错误"}
	ErrUserOldPassword    = &Errno{Code: 20206, Msg: "旧密码错误"}

	// Friend
	ErrFriendExists    = &Errno{Code: 20301, Msg: "Friend already exists"}
	ErrFriendNotExists = &Errno{Code: 20302, Msg: "Friend not friend"}

	// Chat
	ErrGetUserMsg    = &Errno{Code: 20401, Msg: "Chat get message error"}
	ErrRewriteMsg    = &Errno{Code: 20402, Msg: "Chat rewrite the message error"}
	ErrCreateHistory = &Errno{Code: 20403, Msg: "Chat create history error"}
)
