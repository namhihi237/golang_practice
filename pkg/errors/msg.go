package errors

var MsgFlags = map[int]string{
	SUCCESS:        "ok",
	INVALID_PARAMS: "invalid parameter",
	UNAUTHORIZED:   "unauthorized",
	NOT_FOUND:      "not found",
	SERVER_ERROR:   "server error",

	USER_ALREADY_EXIST:  "User name or email already exist",
	HASH_PASSWORD_ERROR: "Hash password error",
}

func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}
	return MsgFlags[SERVER_ERROR]
}
