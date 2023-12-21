// @Author zhangxinmin 2023/12/20 21:25:00
package model

// RegisterRequest 注册请求结构体
type RegisterRequest struct {
	UserAccount   string `json:"userAccount"`
	UserPassword  string `json:"userPassword"`
	CheckPassword string `json:"checkPassword"`
}

// LoginRequest 登录请求结构体
type LoginRequest struct {
	UserAccount   string `json:"userAccount"`
	UserPassword  string `json:"userPassword"`
	CheckPassword string `json:"checkPassword"`
}
