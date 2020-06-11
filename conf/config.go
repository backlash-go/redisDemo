package conf

import (
	"flag"
	"fmt"
	"github.com/spf13/viper"
)

var (
	configFile = flag.String("conf", "conf/app-dev.yaml", "path of config file")
)

func InitialRedis() {
	if *configFile == "" {
		flag.Usage()
	}
	// 读取配置文件
	viper.SetConfigFile(*configFile)
	err := viper.ReadInConfig()
	if err != nil {
		errStr := fmt.Sprintf("viper read config is failed, err is %v configFile is %s ", err, configFile)
		panic(errStr)
	}
	authRedisConf := viper.GetStringMapString("authRedis")
	InitRedis(fmt.Sprintf("%s:%s", authRedisConf["host"], authRedisConf["port"]), authRedisConf["password"], authRedisConf["db"])

}
