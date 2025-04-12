package logger

import (
	"awesomeProject5/pkg/setting"
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

type LoggerZap struct {
	*zap.Logger
}

func NewLogger(config setting.LoggerSetting) *LoggerZap {
	logLevel := config.Log_Level
	// debug --> info --> warning --> error --> fatal --> panic
	var level zapcore.Level
	switch logLevel {
	case "debug":
		level = zapcore.DebugLevel
	case "info":
		level = zapcore.InfoLevel
	case "warning":
		level = zapcore.WarnLevel
	case "error":
		level = zapcore.ErrorLevel
	case "fatal":
		level = zapcore.FatalLevel
	case "panic":
		level = zapcore.PanicLevel
	default:
		level = zapcore.InfoLevel
	}
	encoder := getEncoderLog()
	hook := lumberjack.Logger{
		Filename:   config.File_log_name,
		MaxSize:    config.Max_size,
		MaxBackups: config.Max_backups,
		MaxAge:     config.Max_age,
		Compress:   config.Compress,
	}
	//Lumberjack sẽ tự động:
	//Xoay file log khi đạt MaxSize
	//Giữ tối đa MaxBackups file log cũ
	//Xóa file cũ sau MaxAge ngày
	//Nén file cũ nếu Compress=true

	//Bộ xử lý log chính
	core := zapcore.NewCore(
		encoder, // Format log
		zapcore.NewMultiWriteSyncer(
			zapcore.AddSync(os.Stdout), // Ghi log ra console
			zapcore.AddSync(&hook)),    // Ghi log ra file ( với rotation )
		level, // log level tối thiều
	)
	// logger := zap.New(core, zap.AddCaller())
	return &LoggerZap{
		zap.New(
			core, //Bộ xử lý log chính
			zap.AddCaller(),
			zap.AddStacktrace(zap.ErrorLevel),
		),
	}
}
func getEncoderLog() zapcore.Encoder {
	encodeConfig := zap.NewProductionEncoderConfig()
	// 1739868521.110879 -> 2025-02-18T15:48:41.110+0700
	encodeConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	// ts -> Time
	encodeConfig.TimeKey = "time"
	// from info INFO
	encodeConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	// caller
	encodeConfig.EncodeCaller = zapcore.ShortCallerEncoder
	return zapcore.NewJSONEncoder(encodeConfig)
}
