package e

var MsgUser = map[int]string{
	SUCCESS:                        "ok",
	INVALID_PARAMS:                 "参数错误",
	ERROR_EXISR:                    "当前数据不存在",
	ERROR_AUTH_CHECK_TOKEN_FAIL:    "token格式不正确",
	ERROR_AUTH_CHECK_TOKEN_TIMEOUT: "token过期",
	ERROR_NOT_EXIST_USER:           "不存在该用户",
	ERROR_PASSWORD:                 "密码不正确",
	ERROR_EXIST_USER:               "用户名已存在",
}
