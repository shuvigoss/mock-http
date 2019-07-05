package model

import "net/http"

type Rsp struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

const (
	ServerError         = 500      //服务端异常
	StatusParamError    = 10904001 //请求参数错误
	StatusDatabaseError = 10904002 // 数据库异常
)

func Success(data interface{}) Rsp {
	return Rsp{
		Status:  http.StatusOK,
		Message: "OK",
		Data:    data,
	}
}

func SuccessWithCode(code int, data interface{}) Rsp {
	return Rsp{
		Status:  code,
		Message: "OK",
		Data:    data,
	}
}

func SuccessMsg(data interface{}, msg string) Rsp {
	return Rsp{
		Status:  http.StatusOK,
		Message: msg,
		Data:    data,
	}
}

func FailWithData(code int, data interface{}) Rsp {
	return Rsp{
		Status:  code,
		Message: "FAIL",
		Data:    data,
	}
}

func FailWithMsg(code int, message string) Rsp {
	return Rsp{
		Status:  code,
		Message: message,
		Data:    nil,
	}
}

func IsSuccess(Rsp *Rsp) bool {
	return Rsp.Status == http.StatusOK
}
