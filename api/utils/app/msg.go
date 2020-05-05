package app

import C "github.com/fajarsep12dev/go-api/api/utils/constant"

// MsgFlags for mapping constant code
var MsgFlags = map[int]string{
	C.Success					:"Ok",
	C.Error						:"Fail",
	C.InvalidParam				:"Invalid Parameters",
	C.ErrorAuthCheckTokenFail	:"Token Check Failed",
	C.ErrorAuthCheckTokenTimeOut:"Token Check Timeout",
}

// GetMsg get error information based on Code
func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}

	return MsgFlags[C.Error]
}