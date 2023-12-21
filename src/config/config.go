// @Author zhangxinmin 2023/12/19 23:40:00
package config

// 配置类
import (
	log "github.com/sirupsen/logrus"
	// viper是配置解决库，查找、加载和解组JSON、TOML、YAML、HCL、INI、envfile或Java属性格式的配置文件。
	"github.com/spf13/viper"
	"sync"
)

var (
	config GlobalConfig
	once   sync.Once
)

type SvrConfig struct {
	SvrName string `mapstructure:"svr_name"`
	Port    int    `mapstructure:"port"`
}

// DbConfig 用于db配置时用的结构体
type DbConfig struct {
	Url         string `mapstructure:"url"`
	User        string `mapstructure:"user"`          // 用户名
	Password    string `mapstructure:"password"`      // 密码
	Dbname      string `mapstructure:"dbname"`        // db名
	MaxIdleConn int    `mapstructure:"max_idle_conn"` // 最大空闲连接数
	MaxOpenConn int    `mapstructure:"max_open_conn"` // 最大打开的连接数
	MaxIdleTime int    `mapstructure:"max_idle_time"` // 连接最大空闲时间
}

// GlobalConfig 全局变量
type GlobalConfig struct {
	SvrConfig SvrConfig `mapstructure:"svr_config"`
	DbConfig  DbConfig  `mapstructure:"db"`
}

// GetGlobalConfig 实现一个方法用来返回全局变量
func GetGlobalConfig() *GlobalConfig {
	once.Do(readConfig)
	return &config
}

// 读取配置文件的信息
func readConfig() {
	// 设置读取文件的格式
	viper.SetConfigType("yml")
	viper.SetConfigFile("./src/config/config.yml")
	err := viper.ReadInConfig()
	// 判断是否能够读取
	if err != nil {
		panic("read config file err:" + err.Error())
	}
	// Unmarshal将配置反编组为Struct。确保标签结构的字段设置正确。
	err = viper.Unmarshal(&config)
	if err != nil {
		panic("config file unmarshal err:" + err.Error())
	}
	log.Infof("config === %+v", config)
}
