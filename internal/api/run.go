package api

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/minilabmemo/go-rest-arch/internal/card/delivery/ginrouter"
	"github.com/minilabmemo/go-rest-arch/internal/card/repository/mongo"
	"github.com/minilabmemo/go-rest-arch/internal/card/usecase"
	"github.com/minilabmemo/go-rest-arch/internal/config"

	"github.com/spf13/viper"
	"go.uber.org/zap"
)

func StartHttpServer(errChan chan error) {
	gin.SetMode(gin.ReleaseMode)

	engine := gin.New()
	root := engine.Group("service/api/v1")
	//TODO mid
	authorRepo := mongo.NewMongoCardRepository(true)
	//ar := _articleRepo.NewMysqlArticleRepository(dbConn)

	timeoutContext := time.Duration(viper.GetInt("context.timeout")) * time.Second
	au := usecase.NewArticleUsecase(authorRepo, timeoutContext)
	ginrouter.NewArticleHandler(root, au)

	iu := usecase.NewInfoUsecase(timeoutContext)
	ginrouter.NewInfoHandler(root, iu)
	//loadRoutes(engine)

	endpoint := fmt.Sprintf(":%d", config.ConfigData.Service.Port)
	go func() {
		errChan <- engine.Run(endpoint)
	}()

	zap.S().Infof("Listening on port: %d", config.ConfigData.Service.Port)
}
