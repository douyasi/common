package resp

import "time"

var (
	SuccessCode   = 20000  // success
	RejectionCode = 40000  // invalid params, validation fail etc.
	FailureCode   = 50000  // failure
	// other error or failure please throw JsonError, code > 50000
)

// empty data
type EmptyData struct {
}

// jsonError
type JsonError struct {
	Code int
	Message string
	Data interface{}
}

// jsonResult
type JsonResult struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
	Timestamp int64     `json:"timestamp"`
}

// return Json
func Json(code int, message string, data interface{}) *JsonResult {
	if data == nil {
		data = EmptyData{}
	}
	return &JsonResult{
		code,
		message,
		data,
		time.Now().Unix(),
	}
}

// success response with data
func Success(data interface{}) *JsonResult {
	if data == nil {
		data = EmptyData{}
	}
	return &JsonResult{
		SuccessCode,
		"success",
		data,
			time.Now().Unix(),
	}
}

// error response throw JsonError
func Error(err *JsonError) *JsonResult {
	return &JsonResult{
		err.Code,
		err.Message,
		err.Data,
		time.Now().Unix(),
	}
}

// rejection response when invalid params, validation fail etc.
func Rejection(message string) *JsonResult {
	var data EmptyData
	return &JsonResult{
		RejectionCode,
		message,
		data,
		time.Now().Unix(),
	}
}

// failure response with empty data and default message
func Failure() *JsonResult {
	var data EmptyData
	return &JsonResult{
		FailureCode,
		"failure",
		data,
		time.Now().Unix(),
	}
}
