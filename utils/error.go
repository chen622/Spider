package utils

import "encoding/json"

type Error struct {
	Code int32  `json:"code"`
	Msg  string `json:"msg"`
}

func E_201() (err *Error) {
	err = &Error{Code: 201, Msg: "未开放注册!"}
	return
}

//E_402 表单验证错误
func E_402(errMap map[string]interface{}) (err *Error) {
	msg, _ := json.Marshal(errMap)
	err = &Error{Code: 402, Msg: string(msg)}
	return
}

func E_403() (err *Error) {
	err = &Error{Code: 403, Msg: "用户名或密码错误!"}
	return
}

func E_404() (err *Error) {
	err = &Error{Code: 404, Msg: "未找到目标！"}
	return
}

func E_405() (err *Error) {
	err = &Error{Code: 405, Msg: "用户名已被使用!"}
	return
}

func E_500() (err *Error) {
	err = &Error{Code: 500, Msg: "抱歉！请求异常，请重试!"}
	return
}

func E_501() (err *Error) {
	err = &Error{Code: 501, Msg: "参数错误！"}
	return
}

func E_All() (err *Error) {
	err = &Error{Code: 0, Msg: "抱歉！请求异常，请重试"}
	return
}
