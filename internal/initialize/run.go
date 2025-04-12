package initialize

import (
	"awesomeProject5/global"
	"fmt"
	"go.uber.org/zap"
)

func Run() {
	LoadConfig()
	fmt.Println("Loading configuration...", global.Config.Mysql.Username)
	InitLogger()
	global.Logger.Info("Logger init success", zap.String("oke", "success"))
	InitMySQL()
	r := InitRouter(InitApp(global.Mdb))
	r.Run(":2004")
}
