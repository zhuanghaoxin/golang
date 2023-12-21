// @Author zhangxinmin 2023/12/20 21:26:00
package utils

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net/http"
	"user-center/src/constant"
	"user-center/src/model"
)

func AuthMiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		if session, err := c.Cookie(constant.SessionKey); err == nil {
			// 如果没有出现错误且 session 不为空，说明存在有效的 session
			// 则调用 c.Next() 继续处理后续的请求处理函数，即允许通过该中间件
			if session != "" {
				var user model.User
				if err := json.Unmarshal([]byte(session), &user); err != nil {
					log.Errorf("Json Unmarshal error:%v", err)
					c.Abort()
					return
				}
				c.Next()
				return
			}
		}
		c.JSON(http.StatusUnauthorized, gin.H{"error": "err"})
		c.Abort()
	}
}
