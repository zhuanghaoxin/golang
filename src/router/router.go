// @Author zhangxinmin 2023/12/19 23:42:00
package router

// 路由
import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"strconv"
	"user-center/src/api"
	"user-center/src/config"
	"user-center/src/utils"
)

// InitRouterAndServer 初始化路由和服务端口
func InitRouterAndServer() {
	// 根据输入字符设置gin模式
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	userGroup := r.Group("/api/user")
	// user路由组
	{
		userGroup.POST("/register", api.Register)
		userGroup.POST("/login", api.Login)
		userGroup.POST("/logout", utils.AuthMiddleWare(), api.LogOut)
		userGroup.GET("/current", utils.AuthMiddleWare(), api.GetCurrentUser)
		userGroup.GET("/search", utils.AuthMiddleWare(), api.SearchUser)
		userGroup.POST("/delete", utils.AuthMiddleWare(), api.DeleteUser)
	}

	if err := r.Run(":" + strconv.Itoa(config.GetGlobalConfig().SvrConfig.Port)); err != nil {
		log.Error("start server err:" + err.Error())
	}
}
