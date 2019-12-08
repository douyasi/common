package common

var (
	JsonRespSuccess = 20000  // ok
	JsonRespInvalid = 20500  // invalid
	JsonRespCommonFail = 50000  // fail with error
)

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

// return Json
func Json(code int, message string, data interface{}) *JsonResult {
	return &JsonResult{
		Code:    code,
		Message: message,
		Data:    data,
	}
}

// success response with data
func Success(data interface{}) *JsonResult {
	return &JsonResult{
		Code:    JsonRespSuccess,
		Message: "success",
		Data:    data,
	}
}

// failure response with error
func Failure(err *Error) *JsonResult {
	return &JsonResult{
		Code:    err.Code,
		Message: err.Message,
		Data:    err.Data,
	}
}
