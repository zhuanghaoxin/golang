package common

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// 全局常量，用于设置错误码
const (
	CodeSuccess           ErrCode = 0     // http请求成功
	CodeBodyBindErr       ErrCode = 10001 // 参数绑定错误
	CodeParamErr          ErrCode = 10002 // 请求参数不合法
	CodeRegisterErr       ErrCode = 10003 // 注册错误
	CodeLoginErr          ErrCode = 10003 // 登录错误
	CodeLogoutErr         ErrCode = 10004 // 登出错误
	CodeGetUserInfoErr    ErrCode = 10005 // 获取用户信息错误
	CodeUpdateUserInfoErr ErrCode = 10006 // 更新用户信息错误
	CodeSomethingError    ErrCode = 20000 // 未知错误
)

type (
	DebugType int // debug类型
	ErrCode   int // 错误码
)

// HttpResponse http独立请求返回结构体
type HttpResponse struct {
	Code ErrCode     `json:"code"`
	Data interface{} `json:"data"`
	Msg  string      `json:"message"`
}

func (res *HttpResponse) ResponseWithError(c *gin.Context, code ErrCode, msg string) {
	res.Code = code
	res.Msg = msg
	c.JSON(http.StatusInternalServerError, res)
}

func (res *HttpResponse) ResponseSuccess(c *gin.Context) {
	res.Code = CodeSuccess
	res.Msg = "success"
	c.JSON(http.StatusOK, res)
}

func (res *HttpResponse) ResponseWithData(c *gin.Context, data interface{}) {
	res.Code = CodeSuccess
	res.Msg = "success"
	res.Data = data
	c.JSON(http.StatusOK, res)
}
