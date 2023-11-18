package response

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetResponse() *Response {
	return &Response{
		httpCode: http.StatusOK,
		Result: &Result{
			Code:    0,
			Message: "",
			Data:    nil,
			Cost:    "",
		},
	}
}
func BadRequest(ctx *gin.Context, data ...any) {
	fmt.Println("BadRequest here")
	GetResponse().withDataAndHttpCode(http.StatusBadRequest, ctx, data)
}

// Success 业务成功响应
func Success(ctx *gin.Context, data ...any) {
	if data != nil {
		fmt.Println("Success here")
		GetResponse().WithDataSuccess(ctx, data[0])
		return
	}
	GetResponse().Success(ctx)
}

// Fail 业务失败响应
func Fail(ctx *gin.Context, code int, message *string, data ...any) {
	var msg string
	if message == nil {
		msg = ""
	} else {
		msg = *message
	}
	if data != nil {
		GetResponse().WithData(data[0]).FailCode(ctx, code, msg)
		return
	}
	GetResponse().FailCode(ctx, code, msg)
}

type Result struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
	Cost    string      `json:"cost"`
}

type Response struct {
	httpCode int
	Result   *Result
}

// Fail 错误返回
func (r *Response) Fail(ctx *gin.Context) *Response {
	r.SetCode(http.StatusInternalServerError)
	r.json(ctx)
	return r
}
