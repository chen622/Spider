package utils

type apiJson struct {
	Code int32       `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func Ok(objects interface{}, msg string) (apijson *apiJson) {
	apijson = &apiJson{Code: 200, Data: objects, Msg: msg}
	return
}

func TokenInvalid(msg string) (apijson *apiJson) {
	apijson = &apiJson{Code: 401, Msg: msg}
	return
}

func NotOk(error *Error) (apijson *apiJson) {
	apijson = &apiJson{Code: error.Code, Msg: error.Msg}
	return
}
