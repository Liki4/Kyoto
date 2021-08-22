package toolkit

type SuccessReturn struct {
	Msg   string      `json:"msg"`
	Data  interface{} `json:"data"`
	Error int         `json:"error"`
	//cache bool
	//code string
}

type ErrorReturn struct {
	Code int         `json:"code"`
	Msg  interface{} `json:"msg"`
}

func MakeSuccessReturn(data interface{}) SuccessReturn {
	return SuccessReturn{
		Msg:   "success",
		Data:  data,
		Error: 0,
	}
}

func MakeErrorReturn(code int, msg interface{}) ErrorReturn {
	return ErrorReturn{
		Code: code,
		Msg:  msg,
	}
}
