package errors

var MsgFlags = map[int]string{
	SUCCESS:        "ok",
	INVALID_PARAMS: "invalid parameter",
	UNAUTHORIZED:   "unauthorized",
	NOT_FOUND:      "not found",
	SERVER_ERROR:   "server error",
}

func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}
	return MsgFlags[SERVER_ERROR]
}
