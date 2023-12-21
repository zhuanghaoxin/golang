// @Author zhangxinmin 2023/12/20 21:26:00
package utils

import (
	"crypto/md5"
	"encoding/hex"
	"user-center/src/constant"
	"user-center/src/model"
)

// EncryptMd5 md5 的加密算法
func EncryptMd5(userPassword string) string {
	h := md5.New()
	// 将userPassword与salt进行加密
	h.Write([]byte(constant.SALT + userPassword))
	// 最后返回h的16进制编码
	// sum方法是将括号中的值追加到h后面，但不改变哈希状态
	return hex.EncodeToString(h.Sum(nil))
}

// GetSafetyUser 用户信息脱敏 保留必要的用户信息
func GetSafetyUser(user *model.User) model.User {
	safetyUser := model.User{}
	safetyUser.Id = user.Id
	safetyUser.Username = user.Username
	safetyUser.UserAccount = user.UserAccount
	safetyUser.AvatarUrl = user.AvatarUrl
	safetyUser.Gender = user.Gender
	safetyUser.Phone = user.Phone
	safetyUser.Email = user.Email
	safetyUser.UserStatus = user.UserStatus
	safetyUser.UserRole = user.UserRole
	safetyUser.CreateTime = user.CreateTime
	return safetyUser
}
