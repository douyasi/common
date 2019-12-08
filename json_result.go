package common

var (
	JsonRespSuccess = 20000  // success
	JsonRespCommonFailure = 50000  // failure
)

type JsonError struct {
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
		code,
		message,
		data,
	}
}

// success response with data
func Success(data interface{}) *JsonResult {
	return &JsonResult{
		JsonRespSuccess,
		"success",
		data,
	}
}

// error response throw JsonError
func Error(err *JsonError) *JsonResult {
	return &JsonResult{
		err.Code,
		err.Message,
		err.Data,
	}
}

// failure response with message
func Failure(message string) *JsonResult {
	type EmptyData interface {
	}
	var data EmptyData
	return &JsonResult{
		JsonRespCommonFailure,
		message,
		data,
	}
}
