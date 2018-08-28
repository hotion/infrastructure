// Package code to define some code
package code

import (
	"reflect"
)

const (
	CodeOk = iota
	CodeParamInvalid
	CodeSystemErr
	CodeNoPermission
	CodeServerTimeout
	CodeResourceNotFound

	CodeUserExisted
	CodeNotLogin
	CodeIllegeOP
	CodeWrongTicket

	CodeWrongAuthCode
	CodeWrongExpired
	CodeReqAuthCodeTimes
	CodeRaiseProcessEnd
	CodeNotEnoughRemind
	CodeNotEnoughMoney

	ErrNoSuchCode = "错误码未定义"
)

var messages = map[int]string{
	CodeOk:               "成功",
	CodeParamInvalid:     "参数非法",
	CodeSystemErr:        "系统错误",
	CodeNoPermission:     "没有权限",
	CodeServerTimeout:    "服务超时",
	CodeResourceNotFound: "资源未找到",

	CodeUserExisted: "用户已经存在",
	CodeNotLogin:    "未登录",
	CodeIllegeOP:    "非法操作",
	CodeWrongTicket: "Ticket不正确",

	CodeWrongAuthCode:    "验证码错误",
	CodeWrongExpired:     "验证码过期",
	CodeRaiseProcessEnd:  "投资未开始",
	CodeNotEnoughRemind:  "产品余额不足",
	CodeNotEnoughMoney:   "用户余额不足",
	CodeReqAuthCodeTimes: "上次请求时间小于60s或者相同手机号码请求次数过多",
}

// CodeInfo define a CodeInfo type
type CodeInfo struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

// NewCodeInfo create a new *CodeInfo
func NewCodeInfo(code int, message string) *CodeInfo {
	if message == "" {
		message = GetMessage(code)
	}

	return &CodeInfo{
		Code:    code,
		Message: message,
	}
}

// GetCodeInfo get CodeInfo with specified code
func GetCodeInfo(code int) *CodeInfo {
	return &CodeInfo{
		Code:    code,
		Message: GetMessage(code),
	}
}

// GetMessage get code desc from messages
func GetMessage(code int) string {
	v, ok := messages[code]
	if !ok {
		return ErrNoSuchCode
	}
	return v
}

// FillCodeInfo ... fill a response struct will *CodeInfo
// TODO: validate v
func FillCodeInfo(v interface{}, ci *CodeInfo) interface{} {
	ele := reflect.ValueOf(v).Elem()
	field := ele.FieldByName("CodeInfo")

	if ci.Message == "" {
		ci.Message = GetMessage(ci.Code)
	}

	// set field
	field.FieldByName("Code").SetInt(int64(ci.Code))
	field.FieldByName("Message").SetString(ci.Message)

	return v
}