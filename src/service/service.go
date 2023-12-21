// @Author zhangxinmin 2023/12/20 21:26:00
package service

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"regexp"
	"user-center/src/common"
	"user-center/src/db"
	"user-center/src/model"
	"user-center/src/utils"
)

// Register 写具体的业务逻辑
func Register(req *model.RegisterRequest, c *gin.Context) (int64, error) {
	res := &common.HttpResponse{}
	//账号密码格式验证
	if req.UserAccount == "" || req.UserPassword == "" || req.CheckPassword == "" {
		log.Info("UserAccount or UserPassword or CheckPassword is null")
		res.ResponseWithError(c, common.CodeParamErr, "输入不少于6位")
		return 0, nil
	}
	if len(req.UserAccount) < 6 || len(req.UserPassword) < 6 || len(req.CheckPassword) < 6 {
		log.Info("UserAccount or UserPassword or CheckPassword is too short")
		res.ResponseWithError(c, common.CodeParamErr, "输入不能小于六位")
		return 0, nil
	}
	//账号不能有特殊字符
	_, err := regexp.MatchString("/[`~!@#$%^&*()_\\-+=<>?:\"{}|,.\\/;'\\\\[\\]·~！@#￥%……&*（）——\\-+={}|《》？：“”【】、；‘'，。、]/", req.UserAccount)
	if err != nil {
		log.Infof("UserAccount have not available char:%v", err)
		res.ResponseWithError(c, common.CodeParamErr, "账号不能包含特殊字符")
		return 0, nil
	}
	// 密码与校验密码
	if req.UserPassword != req.CheckPassword {
		log.Info("UserPassword != CheckPassword")
		res.ResponseWithError(c, common.CodeParamErr, "密码与校验密码不相同")
		return 0, nil
	}
	// 账号不能重复
	if db.GetDB().Table("user").Where("userAccount=?", req.UserAccount).RecordNotFound() {
		log.Info("userAccount is already setup")
		res.ResponseWithError(c, common.CodeParamErr, "账号已被注册")
		return 0, nil
	}

	// 对密码进行加密存储
	entryPassword := utils.EncryptMd5(req.UserPassword)
	user := &model.User{
		UserAccount:  req.UserAccount,
		UserPassword: entryPassword,
	}

	//插入数据
	log.Infof("user ====== %+v", user)
	if err = db.GetDB().Table("user").Model(&model.User{}).Create(user).Error; err != nil {
		log.Warningf("Create user fail:%v", err)
		res.ResponseWithError(c, common.CodeUpdateUserInfoErr, "数据库创建用户失败")
		return 0, err
	}
	return user.Id, nil
}

// Login 登陆功能
func Login(req *model.LoginRequest, c *gin.Context) *model.User {
	res := &common.HttpResponse{}
	// 1. 校验
	if req.UserAccount == "" || req.UserPassword == "" {
		log.Info("UserAccount or UserPassword is null")
		res.ResponseWithError(c, common.CodeParamErr, "输入不能为空")
		return nil
	}
	// 2. 判断账号和密码的长度是否符合要求
	if len(req.UserAccount) < 6 || len(req.UserPassword) < 6 {
		log.Info("UserAccount or UserPassword is too short")
		res.ResponseWithError(c, common.CodeParamErr, "输入不能小于六位")
		return nil
	}
	// 3. 账号不能包含特殊字符
	_, err := regexp.MatchString("/[`~!@#$%^&*()_\\-+=<>?:\"{}|,.\\/;'\\\\[\\]·~！@#￥%……&*（）——\\-+={}|《》？：“”【】、；‘'，。、]/", req.UserAccount)
	if err != nil {
		log.Infof("UserAccount have not available char:%v", err)
		res.ResponseWithError(c, common.CodeParamErr, "账号不能包含特殊字符")
		return nil
	}
	// 对用户传递过来的密码进行加密，并与数据库中已经加密的密码进行比对
	encryptPassword := utils.EncryptMd5(req.UserPassword)
	// 查询用户是否存在
	user := &model.User{}
	if err = db.GetDB().Table("user").Where("isDelete=0 and userAccount=? and userPassword=?", req.UserAccount, encryptPassword).First(&user).Error; err != nil {
		log.Warningf("Cannot find the user")
		res.ResponseWithError(c, common.CodeGetUserInfoErr, "查询不到用户信息")
		return nil
	}
	return user
}

// SearchUser 根据用户名查询用户信息
func SearchUser(username string, c *gin.Context) []model.User {
	rsp := &common.HttpResponse{}
	var user []model.User
	if err := db.GetDB().Table("user").Where("isDelete=0 and username like ? or username is null", "%"+username+"%").Find(&user).Error; err != nil {
		log.Warningf("Serch user information fail:%v", err)
		rsp.ResponseWithError(c, common.CodeGetUserInfoErr, "查询用户信息出错")
		return nil
	}
	return user
}

// DeleteById 根据 id 逻辑删除用户
func DeleteById(id int64, c *gin.Context) {
	rsp := &common.HttpResponse{}
	if err := db.GetDB().Table("user").Where("id=?", id).Update("isDelete", 1).Error; err != nil {
		log.Warningf("Delete the user false")
		rsp.ResponseWithError(c, common.CodeGetUserInfoErr, "删除用户失败")
	}
	log.Warningf("Delete the user true,")
}
