// @Author zhangxinmin 2023/12/19 23:39:00
package api

// api接口
import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"strconv"
	"user-center/src/common"
	"user-center/src/constant"
	"user-center/src/model"
	"user-center/src/service"
	"user-center/src/utils"
)

func Register(c *gin.Context) {
	req := &model.RegisterRequest{}
	rsp := &common.HttpResponse{}

	// 从前端接收请求
	if err := c.ShouldBindJSON(&req); err != nil {
		log.Errorf("Register request json err %v", err)
		rsp.ResponseWithError(c, common.CodeBodyBindErr, err.Error())
		return
	}

	// 处理请求
	id, err := service.Register(req, c)
	if err != nil {
		rsp.ResponseWithError(c, common.CodeRegisterErr, err.Error())
		return
	}

	rsp.ResponseWithData(c, id)
}

func Login(c *gin.Context) {
	req := &model.LoginRequest{}
	res := &common.HttpResponse{}
	if err := c.ShouldBindJSON(&req); err != nil {
		log.Errorf("Login request json err %v", err)
		res.ResponseWithError(c, common.CodeBodyBindErr, err.Error())
		return
	}
	log.Infof("login start, user: %s, password: %s", req.UserAccount, req.UserPassword)
	user := service.Login(req, c)
	if user == nil {
		res.ResponseWithError(c, common.CodeLoginErr, "登录失败")
		return
	}
	// 先对user进行序列化
	session, err := json.Marshal(user)
	if err != nil {
		log.Errorf("Json marshal error:%v", err)
		res.ResponseWithError(c, common.CodeSomethingError, "后端出现未知错误")
	}
	// 登陆成功后就设置session
	c.SetCookie(constant.SessionKey, string(session), constant.CookieExpire, "/", "", false, true)

	// 信息脱敏
	safetyUser := utils.GetSafetyUser(user)

	// 返回结果
	res.ResponseWithData(c, safetyUser)
}

// GetCurrentUser 获取当前用户的信息
func GetCurrentUser(c *gin.Context) {
	res := &common.HttpResponse{}
	session, err := c.Cookie(constant.SessionKey)
	if err != nil {
		log.Warningf("Get session fail: %v", err)
		res.ResponseWithError(c, common.CodeSomethingError, "后端出现未知错误")
		return
	}
	var user model.User
	if err := json.Unmarshal([]byte(session), &user); err != nil {
		log.Warningf("Json Unmarshal error:%v", err)
		res.ResponseWithError(c, common.CodeSomethingError, "后端出现未知错误")
		return
	}
	// 用户信息脱敏
	safetyUser := utils.GetSafetyUser(&user)
	res.ResponseWithData(c, safetyUser)
}

// SearchUser 根据username查询用户
func SearchUser(c *gin.Context) {
	res := &common.HttpResponse{}
	// 从url中提取对应的键值对
	username := c.Query("username")
	userList := service.SearchUser(username, c)
	var safetyUserList []model.User
	for _, user := range userList {
		safetyUser := utils.GetSafetyUser(&user)
		safetyUserList = append(safetyUserList, safetyUser)
	}
	res.ResponseWithData(c, safetyUserList)
}

// DeleteUser 删除指定用户
func DeleteUser(c *gin.Context) {
	res := &common.HttpResponse{}
	// 从url提取id，由于提取的id是字符串格式的，但数据库是int类型，因此需要进行类型转换
	uid := c.Query("id")
	log.Info("id:,", uid)
	id, _ := strconv.ParseInt(uid, 10, 64)
	log.Info("id:,", id)
	service.DeleteById(id, c)
	res.ResponseSuccess(c)
}

// LogOut 用户登出，将服务器的session删除
func LogOut(c *gin.Context) {
	res := &common.HttpResponse{}
	c.SetCookie(constant.SessionKey, "", -1, "/", "", false, true)
	res.ResponseSuccess(c)
}
