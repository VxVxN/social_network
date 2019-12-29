package response

type Response struct {
	Data  interface{} `json:"data"`
	Code  int         `json:"code"`
	Error string      `json:"error"`
}

func Success(data interface{}) Response {
	return Response{
		Data:  data,
		Code:  200,
		Error: "",
	}
}

func Error400(text string) Response {
	return error(400, text)
}

func Error500(text string) Response {
	return error(500, text)
}

func error(code int, text string) Response {
	return Response{
		Data:  nil,
		Code:  code,
		Error: text,
	}
}
