package global

import (
	"awesomeProject5/pkg/logger"
	"awesomeProject5/pkg/setting"
	"gorm.io/gorm"
)

var (
	Config setting.Config
	Logger *logger.LoggerZap
	Mdb    *gorm.DB
)
