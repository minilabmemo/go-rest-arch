package api

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/minilabmemo/go-rest-arch/internal/config"

	"go.uber.org/zap"
)

func StartHttpServer(errChan chan error) {
	gin.SetMode(gin.ReleaseMode)

	engine := gin.New()

	loadRoutes(engine)

	endpoint := fmt.Sprintf(":%d", config.ConfigData.Service.Port)
	go func() {
		errChan <- engine.Run(endpoint)
	}()

	zap.S().Infof("Listening on port: %d", config.ConfigData.Service.Port)
}
