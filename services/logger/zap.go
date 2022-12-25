package logger

import "go.uber.org/zap"

//TODO setting
func InitLogger() {
	logger, _ := zap.NewDevelopment()

	zap.ReplaceGlobals(logger) // // 配置 zap 包的全局變量
	zap.S().Info("InitLogger done")
}
