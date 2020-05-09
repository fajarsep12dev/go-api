package app

import C "github.com/fajarsep12dev/go-api/api/utils/constant"

// MsgFlags for mapping constant code
var MsgFlags = map[int]string{
	C.SUCCESS					:"Ok",
	C.ERROR						:"Fail",
	C.INVALID_PARAM				:"Invalid Parameters",
	C.ERROR_AUTH_CHECK_TOKEN_FAIL	:"Token Check Failed",
	C.ERROR_AUTH_CHECK_TOKEN_TIMEOUT:"Token Check Timeout",
}

// GetMsg get error information based on Code
func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}

	return MsgFlags[C.ERROR]
}