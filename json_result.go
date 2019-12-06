package common


type Error struct {
	Code int
	Message string
	Data interface{}
}

type JsonResult struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func Api(code int, message string, data interface{}) *JsonResult {
	return &JsonResult{
		Code:    code,
		Message: message,
		Data:    data,
	}
}

func Success(data interface{}) *JsonResult {
	return &JsonResult{
		Code:    200,
		Message: "success",
		Data:    data,
	}
}

func Fail(err *Error) *JsonResult {
	return &JsonResult{
		Code:    err.Code,
		Message: err.Message,
		Data:    err.Data,
	}
}