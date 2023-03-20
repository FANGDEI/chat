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
	ErrUserNameLength      = &Errno{Code: 20201, Msg: "用户名非法"}
	ErrUserPasswordLength  = &Errno{Code: 20202, Msg: "密码非法"}
	ErrUserExists          = &Errno{Code: 20203, Msg: "用户已存在"}
	ErrUserNotFound        = &Errno{Code: 20204, Msg: "用户不存在"}
	ErrUserPassword        = &Errno{Code: 20205, Msg: "用户名或密码错误"}
	ErrUserOldPassword     = &Errno{Code: 20206, Msg: "旧密码错误"}
	ErrUserFriendsNotFound = &Errno{Code: 20207, Msg: "未查询到用户好友"}

	// Friend
	ErrFriendExists         = &Errno{Code: 20301, Msg: "好友已存在"}
	ErrFriendNotExists      = &Errno{Code: 20302, Msg: "好友不存在"}
	ErrFriendApplyNotExists = &Errno{Code: 20303, Msg: "好友申请不存在"}

	// Chat
	ErrGetUserMsg       = &Errno{Code: 20401, Msg: "消息获取失败"}
	ErrRewriteMsg       = &Errno{Code: 20402, Msg: "消息写回失败"}
	ErrCreateHistory    = &Errno{Code: 20403, Msg: "消息记录创建失败"}
	ErrDuplicateRequest = &Errno{Code: 20404, Msg: "重复请求, 请等待对方验证"}

	// File
	ErrUploadFile = &Errno{Code: 20501, Msg: "上传失败"}

	// Group
	ErrGroupExists           = &Errno{Code: 20601, Msg: "群聊已存在"}
	ErrGroupRole             = &Errno{Code: 20602, Msg: "权限认证失败"}
	ErrGroupNotFound         = &Errno{Code: 20603, Msg: "群聊不存在"}
	ErrGroupMemberAlready    = &Errno{Code: 20604, Msg: "已是群成员"}
	ErrGroupDuplicateRequest = &Errno{Code: 20605, Msg: "重复请求, 请等待群主验证"}
	ErrGroupMemberNot        = &Errno{Code: 20606, Msg: "不是群成员"}
)
