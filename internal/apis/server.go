package apis

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/minilabmemo/go-rest-arch/internal/config"
	"go.uber.org/zap"
)

func StartHttpServer(errChan chan error, engine *gin.Engine) {

	endpoint := fmt.Sprintf(":%d", config.ConfigData.Service.Port)
	go func() {
		errChan <- engine.Run(endpoint)
	}()

	zap.S().Infof("Listening on port: %d", config.ConfigData.Service.Port)
}
