package errors

var MsgFlags = map[int]string{
	SUCCESS:        "ok",
	INVALID_PARAMS: "invalid parameter",
	UNAUTHORIZED:   "unauthorized",
	NOT_FOUND:      "not found",
	SERVER_ERROR:   "server error",

	USER_ALREADY_EXIST:          "User name or email already exist",
	INVALID_USER_NAME_PASSWORD:  "User name or password is invalid",
	HASH_PASSWORD_ERROR:         "Hash password error",
	GEN_TOKEN_ERROR:             "Generate token error",
	INACTIVE_USER:               "User is inactive",
	USER_BLOCKED:                "User is blocked",
	INVALID_TOKEN:               "Invalid token",
	ERROR_EXIST_EMAIL:           "Email already exist",
	USER_DELETED:                "User is deleted",
	ADMIN_DELETED:               "Admin is deleted",
	INACTIVE_ADMIN:              "Admin is inactive",
	UNAUTHORIZED_ACCESS:         "Unauthorized access",
	CATEGORY_IDS_INVALID_PARAMS: "Category ids invalid parameter",
}

func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}
	return MsgFlags[SERVER_ERROR]
}
