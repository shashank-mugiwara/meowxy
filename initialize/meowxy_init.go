package initialize

import (
	"github.com/shashank-mugiwara/meowxy/conf"
	customLogger "github.com/shashank-mugiwara/meowxy/logger"
	"github.com/shashank-mugiwara/meowxy/metrics"
)

func InitRoutes() {
	metrics.RegisterRoutes(conf.GetMeowxyEngine(), customLogger.GetMeowxyLogger())
}

func InitConfig() {
	conf.SetUp("conf/config.ini")
}
