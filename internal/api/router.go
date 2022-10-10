package api

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/minilabmemo/go-rest-arch/internal/card/delivery/ginrouter"
	"github.com/minilabmemo/go-rest-arch/internal/card/repository/mongo"
	"github.com/minilabmemo/go-rest-arch/internal/card/usecase"
	"github.com/minilabmemo/go-rest-arch/internal/config"
	"github.com/spf13/viper"
)

func loadRoutes(engine *gin.Engine) {
	engineGrp := engine.Group("service/api/v1")

	iu := usecase.NewInfoUsecase(*config.ConfigData)
	ginrouter.NewInfoHandler(engineGrp, iu)

	//TODO mid
	authorRepo := mongo.NewMongoCardRepository(true)
	//ar := _articleRepo.NewMysqlArticleRepository(dbConn)
	timeoutContext := time.Duration(viper.GetInt("context.timeout")) * time.Second
	au := usecase.NewArticleUsecase(authorRepo, timeoutContext)
	ginrouter.NewArticleHandler(engineGrp, au)

}
