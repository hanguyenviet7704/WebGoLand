package initialize

import (
	"awesomeProject5/global"
	"awesomeProject5/pkg/logger"
)

func InitLogger() {
	global.Logger = logger.NewLogger(global.Config.Logger)
}
