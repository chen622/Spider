package utils

type Error struct {
	Code int32  `json:"code"`
	Msg  string `json:"msg"`
}

func E_404() (err *Error) {
	err = &Error{Code: 404, Msg: "Url not found"}
	return
}

func E_500() (err *Error) {
	err = &Error{Code: 500, Msg: "Oups! Something went wrong, try again"}
	return
}

func E_All() (err *Error) {
	err = &Error{Code: 0, Msg: "Oups! Something wrong which not expect!"}
	return
}
