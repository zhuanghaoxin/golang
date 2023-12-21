// @Author zhangxinmin 2023/12/19 23:40:00
package db

// 数据库连接
import (
	"github.com/jinzhu/gorm"
	// 管理mysql数据库连接
	"github.com/niuniumart/gosdk/gormcli"
	log "github.com/sirupsen/logrus"
	"user-center/src/config"
	//数据库连接所需

	"sync"
)

// 先规定全局变量,后面提供给其他包使用
var (
	db     *gorm.DB
	dbOnce sync.Once
)

// 数据库连接方法
func connectDb() {
	var err error
	// 设置连接池的配置
	gormcli.Factory = gormcli.GormFactory{
		MaxIdleConn: config.GetGlobalConfig().DbConfig.MaxIdleConn,
		MaxConn:     config.GetGlobalConfig().DbConfig.MaxOpenConn,
		IdleTimeout: config.GetGlobalConfig().DbConfig.MaxIdleTime,
	}
	// 数据库连接
	db, err = gormcli.Factory.CreateGorm(
		config.GetGlobalConfig().DbConfig.User,
		config.GetGlobalConfig().DbConfig.Password,
		config.GetGlobalConfig().DbConfig.Url,
		config.GetGlobalConfig().DbConfig.Dbname,
	)
	// 如果连接失败
	if err != nil {
		log.Fatalf("failed to connect database:%v", err)
		return
	}
	// 测试访问是否正常
	err = db.DB().Ping()
	if err != nil {
		log.Fatalf("Database connection is not available:%v", err)
		return
	}
	log.Info("Database connection is available")
}

func GetDB() *gorm.DB {
	dbOnce.Do(connectDb)
	return db
}
