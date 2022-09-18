package api

import "go.uber.org/zap"

func Run() {
	zap.S().Infof("OK")
}
