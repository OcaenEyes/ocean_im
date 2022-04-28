package errorcode

var (
	Ok = &ErrorCode{Code: 00000, Message: "Success"}

	//系统错误,前缀为 9

	InternalServerError = &ErrorCode{Code: 90001, Message: "内部服务器错误"}

	//数据库错误 8 开头

	DataErr = &ErrorCode{Code: 80001, Message: "数据处理异常"}
	//用户信息错误 7 开头

	AuthErr = &ErrorCode{Code: 70001, Message: "鉴权失败,请重新登录"}

	//群相关错误 4 开头

	GroupReqParamErr  = &ErrorCode{Code: 40001, Message: "参数异常"}
	GroupReqReceiveNo  = &ErrorCode{Code: 40002, Message: "暂无接受到加群请求"}
	GroupReqSendsNo  = &ErrorCode{Code: 40003, Message: "暂无发送加群请求"}
	GroupsNo  = &ErrorCode{Code: 40004, Message: "群不存在"}

	//消息相关错误 3 开头

	MsgNo  = &ErrorCode{Code: 30001, Message: "没有消息"}
	MsgSendNotFriend  = &ErrorCode{Code: 30002, Message: "挺遗憾的，我们不再是好友了~你本次发的消息不会发送成功"}
	MsgSendNotOnline  = &ErrorCode{Code: 30003, Message: "对方不在线"}



	//好友相关错误 2 开头

	FriendReqParamErr  = &ErrorCode{Code: 20001, Message: "参数异常"}
	FriendReqReceiveNo = &ErrorCode{Code: 20002, Message: "暂无接受到好友请求"}
	FriendReqSendsNo   = &ErrorCode{Code: 20003, Message: "暂无发送好友请求"}
	FriendsNo          = &ErrorCode{Code: 20004, Message: "暂无好友"}

	//用户相关错误 1 开头

	UserPassErr   = &ErrorCode{Code: 10001, Message: "账户或密码错误"}
	UserNotExist  = &ErrorCode{Code: 10002, Message: "账户不存在"}
	UserIsExist   = &ErrorCode{Code: 10003, Message: "账户已存在"}
	UserInfoEmpty = &ErrorCode{Code: 10004, Message: "账号或密码未填写"}
)
